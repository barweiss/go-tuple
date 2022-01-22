package tuple

import (
	"constraints"
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

// Equal3 returns whether the host tuple is equal to the other tuple.
// All tuple elements of the host and guest parameters must match the "comparable" built-in constraint.
// To test equality of tuples that hold custom Equalable values, use the Equal3E function.
// To test equality of tuples that hold custom Comparable values, use the Equal3C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal3[Ty1, Ty2, Ty3 comparable](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return host.V1 == guest.V1 && host.V2 == guest.V2 && host.V3 == guest.V3
}

// Equal3E returns whether the host tuple is semantically equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Equalable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal3 function.
// To test equality of tuples that hold custom Comparable values, use the Equal3C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal3E[Ty1 Equalable[Ty1], Ty2 Equalable[Ty2], Ty3 Equalable[Ty3]](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return host.V1.Equal(guest.V1) && host.V2.Equal(guest.V2) && host.V3.Equal(guest.V3)
}

// Equal3C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal3 function.
// To test equality of tuples that hold custom Equalable values, use the Equal3E function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal3C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3]](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return host.V1.CompareTo(guest.V1).EQ() && host.V2.CompareTo(guest.V2).EQ() && host.V3.CompareTo(guest.V3).EQ()
}

// Compare3 returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the Compare3C function.
func Compare3[Ty1, Ty2, Ty3 constraints.Ordered](host, guest T3[Ty1, Ty2, Ty3]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return compareOrdered(host.V1, guest.V1) },

		func() OrderedComparisonResult { return compareOrdered(host.V2, guest.V2) },

		func() OrderedComparisonResult { return compareOrdered(host.V3, guest.V3) },
	)
}

// Compare3C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the Compare3 function.
func Compare3C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3]](host, guest T3[Ty1, Ty2, Ty3]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return host.V1.CompareTo(guest.V1) },

		func() OrderedComparisonResult { return host.V2.CompareTo(guest.V2) },

		func() OrderedComparisonResult { return host.V3.CompareTo(guest.V3) },
	)
}

// LessThan3 returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessThan3C function.
func LessThan3[Ty1, Ty2, Ty3 constraints.Ordered](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return Compare3(host, guest).LT()
}

// LessThan3C returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessThan3 function.
func LessThan3C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3]](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return Compare3C(host, guest).LT()
}

// LessOrEqual3 returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessOrEqual3C function.
func LessOrEqual3[Ty1, Ty2, Ty3 constraints.Ordered](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return Compare3(host, guest).LE()
}

// LessOrEqual3C returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessOrEqual3 function.
func LessOrEqual3C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3]](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return Compare3C(host, guest).LE()
}

// GreaterThan3 returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterThan3C function.
func GreaterThan3[Ty1, Ty2, Ty3 constraints.Ordered](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return Compare3(host, guest).GT()
}

// GreaterThan3C returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterThan3 function.
func GreaterThan3C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3]](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return Compare3C(host, guest).GT()
}

// GreaterOrEqual3 returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterOrEqual3C function.
func GreaterOrEqual3[Ty1, Ty2, Ty3 constraints.Ordered](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return Compare3(host, guest).GE()
}

// GreaterOrEqual3C returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterOrEqual3 function.
func GreaterOrEqual3C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3]](host, guest T3[Ty1, Ty2, Ty3]) bool {
	return Compare3C(host, guest).GE()
}
