package tuple

import (
	"constraints"
	"encoding/json"
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

// Equal4 returns whether the host tuple is equal to the other tuple.
// All tuple elements of the host and guest parameters must match the "comparable" built-in constraint.
// To test equality of tuples that hold custom Equalable values, use the Equal4E function.
// To test equality of tuples that hold custom Comparable values, use the Equal4C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal4[Ty1, Ty2, Ty3, Ty4 comparable](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return host.V1 == guest.V1 && host.V2 == guest.V2 && host.V3 == guest.V3 && host.V4 == guest.V4
}

// Equal4E returns whether the host tuple is semantically equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Equalable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal4 function.
// To test equality of tuples that hold custom Comparable values, use the Equal4C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal4E[Ty1 Equalable[Ty1], Ty2 Equalable[Ty2], Ty3 Equalable[Ty3], Ty4 Equalable[Ty4]](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return host.V1.Equal(guest.V1) && host.V2.Equal(guest.V2) && host.V3.Equal(guest.V3) && host.V4.Equal(guest.V4)
}

// Equal4C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal4 function.
// To test equality of tuples that hold custom Equalable values, use the Equal4E function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal4C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4]](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return host.V1.CompareTo(guest.V1).EQ() && host.V2.CompareTo(guest.V2).EQ() && host.V3.CompareTo(guest.V3).EQ() && host.V4.CompareTo(guest.V4).EQ()
}

// Compare4 returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the Compare4C function.
func Compare4[Ty1, Ty2, Ty3, Ty4 constraints.Ordered](host, guest T4[Ty1, Ty2, Ty3, Ty4]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return compareOrdered(host.V1, guest.V1) },

		func() OrderedComparisonResult { return compareOrdered(host.V2, guest.V2) },

		func() OrderedComparisonResult { return compareOrdered(host.V3, guest.V3) },

		func() OrderedComparisonResult { return compareOrdered(host.V4, guest.V4) },
	)
}

// Compare4C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the Compare4 function.
func Compare4C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4]](host, guest T4[Ty1, Ty2, Ty3, Ty4]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return host.V1.CompareTo(guest.V1) },

		func() OrderedComparisonResult { return host.V2.CompareTo(guest.V2) },

		func() OrderedComparisonResult { return host.V3.CompareTo(guest.V3) },

		func() OrderedComparisonResult { return host.V4.CompareTo(guest.V4) },
	)
}

// LessThan4 returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessThan4C function.
func LessThan4[Ty1, Ty2, Ty3, Ty4 constraints.Ordered](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return Compare4(host, guest).LT()
}

// LessThan4C returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessThan4 function.
func LessThan4C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4]](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return Compare4C(host, guest).LT()
}

// LessOrEqual4 returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessOrEqual4C function.
func LessOrEqual4[Ty1, Ty2, Ty3, Ty4 constraints.Ordered](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return Compare4(host, guest).LE()
}

// LessOrEqual4C returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessOrEqual4 function.
func LessOrEqual4C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4]](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return Compare4C(host, guest).LE()
}

// GreaterThan4 returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterThan4C function.
func GreaterThan4[Ty1, Ty2, Ty3, Ty4 constraints.Ordered](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return Compare4(host, guest).GT()
}

// GreaterThan4C returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterThan4 function.
func GreaterThan4C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4]](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return Compare4C(host, guest).GT()
}

// GreaterOrEqual4 returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterOrEqual4C function.
func GreaterOrEqual4[Ty1, Ty2, Ty3, Ty4 constraints.Ordered](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return Compare4(host, guest).GE()
}

// GreaterOrEqual4C returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterOrEqual4 function.
func GreaterOrEqual4C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4]](host, guest T4[Ty1, Ty2, Ty3, Ty4]) bool {
	return Compare4C(host, guest).GE()
}

func (t T4[Ty1, Ty2, Ty3, Ty4]) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Slice())
}

func (t *T4[Ty1, Ty2, Ty3, Ty4]) UnmarshalJSON(data []byte) error {
	var slice []any
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	unmarshalled, err := FromSlice4[Ty1, Ty2, Ty3, Ty4](slice)
	if err != nil {
		return err
	}

	*t = unmarshalled
	return nil
}
