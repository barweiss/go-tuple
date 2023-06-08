package tuple

import (
	"encoding/json"
	"fmt"

	"golang.org/x/exp/constraints"
)

// T6 is a tuple type holding 6 generic values.
type T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 any] struct {
	V1 Ty1
	V2 Ty2
	V3 Ty3
	V4 Ty4
	V5 Ty5
	V6 Ty6
}

// Len returns the number of values held by the tuple.
func (t T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) Len() int {
	return 6
}

// Values returns the values held by the tuple.
func (t T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) Values() (Ty1, Ty2, Ty3, Ty4, Ty5, Ty6) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6
}

// Array returns an array of the tuple values.
func (t T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) Array() [6]any {
	return [6]any{
		t.V1,
		t.V2,
		t.V3,
		t.V4,
		t.V5,
		t.V6,
	}
}

// Slice returns a slice of the tuple values.
func (t T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) GoString() string {
	return tupGoString(t.Slice())
}

// New6 creates a new tuple holding 6 generic values.
func New6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 any](v1 Ty1, v2 Ty2, v3 Ty3, v4 Ty4, v5 Ty5, v6 Ty6) T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6] {
	return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
	}
}

// FromArray6 returns a tuple from an array of length 6.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 any](arr [6]any) (T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6], error) {
	v1, ok := arr[0].(Ty1)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at array index 0 expected to have type %s but has type %T", typeName[Ty1](), arr[0])
	}
	v2, ok := arr[1].(Ty2)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at array index 1 expected to have type %s but has type %T", typeName[Ty2](), arr[1])
	}
	v3, ok := arr[2].(Ty3)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at array index 2 expected to have type %s but has type %T", typeName[Ty3](), arr[2])
	}
	v4, ok := arr[3].(Ty4)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at array index 3 expected to have type %s but has type %T", typeName[Ty4](), arr[3])
	}
	v5, ok := arr[4].(Ty5)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at array index 4 expected to have type %s but has type %T", typeName[Ty5](), arr[4])
	}
	v6, ok := arr[5].(Ty6)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at array index 5 expected to have type %s but has type %T", typeName[Ty6](), arr[5])
	}

	return New6(v1, v2, v3, v4, v5, v6), nil
}

// FromArray6X returns a tuple from an array of length 6.
// If any of the values can not be converted to the generic type, the function panics.
func FromArray6X[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 any](arr [6]any) T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6] {
	return FromSlice6X[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6](arr[:])
}

// FromSlice6 returns a tuple from a slice of length 6.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 any](values []any) (T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6], error) {
	if len(values) != 6 {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("slice length %d must match number of tuple values 6", len(values))
	}

	v1, ok := values[0].(Ty1)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at slice index 0 expected to have type %s but has type %T", typeName[Ty1](), values[0])
	}
	v2, ok := values[1].(Ty2)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at slice index 1 expected to have type %s but has type %T", typeName[Ty2](), values[1])
	}
	v3, ok := values[2].(Ty3)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at slice index 2 expected to have type %s but has type %T", typeName[Ty3](), values[2])
	}
	v4, ok := values[3].(Ty4)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at slice index 3 expected to have type %s but has type %T", typeName[Ty4](), values[3])
	}
	v5, ok := values[4].(Ty5)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at slice index 4 expected to have type %s but has type %T", typeName[Ty5](), values[4])
	}
	v6, ok := values[5].(Ty6)
	if !ok {
		return T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]{}, fmt.Errorf("value at slice index 5 expected to have type %s but has type %T", typeName[Ty6](), values[5])
	}

	return New6(v1, v2, v3, v4, v5, v6), nil
}

// FromSlice6X returns a tuple from a slice of length 6.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
func FromSlice6X[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 any](values []any) T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6] {
	if len(values) != 6 {
		panic(fmt.Errorf("slice length %d must match number of tuple values 6", len(values)))
	}

	v1 := values[0].(Ty1)
	v2 := values[1].(Ty2)
	v3 := values[2].(Ty3)
	v4 := values[3].(Ty4)
	v5 := values[4].(Ty5)
	v6 := values[5].(Ty6)

	return New6(v1, v2, v3, v4, v5, v6)
}

