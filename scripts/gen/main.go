package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
)

// templateContext is the context passed to the template engine for generating tuple code and test files.
type templateContext struct {
	Indexes             []int
	Len                 int
	TypeName            string
	TypeDecl            string
	GenericTypesDecl    string
	GenericTypesForward string
}

const minTupleLength = 1
const maxTupleLength = 9

var funcMap = template.FuncMap{
	"quote": func(value interface{}) string {
		return fmt.Sprintf("%q", fmt.Sprint(value))
	},
}

//go:embed tuple.tpl
var codeTplContent string

//go:embed tuple_test.tpl
var testTplContent string

// main generates the tuple code and test files by executing the template engine for the "tuple.tpl" and "tuple_test.tpl" files.
func main() {
	outputDir := os.Args[1]

	codeTpl, err := template.New("tuple").Funcs(funcMap).Parse(codeTplContent)
	if err != nil {
		panic(err)
	}

	testTpl, err := template.New("tuple_test").Funcs(funcMap).Parse(testTplContent)
	if err != nil {
		panic(err)
	}

	for tupleLength := minTupleLength; tupleLength <= maxTupleLength; tupleLength++ {
		indexes := make([]int, tupleLength)
		for index := range indexes {
			indexes[index] = index + 1
		}

		decl := genTypesDecl(indexes)
		forward := genTypesForward(indexes)
		context := templateContext{
			Indexes:             indexes,
			Len:                 tupleLength,
			TypeName:            fmt.Sprintf("T%d[%s]", tupleLength, forward),
			TypeDecl:            fmt.Sprintf("T%d[%s]", tupleLength, decl),
			GenericTypesDecl:    decl,
			GenericTypesForward: forward,
		}

		filesToGenerate := []struct {
			fullPath string
			tpl      *template.Template
		}{
			{
				fullPath: path.Join(outputDir, fmt.Sprintf("tuple%d.go", tupleLength)),
				tpl:      codeTpl,
			},
			{
				fullPath: path.Join(outputDir, fmt.Sprintf("tuple%d_test.go", tupleLength)),
				tpl:      testTpl,
			},
		}

		for _, file := range filesToGenerate {
			fmt.Printf("Generating file %q...\n", file.fullPath)
			generateFile(context, file.fullPath, file.tpl)
		}
	}
}

// generateFile generates the file at outputFilePath according to the template tpl.
// The template engine is given the context parameter as data (can be used as "." in the templates).
func generateFile(context templateContext, outputFilePath string, tpl *template.Template) {
	file, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0600))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	err = tpl.Execute(file, context)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("gofmt", "-s", "-w", outputFilePath)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// genTypesDecl generates a "TypeParamDecl" (https://tip.golang.org/ref/spec#Type_parameter_lists) expression,
// used to declare generic types for a type or a function, according to the given element indexes.
func genTypesDecl(indexes []int) string {
	sep := make([]string, len(indexes))
	for index, typeIndex := range indexes {
		sep[index] = fmt.Sprintf("Ty%d any", typeIndex)
	}

	return strings.Join(sep, ", ")
}

// genTypesForward generates a "TypeParamList" (https://tip.golang.org/ref/spec#Type_parameter_lists) expression,
// used to instantiate generic classes, according to the given element indexes.
// Forward refers to forwarding already declared type parameters in order to instantiate the type.
func genTypesForward(indexes []int) string {
	sep := make([]string, len(indexes))
	for index, typeIndex := range indexes {
		sep[index] = fmt.Sprintf("Ty%d", typeIndex)
	}

	return strings.Join(sep, ", ")
}
