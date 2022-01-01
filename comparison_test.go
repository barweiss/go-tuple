package tuple

// approximationHelper is a helper type for testing type approximation.
type approximationHelper string

// intEqualable is a wrapper type for int that implements the Equalable constraint.
type intEqualable int

// stringComparable is a wrapper type for string that implements the Comparable constraint.
type stringComparable string

// Assert implementation.
var _ Equalable[intEqualable] = (intEqualable)(0)
var _ Comparable[stringComparable] = (stringComparable)("")

func (i intEqualable) Equal(other intEqualable) bool {
	return i == other
}

func (s stringComparable) CompareTo(other stringComparable) OrderedComparisonResult {
	return compareOrdered(s, other)
}
