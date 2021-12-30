package tuple_test

import (
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/require"
)

func TestT7_New(t *testing.T) {
	tup := tuple.New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, tuple.T7[string, string, string, string, string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
		V4: "4",
		V5: "5",
		V6: "6",
		V7: "7",
	}, tup)
}

func TestT7_Len(t *testing.T) {
	tup := tuple.New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, 7, tup.Len())
}

func TestT7_Values(t *testing.T) {
	tup := tuple.New7("1", "2", "3", "4", "5", "6", "7")
	v1, v2, v3, v4, v5, v6, v7 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
	require.Equal(t, "4", v4)
	require.Equal(t, "5", v5)
	require.Equal(t, "6", v6)
	require.Equal(t, "7", v7)
}

func TestT7_String(t *testing.T) {
	tup := tuple.New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, `["1" "2" "3" "4" "5" "6" "7"]`, tup.String())
}

func TestT7_GoString(t *testing.T) {
	tup := tuple.New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, `tuple.T7[string, string, string, string, string, string, string]{V1: "1", V2: "2", V3: "3", V4: "4", V5: "5", V6: "6", V7: "7"}`, tup.GoString())
}

func TestT7_ToArray(t *testing.T) {
	tup := tuple.New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, [7]any{
		"1", "2", "3", "4", "5", "6", "7",
	}, tup.Array())
}

func TestT7_ToSlice(t *testing.T) {
	tup := tuple.New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, []any{
		"1", "2", "3", "4", "5", "6", "7",
	}, tup.Slice())
}

func TestT7_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [7]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [7]any{
				"1", "2", "3", "4", "5", "6", "7",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [7]any{0, "1", "2", "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			array:     [7]any{"0", 1, "2", "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			array:     [7]any{"0", "1", 2, "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			array:     [7]any{"0", "1", "2", 3, "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			array:     [7]any{"0", "1", "2", "3", 4, "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 6 bad type",
			array:     [7]any{"0", "1", "2", "3", "4", 5, "6"},
			wantPanic: true,
		},

		{
			name:      "index 7 bad type",
			array:     [7]any{"0", "1", "2", "3", "4", "5", 6},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() tuple.T7[string, string, string, string, string, string, string] {
				return tuple.FromArray7X[string, string, string, string, string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, tuple.New7("1", "2", "3", "4", "5", "6", "7"), do())
		})
	}
}

func TestT7_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [7]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [7]any{
				"1", "2", "3", "4", "5", "6", "7",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [7]any{1, "2", "3", "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			array:   [7]any{"1", 2, "3", "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			array:   [7]any{"1", "2", 3, "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			array:   [7]any{"1", "2", "3", 4, "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			array:   [7]any{"1", "2", "3", "4", 5, "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 6 bad type",
			array:   [7]any{"1", "2", "3", "4", "5", 6, "7"},
			wantErr: true,
		},

		{
			name:    "index 7 bad type",
			array:   [7]any{"1", "2", "3", "4", "5", "6", 7},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := tuple.FromArray7[string, string, string, string, string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tuple.New7("1", "2", "3", "4", "5", "6", "7"), tup)
		})
	}
}

func TestT7_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7",
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
				"3", "4", "5", "6", "7",
			},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0, "1", "2", "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			slice:     []any{"0", 1, "2", "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			slice:     []any{"0", "1", 2, "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			slice:     []any{"0", "1", "2", 3, "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			slice:     []any{"0", "1", "2", "3", 4, "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 6 bad type",
			slice:     []any{"0", "1", "2", "3", "4", 5, "6"},
			wantPanic: true,
		},

		{
			name:      "index 7 bad type",
			slice:     []any{"0", "1", "2", "3", "4", "5", 6},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() tuple.T7[string, string, string, string, string, string, string] {
				return tuple.FromSlice7X[string, string, string, string, string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, tuple.New7("1", "2", "3", "4", "5", "6", "7"), do())
		})
	}
}

func TestT7_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7",
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
				"3", "4", "5", "6", "7",
			},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1, "2", "3", "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			slice:   []any{"1", 2, "3", "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			slice:   []any{"1", "2", 3, "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			slice:   []any{"1", "2", "3", 4, "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			slice:   []any{"1", "2", "3", "4", 5, "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 6 bad type",
			slice:   []any{"1", "2", "3", "4", "5", 6, "7"},
			wantErr: true,
		},

		{
			name:    "index 7 bad type",
			slice:   []any{"1", "2", "3", "4", "5", "6", 7},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := tuple.FromSlice7[string, string, string, string, string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tuple.New7("1", "2", "3", "4", "5", "6", "7"), tup)
		})
	}
}
