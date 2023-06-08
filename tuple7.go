package tuple

import (
	"constraints"
	"encoding/json"
	"fmt"
)

// T7 is a tuple type holding 7 generic values.
type T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 any] struct {
	V1 Ty1
	V2 Ty2
	V3 Ty3
	V4 Ty4
	V5 Ty5
	V6 Ty6
	V7 Ty7
}

// Len returns the number of values held by the tuple.
func (t T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) Len() int {
	return 7
}

// Values returns the values held by the tuple.
func (t T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) Values() (Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7
}

// Array returns an array of the tuple values.
func (t T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) Array() [7]any {
	return [7]any{
		t.V1,
		t.V2,
		t.V3,
		t.V4,
		t.V5,
		t.V6,
		t.V7,
	}
}

// Slice returns a slice of the tuple values.
func (t T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) GoString() string {
	return tupGoString(t.Slice())
}

// New7 creates a new tuple holding 7 generic values.
func New7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 any](v1 Ty1, v2 Ty2, v3 Ty3, v4 Ty4, v5 Ty5, v6 Ty6, v7 Ty7) T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7] {
	return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
	}
}

// FromArray7 returns a tuple from an array of length 7.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 any](arr [7]any) (T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7], error) {
	v1, ok := arr[0].(Ty1)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at array index 0 expected to have type %s but has type %T", typeName[Ty1](), arr[0])
	}
	v2, ok := arr[1].(Ty2)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at array index 1 expected to have type %s but has type %T", typeName[Ty2](), arr[1])
	}
	v3, ok := arr[2].(Ty3)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at array index 2 expected to have type %s but has type %T", typeName[Ty3](), arr[2])
	}
	v4, ok := arr[3].(Ty4)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at array index 3 expected to have type %s but has type %T", typeName[Ty4](), arr[3])
	}
	v5, ok := arr[4].(Ty5)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at array index 4 expected to have type %s but has type %T", typeName[Ty5](), arr[4])
	}
	v6, ok := arr[5].(Ty6)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at array index 5 expected to have type %s but has type %T", typeName[Ty6](), arr[5])
	}
	v7, ok := arr[6].(Ty7)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at array index 6 expected to have type %s but has type %T", typeName[Ty7](), arr[6])
	}

	return New7(v1, v2, v3, v4, v5, v6, v7), nil
}

// FromArray7X returns a tuple from an array of length 7.
// If any of the values can not be converted to the generic type, the function panics.
func FromArray7X[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 any](arr [7]any) T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7] {
	return FromSlice7X[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7](arr[:])
}

// FromSlice7 returns a tuple from a slice of length 7.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 any](values []any) (T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7], error) {
	if len(values) != 7 {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("slice length %d must match number of tuple values 7", len(values))
	}

	v1, ok := values[0].(Ty1)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at slice index 0 expected to have type %s but has type %T", typeName[Ty1](), values[0])
	}
	v2, ok := values[1].(Ty2)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at slice index 1 expected to have type %s but has type %T", typeName[Ty2](), values[1])
	}
	v3, ok := values[2].(Ty3)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at slice index 2 expected to have type %s but has type %T", typeName[Ty3](), values[2])
	}
	v4, ok := values[3].(Ty4)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at slice index 3 expected to have type %s but has type %T", typeName[Ty4](), values[3])
	}
	v5, ok := values[4].(Ty5)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at slice index 4 expected to have type %s but has type %T", typeName[Ty5](), values[4])
	}
	v6, ok := values[5].(Ty6)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at slice index 5 expected to have type %s but has type %T", typeName[Ty6](), values[5])
	}
	v7, ok := values[6].(Ty7)
	if !ok {
		return T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]{}, fmt.Errorf("value at slice index 6 expected to have type %s but has type %T", typeName[Ty7](), values[6])
	}

	return New7(v1, v2, v3, v4, v5, v6, v7), nil
}

// FromSlice7X returns a tuple from a slice of length 7.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
func FromSlice7X[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 any](values []any) T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7] {
	if len(values) != 7 {
		panic(fmt.Errorf("slice length %d must match number of tuple values 7", len(values)))
	}

	v1 := values[0].(Ty1)
	v2 := values[1].(Ty2)
	v3 := values[2].(Ty3)
	v4 := values[3].(Ty4)
	v5 := values[4].(Ty5)
	v6 := values[5].(Ty6)
	v7 := values[6].(Ty7)

	return New7(v1, v2, v3, v4, v5, v6, v7)
}