// Equal6 returns whether the host tuple is equal to the other tuple.
// All tuple elements of the host and guest parameters must match the "comparable" built-in constraint.
// To test equality of tuples that hold custom Equalable values, use the Equal6E function.
// To test equality of tuples that hold custom Comparable values, use the Equal6C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 comparable](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return host.V1 == guest.V1 && host.V2 == guest.V2 && host.V3 == guest.V3 && host.V4 == guest.V4 && host.V5 == guest.V5 && host.V6 == guest.V6
}

// Equal6E returns whether the host tuple is semantically equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Equalable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal6 function.
// To test equality of tuples that hold custom Comparable values, use the Equal6C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal6E[Ty1 Equalable[Ty1], Ty2 Equalable[Ty2], Ty3 Equalable[Ty3], Ty4 Equalable[Ty4], Ty5 Equalable[Ty5], Ty6 Equalable[Ty6]](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return host.V1.Equal(guest.V1) && host.V2.Equal(guest.V2) && host.V3.Equal(guest.V3) && host.V4.Equal(guest.V4) && host.V5.Equal(guest.V5) && host.V6.Equal(guest.V6)
}

// Equal6C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal6 function.
// To test equality of tuples that hold custom Equalable values, use the Equal6E function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal6C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6]](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return host.V1.CompareTo(guest.V1).EQ() && host.V2.CompareTo(guest.V2).EQ() && host.V3.CompareTo(guest.V3).EQ() && host.V4.CompareTo(guest.V4).EQ() && host.V5.CompareTo(guest.V5).EQ() && host.V6.CompareTo(guest.V6).EQ()
}

// Compare6 returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the Compare6C function.
func Compare6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 constraints.Ordered](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return compareOrdered(host.V1, guest.V1) },

		func() OrderedComparisonResult { return compareOrdered(host.V2, guest.V2) },

		func() OrderedComparisonResult { return compareOrdered(host.V3, guest.V3) },

		func() OrderedComparisonResult { return compareOrdered(host.V4, guest.V4) },

		func() OrderedComparisonResult { return compareOrdered(host.V5, guest.V5) },

		func() OrderedComparisonResult { return compareOrdered(host.V6, guest.V6) },
	)
}

// Compare6C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the Compare6 function.
func Compare6C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6]](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return host.V1.CompareTo(guest.V1) },

		func() OrderedComparisonResult { return host.V2.CompareTo(guest.V2) },

		func() OrderedComparisonResult { return host.V3.CompareTo(guest.V3) },

		func() OrderedComparisonResult { return host.V4.CompareTo(guest.V4) },

		func() OrderedComparisonResult { return host.V5.CompareTo(guest.V5) },

		func() OrderedComparisonResult { return host.V6.CompareTo(guest.V6) },
	)
}

// LessThan6 returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessThan6C function.
func LessThan6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 constraints.Ordered](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return Compare6(host, guest).LT()
}

// LessThan6C returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessThan6 function.
func LessThan6C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6]](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return Compare6C(host, guest).LT()
}

// LessOrEqual6 returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessOrEqual6C function.
func LessOrEqual6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 constraints.Ordered](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return Compare6(host, guest).LE()
}

// LessOrEqual6C returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessOrEqual6 function.
func LessOrEqual6C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6]](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return Compare6C(host, guest).LE()
}

// GreaterThan6 returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterThan6C function.
func GreaterThan6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 constraints.Ordered](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return Compare6(host, guest).GT()
}

// GreaterThan6C returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterThan6 function.
func GreaterThan6C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6]](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return Compare6C(host, guest).GT()
}

// GreaterOrEqual6 returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterOrEqual6C function.
func GreaterOrEqual6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6 constraints.Ordered](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return Compare6(host, guest).GE()
}

// GreaterOrEqual6C returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterOrEqual6 function.
func GreaterOrEqual6C[Ty1 Comparable[Ty1], Ty2 Comparable[Ty2], Ty3 Comparable[Ty3], Ty4 Comparable[Ty4], Ty5 Comparable[Ty5], Ty6 Comparable[Ty6]](host, guest T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) bool {
	return Compare6C(host, guest).GE()
}

func (t T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Slice())
}

func (t *T6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6]) UnmarshalJSON(data []byte) error {
	var slice []any
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	unmarshalled, err := FromSlice6[Ty1, Ty2, Ty3, Ty4, Ty5, Ty6](slice)
	if err != nil {
		return err
	}

	*t = unmarshalled
	return nil
}
