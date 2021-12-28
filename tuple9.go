package tuple

import (
	"fmt"
)

// T9 is a tuple type holding 9 generic values.
type T9[Ty1 any, Ty2 any, Ty3 any, Ty4 any, Ty5 any, Ty6 any, Ty7 any, Ty8 any, Ty9 any] struct {
	V1 Ty1
	V2 Ty2
	V3 Ty3
	V4 Ty4
	V5 Ty5
	V6 Ty6
	V7 Ty7
	V8 Ty8
	V9 Ty9
}

// Len returns the number of values held by the tuple.
func (t T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]) Len() int {
	return 9
}

// Values returns the values held by the tuple.
func (t T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]) Values() (Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9
}

// Array returns an array of the tuple values.
func (t T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]) Array() [9]any {
	return [9]any{
		t.V1,
		t.V2,
		t.V3,
		t.V4,
		t.V5,
		t.V6,
		t.V7,
		t.V8,
		t.V9,
	}
}

// Slice returns a slice of the tuple values.
func (t T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]) GoString() string {
	return tupGoString(t.Slice())
}

// New9 creates a new tuple holding 9 generic values.
func New9[Ty1 any, Ty2 any, Ty3 any, Ty4 any, Ty5 any, Ty6 any, Ty7 any, Ty8 any, Ty9 any](v1 Ty1, v2 Ty2, v3 Ty3, v4 Ty4, v5 Ty5, v6 Ty6, v7 Ty7, v8 Ty8, v9 Ty9) T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9] {
	return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
	}
}

// FromArray9 returns a tuple from an array of length 9.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray9[Ty1 any, Ty2 any, Ty3 any, Ty4 any, Ty5 any, Ty6 any, Ty7 any, Ty8 any, Ty9 any](arr [9]any) (T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9], error) {
	v1, ok := arr[0].(Ty1)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at array index 0 expected to have type %s but has type %T", typeName[Ty1](), arr[0])
	}
	v2, ok := arr[1].(Ty2)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at array index 1 expected to have type %s but has type %T", typeName[Ty2](), arr[1])
	}
	v3, ok := arr[2].(Ty3)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at array index 2 expected to have type %s but has type %T", typeName[Ty3](), arr[2])
	}
	v4, ok := arr[3].(Ty4)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at array index 3 expected to have type %s but has type %T", typeName[Ty4](), arr[3])
	}
	v5, ok := arr[4].(Ty5)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at array index 4 expected to have type %s but has type %T", typeName[Ty5](), arr[4])
	}
	v6, ok := arr[5].(Ty6)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at array index 5 expected to have type %s but has type %T", typeName[Ty6](), arr[5])
	}
	v7, ok := arr[6].(Ty7)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at array index 6 expected to have type %s but has type %T", typeName[Ty7](), arr[6])
	}
	v8, ok := arr[7].(Ty8)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at array index 7 expected to have type %s but has type %T", typeName[Ty8](), arr[7])
	}
	v9, ok := arr[8].(Ty9)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at array index 8 expected to have type %s but has type %T", typeName[Ty9](), arr[8])
	}

	return New9(v1, v2, v3, v4, v5, v6, v7, v8, v9), nil
}

// FromArray9X returns a tuple from an array of length 9.
// If any of the values can not be converted to the generic type, the function panics.
func FromArray9X[Ty1 any, Ty2 any, Ty3 any, Ty4 any, Ty5 any, Ty6 any, Ty7 any, Ty8 any, Ty9 any](arr [9]any) T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9] {
	return FromSlice9X[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9](arr[:])
}

// FromSlice9 returns a tuple from a slice of length 9.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice9[Ty1 any, Ty2 any, Ty3 any, Ty4 any, Ty5 any, Ty6 any, Ty7 any, Ty8 any, Ty9 any](values []any) (T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9], error) {
	if len(values) != 9 {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("slice length %d must match number of tuple values 9", len(values))
	}

	v1, ok := values[0].(Ty1)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at slice index 0 expected to have type %s but has type %T", typeName[Ty1](), values[0])
	}
	v2, ok := values[1].(Ty2)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at slice index 1 expected to have type %s but has type %T", typeName[Ty2](), values[1])
	}
	v3, ok := values[2].(Ty3)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at slice index 2 expected to have type %s but has type %T", typeName[Ty3](), values[2])
	}
	v4, ok := values[3].(Ty4)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at slice index 3 expected to have type %s but has type %T", typeName[Ty4](), values[3])
	}
	v5, ok := values[4].(Ty5)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at slice index 4 expected to have type %s but has type %T", typeName[Ty5](), values[4])
	}
	v6, ok := values[5].(Ty6)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at slice index 5 expected to have type %s but has type %T", typeName[Ty6](), values[5])
	}
	v7, ok := values[6].(Ty7)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at slice index 6 expected to have type %s but has type %T", typeName[Ty7](), values[6])
	}
	v8, ok := values[7].(Ty8)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at slice index 7 expected to have type %s but has type %T", typeName[Ty8](), values[7])
	}
	v9, ok := values[8].(Ty9)
	if !ok {
		return T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9]{}, fmt.Errorf("value at slice index 8 expected to have type %s but has type %T", typeName[Ty9](), values[8])
	}

	return New9(v1, v2, v3, v4, v5, v6, v7, v8, v9), nil
}

// FromSlice9X returns a tuple from a slice of length 9.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
func FromSlice9X[Ty1 any, Ty2 any, Ty3 any, Ty4 any, Ty5 any, Ty6 any, Ty7 any, Ty8 any, Ty9 any](values []any) T9[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7, Ty8, Ty9] {
	if len(values) != 9 {
		panic(fmt.Errorf("slice length %d must match number of tuple values 9", len(values)))
	}

	v1 := values[0].(Ty1)
	v2 := values[1].(Ty2)
	v3 := values[2].(Ty3)
	v4 := values[3].(Ty4)
	v5 := values[4].(Ty5)
	v6 := values[5].(Ty6)
	v7 := values[6].(Ty7)
	v8 := values[7].(Ty8)
	v9 := values[8].(Ty9)

	return New9(v1, v2, v3, v4, v5, v6, v7, v8, v9)
}
