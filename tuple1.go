package tuple

import (
	"fmt"
)

// T1 is a tuple type holding 1 generic values.
type T1[Ty1 any] struct {
	V1 Ty1
}

// Len returns the number of values held by the tuple.
func (t T1[Ty1]) Len() int {
	return 1
}

// Values returns the values held by the tuple.
func (t T1[Ty1]) Values() Ty1 {
	return t.V1
}

// Array returns an array of the tuple values.
func (t T1[Ty1]) Array() [1]any {
	return [1]any{
		t.V1,
	}
}

// Slice returns a slice of the tuple values.
func (t T1[Ty1]) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t T1[Ty1]) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t T1[Ty1]) GoString() string {
	return tupGoString(t.Slice())
}

// New1 creates a new tuple holding 1 generic values.
func New1[Ty1 any](v1 Ty1) T1[Ty1] {
	return T1[Ty1]{
		V1: v1,
	}
}

// FromArray1 returns a tuple from an array of length 1.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray1[Ty1 any](arr [1]any) (T1[Ty1], error) {
	v1, ok := arr[0].(Ty1)
	if !ok {
		return T1[Ty1]{}, fmt.Errorf("value at array index 0 expected to have type %s but has type %T", typeName[Ty1](), arr[0])
	}

	return New1(v1), nil
}

// FromArray1X returns a tuple from an array of length 1.
// If any of the values can not be converted to the generic type, the function panics.
func FromArray1X[Ty1 any](arr [1]any) T1[Ty1] {
	return FromSlice1X[Ty1](arr[:])
}

// FromSlice1 returns a tuple from a slice of length 1.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice1[Ty1 any](values []any) (T1[Ty1], error) {
	if len(values) != 1 {
		return T1[Ty1]{}, fmt.Errorf("slice length %d must match number of tuple values 1", len(values))
	}

	v1, ok := values[0].(Ty1)
	if !ok {
		return T1[Ty1]{}, fmt.Errorf("value at slice index 0 expected to have type %s but has type %T", typeName[Ty1](), values[0])
	}

	return New1(v1), nil
}

// FromSlice1X returns a tuple from a slice of length 1.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
func FromSlice1X[Ty1 any](values []any) T1[Ty1] {
	if len(values) != 1 {
		panic(fmt.Errorf("slice length %d must match number of tuple values 1", len(values)))
	}

	v1 := values[0].(Ty1)

	return New1(v1)
}