// Equal7 returns whether the host tuple is equal to the other tuple.
// All tuple elements of the host and guest parameters must match the "comparable" built-in constraint.
// To test equality of tuples that hold custom Equalable values, use the Equal7E function.
// To test equality of tuples that hold custom Comparable values, use the Equal7C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 comparable](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return host.V1 == guest.V1 && host.V2 == guest.V2 && host.V3 == guest.V3 && host.V4 == guest.V4 && host.V5 == guest.V5 && host.V6 == guest.V6 && host.V7 == guest.V7
}

// Equal7E returns whether the host tuple is semantically equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Equalable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal7 function.
// To test equality of tuples that hold custom Comparable values, use the Equal7C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal7E[Ty1 Equalable[Ty1], Ty2 Equalable[Ty2], Ty3 Equalable[Ty3], Ty4 Equalable[Ty4], Ty5 Equalable[Ty5], Ty6 Equalable[Ty6], Ty7 Equalable[Ty7]](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return host.V1.Equal(guest.V1) && host.V2.Equal(guest.V2) && host.V3.Equal(guest.V3) && host.V4.Equal(guest.V4) && host.V5.Equal(guest.V5) && host.V6.Equal(guest.V6) && host.V7.Equal(guest.V7)
}

// Equal7C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal7 function.
// To test equality of tuples that hold custom Equalable values, use the Equal7E function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal7C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6], Ty7 Comparable[Ty7]](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return host.V1.CompareTo(guest.V1).EQ() && host.V2.CompareTo(guest.V2).EQ() && host.V3.CompareTo(guest.V3).EQ() && host.V4.CompareTo(guest.V4).EQ() && host.V5.CompareTo(guest.V5).EQ() && host.V6.CompareTo(guest.V6).EQ() && host.V7.CompareTo(guest.V7).EQ()
}

// Compare7 returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the Compare7C function.
func Compare7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 constraints.Ordered](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return compareOrdered(host.V1, guest.V1) },

		func() OrderedComparisonResult { return compareOrdered(host.V2, guest.V2) },

		func() OrderedComparisonResult { return compareOrdered(host.V3, guest.V3) },

		func() OrderedComparisonResult { return compareOrdered(host.V4, guest.V4) },

		func() OrderedComparisonResult { return compareOrdered(host.V5, guest.V5) },

		func() OrderedComparisonResult { return compareOrdered(host.V6, guest.V6) },

		func() OrderedComparisonResult { return compareOrdered(host.V7, guest.V7) },
	)
}

// Compare7C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the Compare7 function.
func Compare7C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6], Ty7 Comparable[Ty7]](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return host.V1.CompareTo(guest.V1) },

		func() OrderedComparisonResult { return host.V2.CompareTo(guest.V2) },

		func() OrderedComparisonResult { return host.V3.CompareTo(guest.V3) },

		func() OrderedComparisonResult { return host.V4.CompareTo(guest.V4) },

		func() OrderedComparisonResult { return host.V5.CompareTo(guest.V5) },

		func() OrderedComparisonResult { return host.V6.CompareTo(guest.V6) },

		func() OrderedComparisonResult { return host.V7.CompareTo(guest.V7) },
	)
}

// LessThan7 returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessThan7C function.
func LessThan7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 constraints.Ordered](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return Compare7(host, guest).LT()
}

// LessThan7C returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessThan7 function.
func LessThan7C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6], Ty7 Comparable[Ty7]](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return Compare7C(host, guest).LT()
}

// LessOrEqual7 returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessOrEqual7C function.
func LessOrEqual7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 constraints.Ordered](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return Compare7(host, guest).LE()
}

// LessOrEqual7C returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessOrEqual7 function.
func LessOrEqual7C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6], Ty7 Comparable[Ty7]](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return Compare7C(host, guest).LE()
}

// GreaterThan7 returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterThan7C function.
func GreaterThan7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 constraints.Ordered](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return Compare7(host, guest).GT()
}

// GreaterThan7C returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterThan7 function.
func GreaterThan7C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6], Ty7 Comparable[Ty7]](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return Compare7C(host, guest).GT()
}

// GreaterOrEqual7 returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterOrEqual7C function.
func GreaterOrEqual7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7 constraints.Ordered](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return Compare7(host, guest).GE()
}

// GreaterOrEqual7C returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterOrEqual7 function.
func GreaterOrEqual7C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6], Ty7 Comparable[Ty7]](host, guest T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) bool {
	return Compare7C(host, guest).GE()
}

func (t T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Slice())
}

func (t *T7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7]) UnmarshalJSON(data []byte) error {
	var slice []any
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	unmarshalled, err := FromSlice7[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6, Ty7](slice)
	if err != nil {
		return err
	}

	*t = unmarshalled
	return nil
}
