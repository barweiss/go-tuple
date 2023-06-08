package tuple

import (
	"encoding/json"
	"fmt"

	"golang.org/x/exp/constraints"
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

// Equal5 returns whether the host tuple is equal to the other tuple.
// All tuple elements of the host and guest parameters must match the "comparable" built-in constraint.
// To test equality of tuples that hold custom Equalable values, use the Equal5E function.
// To test equality of tuples that hold custom Comparable values, use the Equal5C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal5[Ty1, Ty2, Ty3, Ty4, Ty5 comparable](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return host.V1 == guest.V1 && host.V2 == guest.V2 && host.V3 == guest.V3 && host.V4 == guest.V4 && host.V5 == guest.V5
}

// Equal5E returns whether the host tuple is semantically equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Equalable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal5 function.
// To test equality of tuples that hold custom Comparable values, use the Equal5C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal5E[Ty1 Equalable[Ty1], Ty2 Equalable[Ty2], Ty3 Equalable[Ty3], Ty4 Equalable[Ty4], Ty5 Equalable[Ty5]](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return host.V1.Equal(guest.V1) && host.V2.Equal(guest.V2) && host.V3.Equal(guest.V3) && host.V4.Equal(guest.V4) && host.V5.Equal(guest.V5)
}

// Equal5C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal5 function.
// To test equality of tuples that hold custom Equalable values, use the Equal5E function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal5C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5]](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return host.V1.CompareTo(guest.V1).EQ() && host.V2.CompareTo(guest.V2).EQ() && host.V3.CompareTo(guest.V3).EQ() && host.V4.CompareTo(guest.V4).EQ() && host.V5.CompareTo(guest.V5).EQ()
}

// Compare5 returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the Compare5C function.
func Compare5[Ty1, Ty2, Ty3, Ty4, Ty5 constraints.Ordered](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return compareOrdered(host.V1, guest.V1) },

		func() OrderedComparisonResult { return compareOrdered(host.V2, guest.V2) },

		func() OrderedComparisonResult { return compareOrdered(host.V3, guest.V3) },

		func() OrderedComparisonResult { return compareOrdered(host.V4, guest.V4) },

		func() OrderedComparisonResult { return compareOrdered(host.V5, guest.V5) },
	)
}

// Compare5C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the Compare5 function.
func Compare5C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5]](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return host.V1.CompareTo(guest.V1) },

		func() OrderedComparisonResult { return host.V2.CompareTo(guest.V2) },

		func() OrderedComparisonResult { return host.V3.CompareTo(guest.V3) },

		func() OrderedComparisonResult { return host.V4.CompareTo(guest.V4) },

		func() OrderedComparisonResult { return host.V5.CompareTo(guest.V5) },
	)
}

// LessThan5 returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessThan5C function.
func LessThan5[Ty1, Ty2, Ty3, Ty4, Ty5 constraints.Ordered](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return Compare5(host, guest).LT()
}

// LessThan5C returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessThan5 function.
func LessThan5C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5]](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return Compare5C(host, guest).LT()
}

// LessOrEqual5 returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessOrEqual5C function.
func LessOrEqual5[Ty1, Ty2, Ty3, Ty4, Ty5 constraints.Ordered](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return Compare5(host, guest).LE()
}

// LessOrEqual5C returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessOrEqual5 function.
func LessOrEqual5C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5]](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return Compare5C(host, guest).LE()
}

// GreaterThan5 returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterThan5C function.
func GreaterThan5[Ty1, Ty2, Ty3, Ty4, Ty5 constraints.Ordered](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return Compare5(host, guest).GT()
}

// GreaterThan5C returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterThan5 function.
func GreaterThan5C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5]](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return Compare5C(host, guest).GT()
}

// GreaterOrEqual5 returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterOrEqual5C function.
func GreaterOrEqual5[Ty1, Ty2, Ty3, Ty4, Ty5 constraints.Ordered](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return Compare5(host, guest).GE()
}

// GreaterOrEqual5C returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterOrEqual5 function.
func GreaterOrEqual5C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5]](host, guest T5[Ty1, Ty2, Ty3, Ty4, Ty5]) bool {
	return Compare5C(host, guest).GE()
}

// MarshalJSON marshals the tuple into a JSON array.
func (t T5[Ty1, Ty2, Ty3, Ty4, Ty5]) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Slice())
}

// MarshalJSON unmarshals the tuple from a JSON array.
func (t *T5[Ty1, Ty2, Ty3, Ty4, Ty5]) UnmarshalJSON(data []byte) error {
	var slice []any
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	unmarshalled, err := FromSlice5[Ty1, Ty2, Ty3, Ty4, Ty5](slice)
	if err != nil {
		return err
	}

	*t = unmarshalled
	return nil
}
