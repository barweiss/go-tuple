package tuple

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// T2 is a tuple type holding 2 generic values.
type T2[Ty1, Ty2 any] struct {
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
func New2[Ty1, Ty2 any](v1 Ty1, v2 Ty2) T2[Ty1, Ty2] {
	return T2[Ty1, Ty2]{
		V1: v1,
		V2: v2,
	}
}

// FromArray2 returns a tuple from an array of length 2.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray2[Ty1, Ty2 any](arr [2]any) (T2[Ty1, Ty2], error) {
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
func FromArray2X[Ty1, Ty2 any](arr [2]any) T2[Ty1, Ty2] {
	return FromSlice2X[Ty1, Ty2](arr[:])
}

// FromSlice2 returns a tuple from a slice of length 2.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice2[Ty1, Ty2 any](values []any) (T2[Ty1, Ty2], error) {
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
func FromSlice2X[Ty1, Ty2 any](values []any) T2[Ty1, Ty2] {
	if len(values) != 2 {
		panic(fmt.Errorf("slice length %d must match number of tuple values 2", len(values)))
	}

	v1 := values[0].(Ty1)
	v2 := values[1].(Ty2)

	return New2(v1, v2)
}

// Equal2 returns whether the host tuple is equal to the other tuple.
// All tuple elements of the host and guest parameters must match the "comparable" built-in constraint.
// To test equality of tuples that hold custom Equalable values, use the Equal2E function.
// To test equality of tuples that hold custom Comparable values, use the Equal2C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal2[Ty1, Ty2 comparable](host, guest T2[Ty1, Ty2]) bool {
	return host.V1 == guest.V1 && host.V2 == guest.V2
}

// Equal2E returns whether the host tuple is semantically equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Equalable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal2 function.
// To test equality of tuples that hold custom Comparable values, use the Equal2C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal2E[Ty1 Equalable[Ty1], Ty2 Equalable[Ty2]](host, guest T2[Ty1, Ty2]) bool {
	return host.V1.Equal(guest.V1) && host.V2.Equal(guest.V2)
}

// Equal2C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal2 function.
// To test equality of tuples that hold custom Equalable values, use the Equal2E function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal2C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2]](host, guest T2[Ty1, Ty2]) bool {
	return host.V1.CompareTo(guest.V1).EQ() && host.V2.CompareTo(guest.V2).EQ()
}

// Compare2 returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the Compare2C function.
func Compare2[Ty1, Ty2 constraints.Ordered](host, guest T2[Ty1, Ty2]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return compareOrdered(host.V1, guest.V1) },

		func() OrderedComparisonResult { return compareOrdered(host.V2, guest.V2) },
	)
}

// Compare2C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the Compare2 function.
func Compare2C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2]](host, guest T2[Ty1, Ty2]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return host.V1.CompareTo(guest.V1) },

		func() OrderedComparisonResult { return host.V2.CompareTo(guest.V2) },
	)
}

// LessThan2 returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessThan2C function.
func LessThan2[Ty1, Ty2 constraints.Ordered](host, guest T2[Ty1, Ty2]) bool {
	return Compare2(host, guest).LT()
}

// LessThan2C returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessThan2 function.
func LessThan2C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2]](host, guest T2[Ty1, Ty2]) bool {
	return Compare2C(host, guest).LT()
}

// LessOrEqual2 returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessOrEqual2C function.
func LessOrEqual2[Ty1, Ty2 constraints.Ordered](host, guest T2[Ty1, Ty2]) bool {
	return Compare2(host, guest).LE()
}

// LessOrEqual2C returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessOrEqual2 function.
func LessOrEqual2C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2]](host, guest T2[Ty1, Ty2]) bool {
	return Compare2C(host, guest).LE()
}

// GreaterThan2 returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterThan2C function.
func GreaterThan2[Ty1, Ty2 constraints.Ordered](host, guest T2[Ty1, Ty2]) bool {
	return Compare2(host, guest).GT()
}

// GreaterThan2C returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterThan2 function.
func GreaterThan2C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2]](host, guest T2[Ty1, Ty2]) bool {
	return Compare2C(host, guest).GT()
}

// GreaterOrEqual2 returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterOrEqual2C function.
func GreaterOrEqual2[Ty1, Ty2 constraints.Ordered](host, guest T2[Ty1, Ty2]) bool {
	return Compare2(host, guest).GE()
}

// GreaterOrEqual2C returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterOrEqual2 function.
func GreaterOrEqual2C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2]](host, guest T2[Ty1, Ty2]) bool {
	return Compare2C(host, guest).GE()
}
