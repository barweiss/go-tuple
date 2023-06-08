package tuple

import (
	"constraints"
	"encoding/json"
	"fmt"
)

// T1 is a tuple type holding 1 generic values.
type T1[Ty1 any] struct {
	V1 Ty1
}

// Len returns the number of values held by the tuple.
func (t T1[Ty1]) Len() int {
	return 1
}

// Values returns the values held by the tuple.
func (t T1[Ty1]) Values() Ty1 {
	return t.V1
}

// Array returns an array of the tuple values.
func (t T1[Ty1]) Array() [1]any {
	return [1]any{
		t.V1,
	}
}

// Slice returns a slice of the tuple values.
func (t T1[Ty1]) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t T1[Ty1]) String() string {
	return tupString(t.Slice())
}

// GoString returns a Go-syntax representation of the tuple.
func (t T1[Ty1]) GoString() string {
	return tupGoString(t.Slice())
}

// New1 creates a new tuple holding 1 generic values.
func New1[Ty1 any](v1 Ty1) T1[Ty1] {
	return T1[Ty1]{
		V1: v1,
	}
}

// FromArray1 returns a tuple from an array of length 1.
// If any of the values can not be converted to the generic type, an error is returned.
func FromArray1[Ty1 any](arr [1]any) (T1[Ty1], error) {
	v1, ok := arr[0].(Ty1)
	if !ok {
		return T1[Ty1]{}, fmt.Errorf("value at array index 0 expected to have type %s but has type %T", typeName[Ty1](), arr[0])
	}

	return New1(v1), nil
}

// FromArray1X returns a tuple from an array of length 1.
// If any of the values can not be converted to the generic type, the function panics.
func FromArray1X[Ty1 any](arr [1]any) T1[Ty1] {
	return FromSlice1X[Ty1](arr[:])
}

// FromSlice1 returns a tuple from a slice of length 1.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
func FromSlice1[Ty1 any](values []any) (T1[Ty1], error) {
	if len(values) != 1 {
		return T1[Ty1]{}, fmt.Errorf("slice length %d must match number of tuple values 1", len(values))
	}

	v1, ok := values[0].(Ty1)
	if !ok {
		return T1[Ty1]{}, fmt.Errorf("value at slice index 0 expected to have type %s but has type %T", typeName[Ty1](), values[0])
	}

	return New1(v1), nil
}

// FromSlice1X returns a tuple from a slice of length 1.
// If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
func FromSlice1X[Ty1 any](values []any) T1[Ty1] {
	if len(values) != 1 {
		panic(fmt.Errorf("slice length %d must match number of tuple values 1", len(values)))
	}

	v1 := values[0].(Ty1)

	return New1(v1)
}

// Equal1 returns whether the host tuple is equal to the other tuple.
// All tuple elements of the host and guest parameters must match the "comparable" built-in constraint.
// To test equality of tuples that hold custom Equalable values, use the Equal1E function.
// To test equality of tuples that hold custom Comparable values, use the Equal1C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal1[Ty1 comparable](host, guest T1[Ty1]) bool {
	return host.V1 == guest.V1
}

// Equal1E returns whether the host tuple is semantically equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Equalable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal1 function.
// To test equality of tuples that hold custom Comparable values, use the Equal1C function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal1E[Ty1 Equalable[Ty1]](host, guest T1[Ty1]) bool {
	return host.V1.Equal(guest.V1)
}

// Equal1C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To test equality of tuples that hold built-in "comparable" values, use the Equal1 function.
// To test equality of tuples that hold custom Equalable values, use the Equal1E function.
// Otherwise, use Equal or reflect.DeepEqual to test tuples of any types.
func Equal1C[Ty1 Comparable[Ty1]](host, guest T1[Ty1]) bool {
	return host.V1.CompareTo(guest.V1).EQ()
}

// Compare1 returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the Compare1C function.
func Compare1[Ty1 constraints.Ordered](host, guest T1[Ty1]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return compareOrdered(host.V1, guest.V1) },
	)
}

// Compare1C returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the Compare1 function.
func Compare1C[Ty1 Comparable[Ty1]](host, guest T1[Ty1]) OrderedComparisonResult {
	return multiCompare(
		func() OrderedComparisonResult { return host.V1.CompareTo(guest.V1) },
	)
}

// LessThan1 returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessThan1C function.
func LessThan1[Ty1 constraints.Ordered](host, guest T1[Ty1]) bool {
	return Compare1(host, guest).LT()
}

// LessThan1C returns whether the host tuple is semantically less than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessThan1 function.
func LessThan1C[Ty1 Comparable[Ty1]](host, guest T1[Ty1]) bool {
	return Compare1C(host, guest).LT()
}

// LessOrEqual1 returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the LessOrEqual1C function.
func LessOrEqual1[Ty1 constraints.Ordered](host, guest T1[Ty1]) bool {
	return Compare1(host, guest).LE()
}

// LessOrEqual1C returns whether the host tuple is semantically less than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the LessOrEqual1 function.
func LessOrEqual1C[Ty1 Comparable[Ty1]](host, guest T1[Ty1]) bool {
	return Compare1C(host, guest).LE()
}

// GreaterThan1 returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterThan1C function.
func GreaterThan1[Ty1 constraints.Ordered](host, guest T1[Ty1]) bool {
	return Compare1(host, guest).GT()
}

// GreaterThan1C returns whether the host tuple is semantically greater than the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterThan1 function.
func GreaterThan1C[Ty1 Comparable[Ty1]](host, guest T1[Ty1]) bool {
	return Compare1C(host, guest).GT()
}

// GreaterOrEqual1 returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the "Ordered" constraint.
// To compare tuples that hold custom comparable values, use the GreaterOrEqual1C function.
func GreaterOrEqual1[Ty1 constraints.Ordered](host, guest T1[Ty1]) bool {
	return Compare1(host, guest).GE()
}

// GreaterOrEqual1C returns whether the host tuple is semantically greater than or equal to the guest tuple.
// All tuple elements of the host and guest parameters must match the Comparable constraint.
// To compare tuples that hold built-in "Ordered" values, use the GreaterOrEqual1 function.
func GreaterOrEqual1C[Ty1 Comparable[Ty1]](host, guest T1[Ty1]) bool {
	return Compare1C(host, guest).GE()
}

func (t T1[Ty1]) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Slice())
}

func (t *T1[Ty1]) UnmarshalJSON(data []byte) error {
	var slice []any
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	unmarshalled, err := FromSlice1[Ty1](slice)
	if err != nil {
		return err
	}

	*t = unmarshalled
	return nil
}
