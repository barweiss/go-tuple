package tuple

import (
	"fmt"
)

// T5 is a tuple type holding 5 generic values.
type T5[Ty1, Ty2, Ty3, Ty4, Ty5 any] struct {
	V1 Ty1
	V2 Ty2
	V3 Ty3
	V4 Ty4
	V5 Ty5
}

// Len returns the number of values held by the tuple.
func (t T5[Ty1, Ty2, Ty3, Ty4, Ty5]) Len() int {
	return 5
}

// Values returns the values held by the tuple.
func (t T5[Ty1, Ty2, Ty3, Ty4, Ty5]) Values() (Ty1, Ty2, Ty3, Ty4, Ty5) {
	return t.V1, t.V2, t.V3, t.V4, t.V5
}

// Array returns an array of the tuple values.
func (t T5[Ty1, Ty2, Ty3, Ty4, Ty5]) Array() [5]any {
	return [5]any{
		t.V1,
		t.V2,
		t.V3,
		t.V4,
		t.V5,
	}
}

// Slice returns a slice of the tuple values.
func (t T5[Ty1, Ty2, Ty3, Ty4, Ty5]) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t T5[Ty1, Ty2, Ty3, Ty4, Ty5]) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t T5[Ty1, Ty2, Ty3, Ty4, Ty5]) GoString() string {
	return tupGoString(t.Slice())
}

// New5 creates a new tuple holding 5 generic values.
func New5[Ty1, Ty2, Ty3, Ty4, Ty5 any](v1 Ty1, v2 Ty2, v3 Ty3, v4 Ty4, v5 Ty5) T5[Ty1, Ty2, Ty3, Ty4, Ty5] {
	return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
	}
}

// FromArray5 returns a tuple from an array of length 5.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray5[Ty1, Ty2, Ty3, Ty4, Ty5 any](arr [5]any) (T5[Ty1, Ty2, Ty3, Ty4, Ty5], error) {
	v1, ok := arr[0].(Ty1)
	if !ok {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("value at array index 0 expected to have type %s but has type %T", typeName[Ty1](), arr[0])
	}
	v2, ok := arr[1].(Ty2)
	if !ok {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("value at array index 1 expected to have type %s but has type %T", typeName[Ty2](), arr[1])
	}
	v3, ok := arr[2].(Ty3)
	if !ok {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("value at array index 2 expected to have type %s but has type %T", typeName[Ty3](), arr[2])
	}
	v4, ok := arr[3].(Ty4)
	if !ok {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("value at array index 3 expected to have type %s but has type %T", typeName[Ty4](), arr[3])
	}
	v5, ok := arr[4].(Ty5)
	if !ok {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("value at array index 4 expected to have type %s but has type %T", typeName[Ty5](), arr[4])
	}

	return New5(v1, v2, v3, v4, v5), nil
}

// FromArray5X returns a tuple from an array of length 5.
// If any of the values can not be converted to the generic type, the function panics.
func FromArray5X[Ty1, Ty2, Ty3, Ty4, Ty5 any](arr [5]any) T5[Ty1, Ty2, Ty3, Ty4, Ty5] {
	return FromSlice5X[Ty1, Ty2, Ty3, Ty4, Ty5](arr[:])
}

// FromSlice5 returns a tuple from a slice of length 5.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice5[Ty1, Ty2, Ty3, Ty4, Ty5 any](values []any) (T5[Ty1, Ty2, Ty3, Ty4, Ty5], error) {
	if len(values) != 5 {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("slice length %d must match number of tuple values 5", len(values))
	}

	v1, ok := values[0].(Ty1)
	if !ok {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("value at slice index 0 expected to have type %s but has type %T", typeName[Ty1](), values[0])
	}
	v2, ok := values[1].(Ty2)
	if !ok {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("value at slice index 1 expected to have type %s but has type %T", typeName[Ty2](), values[1])
	}
	v3, ok := values[2].(Ty3)
	if !ok {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("value at slice index 2 expected to have type %s but has type %T", typeName[Ty3](), values[2])
	}
	v4, ok := values[3].(Ty4)
	if !ok {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("value at slice index 3 expected to have type %s but has type %T", typeName[Ty4](), values[3])
	}
	v5, ok := values[4].(Ty5)
	if !ok {
		return T5[Ty1, Ty2, Ty3, Ty4, Ty5]{}, fmt.Errorf("value at slice index 4 expected to have type %s but has type %T", typeName[Ty5](), values[4])
	}

	return New5(v1, v2, v3, v4, v5), nil
}

// FromSlice5X returns a tuple from a slice of length 5.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
func FromSlice5X[Ty1, Ty2, Ty3, Ty4, Ty5 any](values []any) T5[Ty1, Ty2, Ty3, Ty4, Ty5] {
	if len(values) != 5 {
		panic(fmt.Errorf("slice length %d must match number of tuple values 5", len(values)))
	}

	v1 := values[0].(Ty1)
	v2 := values[1].(Ty2)
	v3 := values[2].(Ty3)
	v4 := values[3].(Ty4)
	v5 := values[4].(Ty5)

	return New5(v1, v2, v3, v4, v5)
}
