package tuple_test

import (
	"testing"

	"github.com/barweiss/go-tuple"
	"github.com/stretchr/testify/require"
)

func TestT3_New(t *testing.T) {
	tup := tuple.New3("1", "2", "3")
	require.Equal(t, tuple.T3[string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
	}, tup)
}

func TestT3_Len(t *testing.T) {
	tup := tuple.New3("1", "2", "3")
	require.Equal(t, 3, tup.Len())
}

func TestT3_Values(t *testing.T) {
	tup := tuple.New3("1", "2", "3")
	v1, v2, v3 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
}

func TestT3_String(t *testing.T) {
	tup := tuple.New3("1", "2", "3")
	require.Equal(t, `["1" "2" "3"]`, tup.String())
}

func TestT3_GoString(t *testing.T) {
	tup := tuple.New3("1", "2", "3")
	require.Equal(t, `tuple.T3[string, string, string]{V1: "1", V2: "2", V3: "3"}`, tup.GoString())
}

func TestT3_ToArray(t *testing.T) {
	tup := tuple.New3("1", "2", "3")
	require.Equal(t, [3]any{
		"1", "2", "3",
	}, tup.Array())
}

func TestT3_ToSlice(t *testing.T) {
	tup := tuple.New3("1", "2", "3")
	require.Equal(t, []any{
		"1", "2", "3",
	}, tup.Slice())
}

func TestT3_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [3]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [3]any{
				"1", "2", "3",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [3]any{0, "1", "2"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			array:     [3]any{"0", 1, "2"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			array:     [3]any{"0", "1", 2},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() tuple.T3[string, string, string] {
				return tuple.FromArray3X[string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, tuple.New3("1", "2", "3"), do())
		})
	}
}

func TestT3_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [3]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [3]any{
				"1", "2", "3",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [3]any{1, "2", "3"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			array:   [3]any{"1", 2, "3"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			array:   [3]any{"1", "2", 3},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := tuple.FromArray3[string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tuple.New3("1", "2", "3"), tup)
		})
	}
}

func TestT3_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3",
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
				"3",
			},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0, "1", "2"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			slice:     []any{"0", 1, "2"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			slice:     []any{"0", "1", 2},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() tuple.T3[string, string, string] {
				return tuple.FromSlice3X[string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, tuple.New3("1", "2", "3"), do())
		})
	}
}

func TestT3_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3",
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
				"3",
			},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1, "2", "3"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			slice:   []any{"1", 2, "3"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			slice:   []any{"1", "2", 3},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := tuple.FromSlice3[string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tuple.New3("1", "2", "3"), tup)
		})
	}
}
