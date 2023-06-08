// package tuple defines tuple types that can hold multiple values of varying types. Currently, tuples with up to 9 values are supported.
//
// Tuple methods:
//
// * Len      returns the number of values held by the tuple.
// * Values   returns the values held by the tuple.
// * Array    returns an array of the tuple values.
// * Slice    returns a slice of the tuple values.
// * String   returns the string representation of the tuple.
// * GoString returns a Go-syntax representation of the tuple.
//
// Tuple creation functions:
//
// * New<N>        creates a new tuple holding N generic values.
// * FromArray<N>  returns a tuple from an array of length N.
//    If any of the values can not be converted to the generic type, an error is returned.
// * FromArray<N>X returns a tuple from an array of length N.
//    If any of the values can not be converted to the generic type, the function panics.
// * FromSlice<N>  returns a tuple from a slice of length N.
//    If the length of the slice doesn't match, or any of the values can not be converted to the generic type, an error is returned.
// * FromSlice<N>X returns a tuple from a slice of length N.
//    If the length of the slice doesn't match, or any of the values can not be converted to the generic type, the function panics.
//
// Tuple comparison functions:
//
// * Equal<N> returns whether the host tuple is equal to the other tuple.
// * Compare<N> returns whether the host tuple is semantically less than, equal to, or greater than the guest tuple.
// * LessThan<N> returns whether the host tuple is semantically less than the guest tuple.
// * LessOrEqual<N> returns whether the host tuple is semantically less than or equal to the guest tuple.
// * GreaterThan<N> returns whether the host tuple is semantically greater than the guest tuple.
// * GreaterOrEqual<N> returns whether the host tuple is semantically greater than or equal to the guest tuple.
//
// Tuple comparison functions may have an "C" or "E" suffix as overload with additional supported type constraints.
// Comparison functions ending with "C" accept the "Comparable" constraint.
// Comparison functions ending with "E" accept the "Equalable contraint.
package tuple
