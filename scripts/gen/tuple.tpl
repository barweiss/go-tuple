package tuple

import (
	"encoding/json"
	"fmt"

	"golang.org/x/exp/constraints"
)

{{/* $typeRef can be used when the context of dot changes. */}}
{{$typeRef := typeRef .Indexes}}

// T{{.Len}} is a tuple type holding {{.Len}} generic values.
type T{{.Len}}[{{genericTypesDecl .Indexes "any"}}] struct {
	{{range .Indexes -}}
	V{{.}} Ty{{.}}
	{{end -}}
}

// Len returns the number of values held by the tuple.
func (t {{$typeRef}}) Len() int {
	return {{.Len}}
}

// Values returns the values held by the tuple.
func (t {{$typeRef}}) Values() ({{.GenericTypesForward}}) {
	return {{range $index, $num := .Indexes -}}
		{{- if gt $index 0}}, {{end -}}
		t.V{{$num}}
	{{- end}}
}

// Array returns an array of the tuple values.
func (t {{$typeRef}}) Array() [{{.Len}}]any {
	return [{{.Len}}]any{
		{{ range .Indexes -}}
		t.V{{.}},
		{{end}}
	}
}

// Slice returns a slice of the tuple values.
func (t {{$typeRef}}) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t {{$typeRef}}) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t {{$typeRef}}) GoString() string {
	return tupGoString(t.Slice())
}

// New{{.Len}} creates a new tuple holding {{.Len}} generic values.
func New{{.Len}}[{{genericTypesDecl .Indexes "any"}}](
	{{- range $index, $num := .Indexes -}}
	{{- if gt $index 0}}, {{end -}}
	v{{.}} Ty{{.}}
	{{- end -}}
) {{$typeRef}} {
	return {{$typeRef}}{
		{{range .Indexes -}}
		V{{.}}: v{{.}},
		{{end}}
	}
}

// FromArray{{.Len}} returns a tuple from an array of length {{.Len}}.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray{{.Len}}[{{genericTypesDecl .Indexes "any"}}](arr [{{.Len}}]any) ({{$typeRef}}, error) {
	{{range $index, $num := .Indexes -}}
	v{{$num}}, ok := arr[{{$index}}].(Ty{{$num}})
	if !ok {
		return {{$typeRef}}{}, fmt.Errorf("value at array index {{$index}} expected to have type %s but has type %T", typeName[Ty{{$num}}](), arr[{{$index}}])
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
func FromArray{{.Len}}X[{{genericTypesDecl .Indexes "any"}}](arr [{{.Len}}]any) {{$typeRef}} {
	return FromSlice{{.Len}}X[{{.GenericTypesForward}}](arr[:])
}

// FromSlice{{.Len}} returns a tuple from a slice of length {{.Len}}.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice{{.Len}}[{{genericTypesDecl .Indexes "any"}}](values []any) ({{$typeRef}}, error) {
	if len(values) != {{.Len}} {
		return {{$typeRef}}{}, fmt.Errorf("slice length %d must match number of tuple values {{.Len}}", len(values))
	}

	{{range $index, $num := .Indexes -}}
	v{{$num}}, ok := values[{{$index}}].(Ty{{$num}})
	if !ok {
		return {{$typeRef}}{}, fmt.Errorf("value at slice index {{$index}} expected to have type %s but has type %T", typeName[Ty{{$num}}](), values[{{$index}}])
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
func FromSlice{{.Len}}X[{{genericTypesDecl .Indexes "any"}}](values []any) {{$typeRef}} {
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

// Equal{{.Len}} returns whether the host tuple is equal to the other tuple.
// All tuple elements of the host and guest parameters must match the "comparable" built-in constraint.
// To test equality of tuples that hold custom Equalable values, use the Equal{{.Len}}E function.
// To test equality of tuples that hold custom Comparable values, use the Equal{{.Len}}C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal{{.Len}}[{{genericTypesDecl .Indexes "comparable"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return {{range $index, $num := .Indexes}}{{if gt $index 0}} && {{end}}host.V{{$num}} == guest.V{{$num}}{{end}}
}

// Equal{{.Len}}E returns whether the host tuple is semantically equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Equalable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal{{.Len}} function.
// To test equality of tuples that hold custom Comparable values, use the Equal{{.Len}}C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal{{.Len}}E[{{genericTypesDeclGenericConstraint .Indexes "Equalable"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return {{range $index, $num := .Indexes}}{{if gt $index 0}} && {{end}}host.V{{$num}}.Equal(guest.V{{$num}}){{end}}
}

// Equal{{.Len}}C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal{{.Len}} function.
// To test equality of tuples that hold custom Equalable values, use the Equal{{.Len}}E function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal{{.Len}}C[{{genericTypesDeclGenericConstraint .Indexes "Comparable"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return {{range $index, $num := .Indexes}}{{if gt $index 0}} && {{end}}host.V{{$num}}.CompareTo(guest.V{{$num}}).EQ(){{end}}
}

