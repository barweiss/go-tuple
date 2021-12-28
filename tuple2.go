package tuple

import (
	"fmt"
)

// T2 is a tuple type holding 2 generic values.
type T2[Ty1 any, Ty2 any] struct {
	V1 Ty1
	V2 Ty2
}

// Len returns the number of values held by the tuple.
func (t T2[Ty1, Ty2]) Len() int {
	return 2
}

// Values returns the values held by the tuple.
func (t T2[Ty1, Ty2]) Values() (Ty1, Ty2) {
	return t.V1, t.V2
}

// Array returns an array of the tuple values.
func (t T2[Ty1, Ty2]) Array() [2]any {
	return [2]any{
		t.V1,
		t.V2,
	}
}

// Slice returns a slice of the tuple values.
func (t T2[Ty1, Ty2]) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t T2[Ty1, Ty2]) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t T2[Ty1, Ty2]) GoString() string {
	return tupGoString(t.Slice())
}

// New2 creates a new tuple holding 2 generic values.
func New2[Ty1 any, Ty2 any](v1 Ty1, v2 Ty2) T2[Ty1, Ty2] {
	return T2[Ty1, Ty2]{
		V1: v1,
		V2: v2,
	}
}

// FromArray2 returns a tuple from an array of length 2.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray2[Ty1 any, Ty2 any](arr [2]any) (T2[Ty1, Ty2], error) {
	v1, ok := arr[0].(Ty1)
	if !ok {
		return T2[Ty1, Ty2]{}, fmt.Errorf("value at array index 0 expected to have type %s but has type %T", typeName[Ty1](), arr[0])
	}
	v2, ok := arr[1].(Ty2)
	if !ok {
		return T2[Ty1, Ty2]{}, fmt.Errorf("value at array index 1 expected to have type %s but has type %T", typeName[Ty2](), arr[1])
	}

	return New2(v1, v2), nil
}

// FromArray2X returns a tuple from an array of length 2.
// If any of the values can not be converted to the generic type, the function panics.
func FromArray2X[Ty1 any, Ty2 any](arr [2]any) T2[Ty1, Ty2] {
	return FromSlice2X[Ty1, Ty2](arr[:])
}

// FromSlice2 returns a tuple from a slice of length 2.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice2[Ty1 any, Ty2 any](values []any) (T2[Ty1, Ty2], error) {
	if len(values) != 2 {
		return T2[Ty1, Ty2]{}, fmt.Errorf("slice length %d must match number of tuple values 2", len(values))
	}

	v1, ok := values[0].(Ty1)
	if !ok {
		return T2[Ty1, Ty2]{}, fmt.Errorf("value at slice index 0 expected to have type %s but has type %T", typeName[Ty1](), values[0])
	}
	v2, ok := values[1].(Ty2)
	if !ok {
		return T2[Ty1, Ty2]{}, fmt.Errorf("value at slice index 1 expected to have type %s but has type %T", typeName[Ty2](), values[1])
	}

	return New2(v1, v2), nil
}

// FromSlice2X returns a tuple from a slice of length 2.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
func FromSlice2X[Ty1 any, Ty2 any](values []any) T2[Ty1, Ty2] {
	if len(values) != 2 {
		panic(fmt.Errorf("slice length %d must match number of tuple values 2", len(values)))
	}

	v1 := values[0].(Ty1)
	v2 := values[1].(Ty2)

	return New2(v1, v2)
}
