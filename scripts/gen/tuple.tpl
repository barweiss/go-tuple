package tuple

import (
	"fmt"
)

{{/* $typeName can be used when the context of dot changes. */}}
{{$typeName := .TypeName}}

// T{{.Len}} is a tuple type holding {{.Len}} generic values.
type T{{.Len}}[{{.GenericTypesDecl}}] struct {
	{{range .Indexes -}}
	V{{.}} Ty{{.}}
	{{end -}}
}

// Len returns the number of values held by the tuple.
func (t {{.TypeName}}) Len() int {
	return {{.Len}}
}

// Values returns the values held by the tuple.
func (t {{.TypeName}}) Values() ({{.GenericTypesForward}}) {
	return {{range $index, $num := .Indexes -}}
		{{- if gt $index 0}}, {{end -}}
		t.V{{$num}}
	{{- end}}
}

// Array returns an array of the tuple values.
func (t {{.TypeName}}) Array() [{{.Len}}]any {
	return [{{.Len}}]any{
		{{ range .Indexes -}}
		t.V{{.}},
		{{end}}
	}
}

// Slice returns a slice of the tuple values.
func (t {{.TypeName}}) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t {{.TypeName}}) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t {{.TypeName}}) GoString() string {
	return tupGoString(t.Slice())
}

// New{{.Len}} creates a new tuple holding {{.Len}} generic values.
func New{{.Len}}[{{.GenericTypesDecl}}](
	{{- range $index, $num := .Indexes -}}
	{{- if gt $index 0}}, {{end -}}
	v{{.}} Ty{{.}}
	{{- end -}}
) {{.TypeName}} {
	return {{.TypeName}}{
		{{range .Indexes -}}
		V{{.}}: v{{.}},
		{{end}}
	}
}

// FromArray{{.Len}} returns a tuple from an array of length {{.Len}}.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray{{.Len}}[{{.GenericTypesDecl}}](arr [{{.Len}}]any) ({{.TypeName}}, error) {
	{{range $index, $num := .Indexes -}}
	v{{$num}}, ok := arr[{{$index}}].(Ty{{$num}})
	if !ok {
		return {{$typeName}}{}, fmt.Errorf("value at array index {{$index}} expected to have type %s but has type %T", typeName[Ty{{$num}}](), arr[{{$index}}])
	}
	{{end}}
	return New{{.Len}}(
		{{- range $index, $num := .Indexes -}}
		{{- if gt $index 0}}, {{end -}}
		v{{$num}}
		{{- end -}}
	), nil
}

// FromArray{{.Len}}X returns a tuple from an array of length {{.Len}}.
// If any of the values can not be converted to the generic type, the function panics.
func FromArray{{.Len}}X[{{.GenericTypesDecl}}](arr [{{.Len}}]any) {{.TypeName}} {
	return FromSlice{{.Len}}X[{{.GenericTypesForward}}](arr[:])
}

// FromSlice{{.Len}} returns a tuple from a slice of length {{.Len}}.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice{{.Len}}[{{.GenericTypesDecl}}](values []any) ({{.TypeName}}, error) {
	if len(values) != {{.Len}} {
		return {{.TypeName}}{}, fmt.Errorf("slice length %d must match number of tuple values {{.Len}}", len(values))
	}

	{{range $index, $num := .Indexes -}}
	v{{$num}}, ok := values[{{$index}}].(Ty{{$num}})
	if !ok {
		return {{$typeName}}{}, fmt.Errorf("value at slice index {{$index}} expected to have type %s but has type %T", typeName[Ty{{$num}}](), values[{{$index}}])
	}
	{{end}}
	return New{{.Len}}(
		{{- range $index, $num := .Indexes -}}
		{{- if gt $index 0}}, {{end -}}
		v{{$num}}
		{{- end -}}
	), nil
}

// FromSlice{{.Len}}X returns a tuple from a slice of length {{.Len}}.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
func FromSlice{{.Len}}X[{{.GenericTypesDecl}}](values []any) {{.TypeName}} {
	if len(values) != {{.Len}} {
		panic(fmt.Errorf("slice length %d must match number of tuple values {{.Len}}", len(values)))
	}

	{{range $index, $num := .Indexes -}}
	v{{$num}} := values[{{$index}}].(Ty{{$num}})
	{{end}}
	return New{{.Len}}(
		{{- range $index, $num := .Indexes -}}
		{{- if gt $index 0}}, {{end -}}
		v{{$num}}
		{{- end -}}
	)
}
