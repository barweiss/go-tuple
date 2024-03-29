package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"text/template"
)

// templateContext is the context passed to the template engine for generating tuple code and test files.
type templateContext struct {
	Indexes             []int
	Len                 int
	GenericTypesForward string
}

const minTupleLength = 1
const maxTupleLength = 9

var funcMap = template.FuncMap{
	"quote": func(value interface{}) string {
		return strconv.Quote(fmt.Sprint(value))
	},
	"inc": func(value int) int {
		return value + 1
	},
	"typeRef": func(indexes []int, suffix ...string) string {
		if len(suffix) > 1 {
			panic(fmt.Errorf("typeRef accepts at most 1 suffix argument"))
		}

		var typeNameSuffix string
		if len(suffix) == 1 {
			typeNameSuffix = suffix[0]
		}

		return fmt.Sprintf("T%d%s[%s]", len(indexes), typeNameSuffix, genTypesForward(indexes))
	},
	"genericTypesDecl":                  genTypesDecl,
	"genericTypesDeclGenericConstraint": genTypesDeclGenericConstraint,
	"buildSingleTypedOverload": func(indexes []int, typ string) string {
		typesArray := make([]string, 0, len(indexes))
		for range indexes {
			typesArray = append(typesArray, typ)
		}

		return fmt.Sprintf("T%d[%s]", len(indexes), strings.Join(typesArray, ", "))
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

		context := templateContext{
			Indexes:             indexes,
			Len:                 tupleLength,
			GenericTypesForward: genTypesForward(indexes),
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

func genTypesDeclGenericConstraint(indexes []int, constraint string) string {
	sep := make([]string, len(indexes))
	for index, typeIndex := range indexes {
		typ := fmt.Sprintf("Ty%d", typeIndex)
		sep[index] = fmt.Sprintf("%s %s[%s]", typ, constraint, typ)
	}

	return strings.Join(sep, ", ")
}

// genTypesDecl generates a "TypeParamDecl" (https://tip.golang.org/ref/spec#Type_parameter_lists) expression,
// used to declare generic types for a type or a function, according to the given element indexes.
func genTypesDecl(indexes []int, constraint string) string {
	sep := make([]string, len(indexes))
	for index, typeIndex := range indexes {
		sep[index] = fmt.Sprintf("Ty%d", typeIndex)
	}

	return strings.Join(sep, ", ") + " " + constraint
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
