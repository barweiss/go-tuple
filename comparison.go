package tuple

import (
	"golang.org/x/exp/constraints"
)

// OrderedComparisonResult represents the result of a tuple ordered comparison.
// OrderedComparisonResult == 0 represents that the tuples are equal.
// OrderedComparisonResult < 0 represent that the host tuple is less than the guest tuple.
// OrderedComparisonResult > 0 represent that the host tuple is greater than the guest tuple.
type OrderedComparisonResult int

// Comparable is a constraint interface for complex tuple elements that can be compared to other instances.
// In order to compare tuples, either all of their elements must be Ordered, or Comparable.
type Comparable[T any] interface {
	CompareTo(guest T) OrderedComparisonResult
}

// Equalable is a constraint interface for complex tuple elements whose equality to other instances can be tested.
type Equalable[T any] interface {
	Equal(guest T) bool
}

// Equal returns whether the compared values are equal.
func (result OrderedComparisonResult) Equal() bool {
	return result == 0
}

// LessThan returns whether the host is less than the guest.
func (result OrderedComparisonResult) LessThan() bool {
	return result < 0
}

// LessOrEqual returns whether the host is less than or equal to the guest.
func (result OrderedComparisonResult) LessOrEqual() bool {
	return result <= 0
}

// GreaterThan returns whether the host is greater than the guest.
func (result OrderedComparisonResult) GreaterThan() bool {
	return result > 0
}

// GreaterOrEqual returns whether the host is greater than or equal to the guest.
func (result OrderedComparisonResult) GreaterOrEqual() bool {
	return result >= 0
}

// EQ is short for Equal and returns whether the compared values are equal.
func (result OrderedComparisonResult) EQ() bool {
	return result.Equal()
}

// LT is short for LessThan and returns whether the host is less than the guest.
func (result OrderedComparisonResult) LT() bool {
	return result.LessThan()
}

// LE is short for LessOrEqual and returns whether the host is less than or equal to the guest.
func (result OrderedComparisonResult) LE() bool {
	return result.LessOrEqual()
}

// GT is short for GreaterThan and returns whether the host is greater than the guest.
func (result OrderedComparisonResult) GT() bool {
	return result.GreaterThan()
}

// GE is short for GreaterOrEqual and returns whether the host is greater than or equal to the guest.
func (result OrderedComparisonResult) GE() bool {
	return result.GreaterOrEqual()
}

// multiCompare calls and compares the predicates by order.
// multiCompare will short-circuit once one of the predicates returns a non-equal result, and the rest
// of the predicates will not be called.
func multiCompare(predicates ...func() OrderedComparisonResult) OrderedComparisonResult {
	for _, pred := range predicates {
		if result := pred(); !result.Equal() {
			return result
		}
	}

	return 0
}

// compareOrdered returns the comparison result between the host and guest values provided they match the Ordered constraint.
func compareOrdered[T constraints.Ordered](host, guest T) OrderedComparisonResult {
	if host < guest {
		return -1
	}
	if host > guest {
		return 1
	}

	return 0
}
