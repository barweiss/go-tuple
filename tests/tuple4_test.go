package tuple_test

import (
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/require"
)

func TestT4_New(t *testing.T) {
	tup := tuple.New4("1", "2", "3", "4")
	require.Equal(t, tuple.T4[string, string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
		V4: "4",
	}, tup)
}

func TestT4_Len(t *testing.T) {
	tup := tuple.New4("1", "2", "3", "4")
	require.Equal(t, 4, tup.Len())
}

func TestT4_Values(t *testing.T) {
	tup := tuple.New4("1", "2", "3", "4")
	v1, v2, v3, v4 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
	require.Equal(t, "4", v4)
}

func TestT4_String(t *testing.T) {
	tup := tuple.New4("1", "2", "3", "4")
	require.Equal(t, `["1" "2" "3" "4"]`, tup.String())
}

func TestT4_GoString(t *testing.T) {
	tup := tuple.New4("1", "2", "3", "4")
	require.Equal(t, `tuple.T4[string, string, string, string]{V1: "1", V2: "2", V3: "3", V4: "4"}`, tup.GoString())
}

func TestT4_ToArray(t *testing.T) {
	tup := tuple.New4("1", "2", "3", "4")
	require.Equal(t, [4]any{
		"1", "2", "3", "4",
	}, tup.Array())
}

func TestT4_ToSlice(t *testing.T) {
	tup := tuple.New4("1", "2", "3", "4")
	require.Equal(t, []any{
		"1", "2", "3", "4",
	}, tup.Slice())
}

func TestT4_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [4]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [4]any{
				"1", "2", "3", "4",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [4]any{0, "1", "2", "3"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			array:     [4]any{"0", 1, "2", "3"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			array:     [4]any{"0", "1", 2, "3"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			array:     [4]any{"0", "1", "2", 3},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() tuple.T4[string, string, string, string] {
				return tuple.FromArray4X[string, string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, tuple.New4("1", "2", "3", "4"), do())
		})
	}
}

func TestT4_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [4]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [4]any{
				"1", "2", "3", "4",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [4]any{1, "2", "3", "4"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			array:   [4]any{"1", 2, "3", "4"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			array:   [4]any{"1", "2", 3, "4"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			array:   [4]any{"1", "2", "3", 4},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := tuple.FromArray4[string, string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tuple.New4("1", "2", "3", "4"), tup)
		})
	}
}

func TestT4_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4",
			},
			wantPanic: false,
		},
		{
			name:      "slice empty",
			slice:     []any{},
			wantPanic: true,
		},
		{
			name: "slice too short",
			slice: []any{
				"3", "4",
			},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0, "1", "2", "3"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			slice:     []any{"0", 1, "2", "3"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			slice:     []any{"0", "1", 2, "3"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			slice:     []any{"0", "1", "2", 3},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() tuple.T4[string, string, string, string] {
				return tuple.FromSlice4X[string, string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, tuple.New4("1", "2", "3", "4"), do())
		})
	}
}

func TestT4_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4",
			},
			wantErr: false,
		},
		{
			name:    "slice empty",
			slice:   []any{},
			wantErr: true,
		},
		{
			name: "slice too short",
			slice: []any{
				"3", "4",
			},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1, "2", "3", "4"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			slice:   []any{"1", 2, "3", "4"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			slice:   []any{"1", "2", 3, "4"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			slice:   []any{"1", "2", "3", 4},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := tuple.FromSlice4[string, string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tuple.New4("1", "2", "3", "4"), tup)
		})
	}
}
