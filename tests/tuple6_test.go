package tuple_test

import (
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/require"
)

func TestT6_New(t *testing.T) {
	tup := tuple.New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, tuple.T6[string, string, string, string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
		V4: "4",
		V5: "5",
		V6: "6",
	}, tup)
}

func TestT6_Len(t *testing.T) {
	tup := tuple.New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, 6, tup.Len())
}

func TestT6_Values(t *testing.T) {
	tup := tuple.New6("1", "2", "3", "4", "5", "6")
	v1, v2, v3, v4, v5, v6 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
	require.Equal(t, "4", v4)
	require.Equal(t, "5", v5)
	require.Equal(t, "6", v6)
}

func TestT6_String(t *testing.T) {
	tup := tuple.New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, `["1" "2" "3" "4" "5" "6"]`, tup.String())
}

func TestT6_GoString(t *testing.T) {
	tup := tuple.New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, `tuple.T6[string, string, string, string, string, string]{V1: "1", V2: "2", V3: "3", V4: "4", V5: "5", V6: "6"}`, tup.GoString())
}

func TestT6_ToArray(t *testing.T) {
	tup := tuple.New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, [6]any{
		"1", "2", "3", "4", "5", "6",
	}, tup.Array())
}

func TestT6_ToSlice(t *testing.T) {
	tup := tuple.New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, []any{
		"1", "2", "3", "4", "5", "6",
	}, tup.Slice())
}

func TestT6_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [6]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [6]any{
				"1", "2", "3", "4", "5", "6",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [6]any{0, "1", "2", "3", "4", "5"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			array:     [6]any{"0", 1, "2", "3", "4", "5"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			array:     [6]any{"0", "1", 2, "3", "4", "5"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			array:     [6]any{"0", "1", "2", 3, "4", "5"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			array:     [6]any{"0", "1", "2", "3", 4, "5"},
			wantPanic: true,
		},

		{
			name:      "index 6 bad type",
			array:     [6]any{"0", "1", "2", "3", "4", 5},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() tuple.T6[string, string, string, string, string, string] {
				return tuple.FromArray6X[string, string, string, string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, tuple.New6("1", "2", "3", "4", "5", "6"), do())
		})
	}
}

func TestT6_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [6]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [6]any{
				"1", "2", "3", "4", "5", "6",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [6]any{1, "2", "3", "4", "5", "6"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			array:   [6]any{"1", 2, "3", "4", "5", "6"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			array:   [6]any{"1", "2", 3, "4", "5", "6"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			array:   [6]any{"1", "2", "3", 4, "5", "6"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			array:   [6]any{"1", "2", "3", "4", 5, "6"},
			wantErr: true,
		},

		{
			name:    "index 6 bad type",
			array:   [6]any{"1", "2", "3", "4", "5", 6},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := tuple.FromArray6[string, string, string, string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tuple.New6("1", "2", "3", "4", "5", "6"), tup)
		})
	}
}

func TestT6_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5", "6",
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
				"3", "4", "5", "6",
			},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5", "6",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0, "1", "2", "3", "4", "5"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			slice:     []any{"0", 1, "2", "3", "4", "5"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			slice:     []any{"0", "1", 2, "3", "4", "5"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			slice:     []any{"0", "1", "2", 3, "4", "5"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			slice:     []any{"0", "1", "2", "3", 4, "5"},
			wantPanic: true,
		},

		{
			name:      "index 6 bad type",
			slice:     []any{"0", "1", "2", "3", "4", 5},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() tuple.T6[string, string, string, string, string, string] {
				return tuple.FromSlice6X[string, string, string, string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, tuple.New6("1", "2", "3", "4", "5", "6"), do())
		})
	}
}

func TestT6_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5", "6",
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
				"3", "4", "5", "6",
			},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5", "6",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1, "2", "3", "4", "5", "6"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			slice:   []any{"1", 2, "3", "4", "5", "6"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			slice:   []any{"1", "2", 3, "4", "5", "6"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			slice:   []any{"1", "2", "3", 4, "5", "6"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			slice:   []any{"1", "2", "3", "4", 5, "6"},
			wantErr: true,
		},

		{
			name:    "index 6 bad type",
			slice:   []any{"1", "2", "3", "4", "5", 6},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := tuple.FromSlice6[string, string, string, string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tuple.New6("1", "2", "3", "4", "5", "6"), tup)
		})
	}
}
