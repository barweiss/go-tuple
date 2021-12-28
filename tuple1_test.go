package tuple

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT1_New(t *testing.T) {
	tup := New1("1")
	require.Equal(t, T1[string]{
		V1: "1",
	}, tup)
}

func TestT1_Len(t *testing.T) {
	tup := New1("1")
	require.Equal(t, 1, tup.Len())
}

func TestT1_Values(t *testing.T) {
	tup := New1("1")
	v1 := tup.Values()
	require.Equal(t, "1", v1)
}

func TestT1_String(t *testing.T) {
	tup := New1("1")
	require.Equal(t, `["1"]`, tup.String())
}

func TestT1_GoString(t *testing.T) {
	tup := New1("1")
	require.Equal(t, `T1[string]{V1: "1"}`, tup.GoString())
}

func TestT1_ToArray(t *testing.T) {
	tup := New1("1")
	require.Equal(t, [1]any{
		"1",
	}, tup.Array())
}

func TestT1_ToSlice(t *testing.T) {
	tup := New1("1")
	require.Equal(t, []any{
		"1",
	}, tup.Slice())
}

func TestT1_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [1]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [1]any{
				"1",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [1]any{0},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T1[string] {
				return FromArray1X[string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New1("1"), do())
		})
	}
}

func TestT1_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [1]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [1]any{
				"1",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [1]any{1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromArray1[string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New1("1"), tup)
		})
	}
}

func TestT1_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1",
			},
			wantPanic: false,
		},
		{
			name:      "slice empty",
			slice:     []any{},
			wantPanic: true,
		},
		{
			name:      "slice too short",
			slice:     []any{},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T1[string] {
				return FromSlice1X[string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New1("1"), do())
		})
	}
}

func TestT1_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1",
			},
			wantErr: false,
		},
		{
			name:    "slice empty",
			slice:   []any{},
			wantErr: true,
		},
		{
			name:    "slice too short",
			slice:   []any{},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromSlice1[string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New1("1"), tup)
		})
	}
}
