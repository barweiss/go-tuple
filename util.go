package tuple

import (
	"fmt"
	"reflect"
	"strings"
)

// typeName returns the name of the type parameters.
func typeName[T any]() string {
	var val T
	ty := reflect.TypeOf(&val).Elem()
	return ty.String()
}

// tupString returns a string representation of the tuple values.
func tupString(values []any) string {
	valuesStr := make([]string, len(values))
	for i, val := range values {
		valuesStr[i] = fmt.Sprintf("%#v", val)
	}

	return "[" + strings.Join(valuesStr, " ") + "]"
}

// tupGoString returns a Go-syntax representation of a tuple holding the given values.
func tupGoString(values []any) string {
	types := make([]string, len(values))
	for i, val := range values {
		types[i] = fmt.Sprintf("%T", val)
	}

	fields := make([]string, len(values))
	for i, val := range values {
		fields[i] = fmt.Sprintf("V%d: %#v", i+1, val)
	}

	return fmt.Sprintf("tuple.T%d[%s]{%s}",
		len(values),
		strings.Join(types, ", "),
		strings.Join(fields, ", "),
	)
}
