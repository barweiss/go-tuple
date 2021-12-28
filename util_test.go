package tuple

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_typeName_builtin(t *testing.T) {
	require.Equal(t, "string", typeName[string]())
}

func Test_typeName_namedStruct(t *testing.T) {
	type dummy struct{}
	require.Equal(t, "tuple.dummy", typeName[dummy]())
}

func Test_typeName_namedInterface(t *testing.T) {
	type dummy interface{}
	require.Equal(t, "tuple.dummy", typeName[dummy]())
}

func Test_typeName_unnamedInterface(t *testing.T) {
	require.Equal(t, "interface {}", typeName[interface{}]())
}

func Test_typeName_unnamedStruct(t *testing.T) {
	require.Equal(t, "struct {}", typeName[struct{}]())
}

func Test_typeName_func(t *testing.T) {
	type dummy struct{}
	require.Equal(t, "func(interface {}) tuple.dummy", typeName[func(interface{}) dummy]())
}

func Test_typeName_channel(t *testing.T) {
	type dummy struct{}
	require.Equal(t, "chan tuple.dummy", typeName[chan dummy]())
}
