package tuple_test

import (
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/require"
)

func TestT2_New(t *testing.T) {
	tup := tuple.New2("1", "2")
	require.Equal(t, tuple.T2[string, string]{
		V1: "1",
		V2: "2",
	}, tup)
}

func TestT2_Len(t *testing.T) {
	tup := tuple.New2("1", "2")
	require.Equal(t, 2, tup.Len())
}

func TestT2_Values(t *testing.T) {
	tup := tuple.New2("1", "2")
	v1, v2 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
}

func TestT2_String(t *testing.T) {
	tup := tuple.New2("1", "2")
	require.Equal(t, `["1" "2"]`, tup.String())
}

func TestT2_GoString(t *testing.T) {
	tup := tuple.New2("1", "2")
	require.Equal(t, `tuple.T2[string, string]{V1: "1", V2: "2"}`, tup.GoString())
}

func TestT2_ToArray(t *testing.T) {
	tup := tuple.New2("1", "2")
	require.Equal(t, [2]any{
		"1", "2",
	}, tup.Array())
}

func TestT2_ToSlice(t *testing.T) {
	tup := tuple.New2("1", "2")
	require.Equal(t, []any{
		"1", "2",
	}, tup.Slice())
}

func TestT2_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [2]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [2]any{
				"1", "2",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [2]any{0, "1"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			array:     [2]any{"0", 1},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() tuple.T2[string, string] {
				return tuple.FromArray2X[string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, tuple.New2("1", "2"), do())
		})
	}
}

func TestT2_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [2]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [2]any{
				"1", "2",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [2]any{1, "2"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			array:   [2]any{"1", 2},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := tuple.FromArray2[string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tuple.New2("1", "2"), tup)
		})
	}
}

func TestT2_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2",
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
				"1", "2",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0, "1"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			slice:     []any{"0", 1},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() tuple.T2[string, string] {
				return tuple.FromSlice2X[string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, tuple.New2("1", "2"), do())
		})
	}
}

func TestT2_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2",
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
				"1", "2",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1, "2"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			slice:   []any{"1", 2},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := tuple.FromSlice2[string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tuple.New2("1", "2"), tup)
		})
	}
}
