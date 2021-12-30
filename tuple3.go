package tuple

import (
	"fmt"
)

// T3 is a tuple type holding 3 generic values.
type T3[Ty1, Ty2, Ty3 any] struct {
	V1 Ty1
	V2 Ty2
	V3 Ty3
}

// Len returns the number of values held by the tuple.
func (t T3[Ty1, Ty2, Ty3]) Len() int {
	return 3
}

// Values returns the values held by the tuple.
func (t T3[Ty1, Ty2, Ty3]) Values() (Ty1, Ty2, Ty3) {
	return t.V1, t.V2, t.V3
}

// Array returns an array of the tuple values.
func (t T3[Ty1, Ty2, Ty3]) Array() [3]any {
	return [3]any{
		t.V1,
		t.V2,
		t.V3,
	}
}

// Slice returns a slice of the tuple values.
func (t T3[Ty1, Ty2, Ty3]) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t T3[Ty1, Ty2, Ty3]) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t T3[Ty1, Ty2, Ty3]) GoString() string {
	return tupGoString(t.Slice())
}

// New3 creates a new tuple holding 3 generic values.
func New3[Ty1, Ty2, Ty3 any](v1 Ty1, v2 Ty2, v3 Ty3) T3[Ty1, Ty2, Ty3] {
	return T3[Ty1, Ty2, Ty3]{
		V1: v1,
		V2: v2,
		V3: v3,
	}
}

// FromArray3 returns a tuple from an array of length 3.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray3[Ty1, Ty2, Ty3 any](arr [3]any) (T3[Ty1, Ty2, Ty3], error) {
	v1, ok := arr[0].(Ty1)
	if !ok {
		return T3[Ty1, Ty2, Ty3]{}, fmt.Errorf("value at array index 0 expected to have type %s but has type %T", typeName[Ty1](), arr[0])
	}
	v2, ok := arr[1].(Ty2)
	if !ok {
		return T3[Ty1, Ty2, Ty3]{}, fmt.Errorf("value at array index 1 expected to have type %s but has type %T", typeName[Ty2](), arr[1])
	}
	v3, ok := arr[2].(Ty3)
	if !ok {
		return T3[Ty1, Ty2, Ty3]{}, fmt.Errorf("value at array index 2 expected to have type %s but has type %T", typeName[Ty3](), arr[2])
	}

	return New3(v1, v2, v3), nil
}

// FromArray3X returns a tuple from an array of length 3.
// If any of the values can not be converted to the generic type, the function panics.
func FromArray3X[Ty1, Ty2, Ty3 any](arr [3]any) T3[Ty1, Ty2, Ty3] {
	return FromSlice3X[Ty1, Ty2, Ty3](arr[:])
}

// FromSlice3 returns a tuple from a slice of length 3.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice3[Ty1, Ty2, Ty3 any](values []any) (T3[Ty1, Ty2, Ty3], error) {
	if len(values) != 3 {
		return T3[Ty1, Ty2, Ty3]{}, fmt.Errorf("slice length %d must match number of tuple values 3", len(values))
	}

	v1, ok := values[0].(Ty1)
	if !ok {
		return T3[Ty1, Ty2, Ty3]{}, fmt.Errorf("value at slice index 0 expected to have type %s but has type %T", typeName[Ty1](), values[0])
	}
	v2, ok := values[1].(Ty2)
	if !ok {
		return T3[Ty1, Ty2, Ty3]{}, fmt.Errorf("value at slice index 1 expected to have type %s but has type %T", typeName[Ty2](), values[1])
	}
	v3, ok := values[2].(Ty3)
	if !ok {
		return T3[Ty1, Ty2, Ty3]{}, fmt.Errorf("value at slice index 2 expected to have type %s but has type %T", typeName[Ty3](), values[2])
	}

	return New3(v1, v2, v3), nil
}

// FromSlice3X returns a tuple from a slice of length 3.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
func FromSlice3X[Ty1, Ty2, Ty3 any](values []any) T3[Ty1, Ty2, Ty3] {
	if len(values) != 3 {
		panic(fmt.Errorf("slice length %d must match number of tuple values 3", len(values)))
	}

	v1 := values[0].(Ty1)
	v2 := values[1].(Ty2)
	v3 := values[2].(Ty3)

	return New3(v1, v2, v3)
}
