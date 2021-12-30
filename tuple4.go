package tuple

import (
	"fmt"
)

// T4 is a tuple type holding 4 generic values.
type T4[Ty1, Ty2, Ty3, Ty4 any] struct {
	V1 Ty1
	V2 Ty2
	V3 Ty3
	V4 Ty4
}

// Len returns the number of values held by the tuple.
func (t T4[Ty1, Ty2, Ty3, Ty4]) Len() int {
	return 4
}

// Values returns the values held by the tuple.
func (t T4[Ty1, Ty2, Ty3, Ty4]) Values() (Ty1, Ty2, Ty3, Ty4) {
	return t.V1, t.V2, t.V3, t.V4
}

// Array returns an array of the tuple values.
func (t T4[Ty1, Ty2, Ty3, Ty4]) Array() [4]any {
	return [4]any{
		t.V1,
		t.V2,
		t.V3,
		t.V4,
	}
}

// Slice returns a slice of the tuple values.
func (t T4[Ty1, Ty2, Ty3, Ty4]) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t T4[Ty1, Ty2, Ty3, Ty4]) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t T4[Ty1, Ty2, Ty3, Ty4]) GoString() string {
	return tupGoString(t.Slice())
}

// New4 creates a new tuple holding 4 generic values.
func New4[Ty1, Ty2, Ty3, Ty4 any](v1 Ty1, v2 Ty2, v3 Ty3, v4 Ty4) T4[Ty1, Ty2, Ty3, Ty4] {
	return T4[Ty1, Ty2, Ty3, Ty4]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
	}
}

// FromArray4 returns a tuple from an array of length 4.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray4[Ty1, Ty2, Ty3, Ty4 any](arr [4]any) (T4[Ty1, Ty2, Ty3, Ty4], error) {
	v1, ok := arr[0].(Ty1)
	if !ok {
		return T4[Ty1, Ty2, Ty3, Ty4]{}, fmt.Errorf("value at array index 0 expected to have type %s but has type %T", typeName[Ty1](), arr[0])
	}
	v2, ok := arr[1].(Ty2)
	if !ok {
		return T4[Ty1, Ty2, Ty3, Ty4]{}, fmt.Errorf("value at array index 1 expected to have type %s but has type %T", typeName[Ty2](), arr[1])
	}
	v3, ok := arr[2].(Ty3)
	if !ok {
		return T4[Ty1, Ty2, Ty3, Ty4]{}, fmt.Errorf("value at array index 2 expected to have type %s but has type %T", typeName[Ty3](), arr[2])
	}
	v4, ok := arr[3].(Ty4)
	if !ok {
		return T4[Ty1, Ty2, Ty3, Ty4]{}, fmt.Errorf("value at array index 3 expected to have type %s but has type %T", typeName[Ty4](), arr[3])
	}

	return New4(v1, v2, v3, v4), nil
}

// FromArray4X returns a tuple from an array of length 4.
// If any of the values can not be converted to the generic type, the function panics.
func FromArray4X[Ty1, Ty2, Ty3, Ty4 any](arr [4]any) T4[Ty1, Ty2, Ty3, Ty4] {
	return FromSlice4X[Ty1, Ty2, Ty3, Ty4](arr[:])
}

// FromSlice4 returns a tuple from a slice of length 4.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice4[Ty1, Ty2, Ty3, Ty4 any](values []any) (T4[Ty1, Ty2, Ty3, Ty4], error) {
	if len(values) != 4 {
		return T4[Ty1, Ty2, Ty3, Ty4]{}, fmt.Errorf("slice length %d must match number of tuple values 4", len(values))
	}

	v1, ok := values[0].(Ty1)
	if !ok {
		return T4[Ty1, Ty2, Ty3, Ty4]{}, fmt.Errorf("value at slice index 0 expected to have type %s but has type %T", typeName[Ty1](), values[0])
	}
	v2, ok := values[1].(Ty2)
	if !ok {
		return T4[Ty1, Ty2, Ty3, Ty4]{}, fmt.Errorf("value at slice index 1 expected to have type %s but has type %T", typeName[Ty2](), values[1])
	}
	v3, ok := values[2].(Ty3)
	if !ok {
		return T4[Ty1, Ty2, Ty3, Ty4]{}, fmt.Errorf("value at slice index 2 expected to have type %s but has type %T", typeName[Ty3](), values[2])
	}
	v4, ok := values[3].(Ty4)
	if !ok {
		return T4[Ty1, Ty2, Ty3, Ty4]{}, fmt.Errorf("value at slice index 3 expected to have type %s but has type %T", typeName[Ty4](), values[3])
	}

	return New4(v1, v2, v3, v4), nil
}

// FromSlice4X returns a tuple from a slice of length 4.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
func FromSlice4X[Ty1, Ty2, Ty3, Ty4 any](values []any) T4[Ty1, Ty2, Ty3, Ty4] {
	if len(values) != 4 {
		panic(fmt.Errorf("slice length %d must match number of tuple values 4", len(values)))
	}

	v1 := values[0].(Ty1)
	v2 := values[1].(Ty2)
	v3 := values[2].(Ty3)
	v4 := values[3].(Ty4)

	return New4(v1, v2, v3, v4)
}