// Compare{{.Len}} returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the Compare{{.Len}}C function.
func Compare{{.Len}}[{{genericTypesDecl .Indexes "constraints.Ordered"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) OrderedComparisonResult {
	return multiCompare({{range .Indexes}}
		func () OrderedComparisonResult { return compareOrdered(host.V{{.}}, guest.V{{.}}) },
	{{end}})
}

// Compare{{.Len}}C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the Compare{{.Len}} function.
func Compare{{.Len}}C[{{genericTypesDeclGenericConstraint .Indexes "Comparable"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) OrderedComparisonResult {
	return multiCompare({{range .Indexes}}
		func () OrderedComparisonResult { return host.V{{.}}.CompareTo(guest.V{{.}}) },
	{{end}})
}

// LessThan{{.Len}} returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessThan{{.Len}}C function.
func LessThan{{.Len}}[{{genericTypesDecl .Indexes "constraints.Ordered"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return Compare{{.Len}}(host, guest).LT()
}

// LessThan{{.Len}}C returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessThan{{.Len}} function.
func LessThan{{.Len}}C[{{genericTypesDeclGenericConstraint .Indexes "Comparable"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return Compare{{.Len}}C(host, guest).LT()
}

// LessOrEqual{{.Len}} returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessOrEqual{{.Len}}C function.
func LessOrEqual{{.Len}}[{{genericTypesDecl .Indexes "constraints.Ordered"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return Compare{{.Len}}(host, guest).LE()
}

// LessOrEqual{{.Len}}C returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessOrEqual{{.Len}} function.
func LessOrEqual{{.Len}}C[{{genericTypesDeclGenericConstraint .Indexes "Comparable"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return Compare{{.Len}}C(host, guest).LE()
}

// GreaterThan{{.Len}} returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterThan{{.Len}}C function.
func GreaterThan{{.Len}}[{{genericTypesDecl .Indexes "constraints.Ordered"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return Compare{{.Len}}(host, guest).GT()
}

// GreaterThan{{.Len}}C returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterThan{{.Len}} function.
func GreaterThan{{.Len}}C[{{genericTypesDeclGenericConstraint .Indexes "Comparable"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return Compare{{.Len}}C(host, guest).GT()
}

// GreaterOrEqual{{.Len}} returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterOrEqual{{.Len}}C function.
func GreaterOrEqual{{.Len}}[{{genericTypesDecl .Indexes "constraints.Ordered"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return Compare{{.Len}}(host, guest).GE()
}

// GreaterOrEqual{{.Len}}C returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterOrEqual{{.Len}} function.
func GreaterOrEqual{{.Len}}C[{{genericTypesDeclGenericConstraint .Indexes "Comparable"}}](host, guest T{{.Len}}[{{.GenericTypesForward}}]) bool {
	return Compare{{.Len}}C(host, guest).GE()
}

// MarshalJSON marshals the tuple into a JSON array.
func (t {{$typeRef}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Slice())
}

// MarshalJSON unmarshals the tuple from a JSON array.
func (t *{{$typeRef}}) UnmarshalJSON(data []byte) error {
	// Working with json.RawMessage instead of any enables custom struct support.
	var slice []json.RawMessage
	if err := json.Unmarshal(data, &slice); err != nil {
		return fmt.Errorf("unable to unmarshal json array for tuple: %w", err)
	}

	if len(slice) != {{.Len}} {
		return fmt.Errorf("unmarshalled json array length %d must match number of tuple values {{.Len}}", len(slice))
	}

	{{- range $index, $num := .Indexes}}
	if err := json.Unmarshal(slice[{{$index}}], &t.V{{.}}); err != nil {
		return fmt.Errorf("value %q at slice index {{$index}} failed to unmarshal: %w", string(slice[{{$index}}]), err)
	}
	{{end -}}

	return nil
}
