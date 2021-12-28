package tuple

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT9_New(t *testing.T) {
	tup := New9("1", "2", "3", "4", "5", "6", "7", "8", "9")
	require.Equal(t, T9[string, string, string, string, string, string, string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
		V4: "4",
		V5: "5",
		V6: "6",
		V7: "7",
		V8: "8",
		V9: "9",
	}, tup)
}

func TestT9_Len(t *testing.T) {
	tup := New9("1", "2", "3", "4", "5", "6", "7", "8", "9")
	require.Equal(t, 9, tup.Len())
}

func TestT9_Values(t *testing.T) {
	tup := New9("1", "2", "3", "4", "5", "6", "7", "8", "9")
	v1, v2, v3, v4, v5, v6, v7, v8, v9 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
	require.Equal(t, "4", v4)
	require.Equal(t, "5", v5)
	require.Equal(t, "6", v6)
	require.Equal(t, "7", v7)
	require.Equal(t, "8", v8)
	require.Equal(t, "9", v9)
}

func TestT9_String(t *testing.T) {
	tup := New9("1", "2", "3", "4", "5", "6", "7", "8", "9")
	require.Equal(t, `["1" "2" "3" "4" "5" "6" "7" "8" "9"]`, tup.String())
}

func TestT9_GoString(t *testing.T) {
	tup := New9("1", "2", "3", "4", "5", "6", "7", "8", "9")
	require.Equal(t, `T9[string, string, string, string, string, string, string, string, string]{V1: "1", V2: "2", V3: "3", V4: "4", V5: "5", V6: "6", V7: "7", V8: "8", V9: "9"}`, tup.GoString())
}

func TestT9_ToArray(t *testing.T) {
	tup := New9("1", "2", "3", "4", "5", "6", "7", "8", "9")
	require.Equal(t, [9]any{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}, tup.Array())
}

func TestT9_ToSlice(t *testing.T) {
	tup := New9("1", "2", "3", "4", "5", "6", "7", "8", "9")
	require.Equal(t, []any{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}, tup.Slice())
}

func TestT9_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [9]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [9]any{
				"1", "2", "3", "4", "5", "6", "7", "8", "9",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [9]any{0, "1", "2", "3", "4", "5", "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			array:     [9]any{"0", 1, "2", "3", "4", "5", "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			array:     [9]any{"0", "1", 2, "3", "4", "5", "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			array:     [9]any{"0", "1", "2", 3, "4", "5", "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			array:     [9]any{"0", "1", "2", "3", 4, "5", "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 6 bad type",
			array:     [9]any{"0", "1", "2", "3", "4", 5, "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 7 bad type",
			array:     [9]any{"0", "1", "2", "3", "4", "5", 6, "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 8 bad type",
			array:     [9]any{"0", "1", "2", "3", "4", "5", "6", 7, "8"},
			wantPanic: true,
		},

		{
			name:      "index 9 bad type",
			array:     [9]any{"0", "1", "2", "3", "4", "5", "6", "7", 8},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T9[string, string, string, string, string, string, string, string, string] {
				return FromArray9X[string, string, string, string, string, string, string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New9("1", "2", "3", "4", "5", "6", "7", "8", "9"), do())
		})
	}
}

func TestT9_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [9]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [9]any{
				"1", "2", "3", "4", "5", "6", "7", "8", "9",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [9]any{1, "2", "3", "4", "5", "6", "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			array:   [9]any{"1", 2, "3", "4", "5", "6", "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			array:   [9]any{"1", "2", 3, "4", "5", "6", "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			array:   [9]any{"1", "2", "3", 4, "5", "6", "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			array:   [9]any{"1", "2", "3", "4", 5, "6", "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 6 bad type",
			array:   [9]any{"1", "2", "3", "4", "5", 6, "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 7 bad type",
			array:   [9]any{"1", "2", "3", "4", "5", "6", 7, "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 8 bad type",
			array:   [9]any{"1", "2", "3", "4", "5", "6", "7", 8, "9"},
			wantErr: true,
		},

		{
			name:    "index 9 bad type",
			array:   [9]any{"1", "2", "3", "4", "5", "6", "7", "8", 9},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromArray9[string, string, string, string, string, string, string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New9("1", "2", "3", "4", "5", "6", "7", "8", "9"), tup)
		})
	}
}

func TestT9_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7", "8", "9",
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
				"3", "4", "5", "6", "7", "8", "9",
			},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7", "8", "9",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0, "1", "2", "3", "4", "5", "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			slice:     []any{"0", 1, "2", "3", "4", "5", "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			slice:     []any{"0", "1", 2, "3", "4", "5", "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			slice:     []any{"0", "1", "2", 3, "4", "5", "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			slice:     []any{"0", "1", "2", "3", 4, "5", "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 6 bad type",
			slice:     []any{"0", "1", "2", "3", "4", 5, "6", "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 7 bad type",
			slice:     []any{"0", "1", "2", "3", "4", "5", 6, "7", "8"},
			wantPanic: true,
		},

		{
			name:      "index 8 bad type",
			slice:     []any{"0", "1", "2", "3", "4", "5", "6", 7, "8"},
			wantPanic: true,
		},

		{
			name:      "index 9 bad type",
			slice:     []any{"0", "1", "2", "3", "4", "5", "6", "7", 8},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T9[string, string, string, string, string, string, string, string, string] {
				return FromSlice9X[string, string, string, string, string, string, string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New9("1", "2", "3", "4", "5", "6", "7", "8", "9"), do())
		})
	}
}

func TestT9_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7", "8", "9",
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
				"3", "4", "5", "6", "7", "8", "9",
			},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7", "8", "9",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1, "2", "3", "4", "5", "6", "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			slice:   []any{"1", 2, "3", "4", "5", "6", "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			slice:   []any{"1", "2", 3, "4", "5", "6", "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			slice:   []any{"1", "2", "3", 4, "5", "6", "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			slice:   []any{"1", "2", "3", "4", 5, "6", "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 6 bad type",
			slice:   []any{"1", "2", "3", "4", "5", 6, "7", "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 7 bad type",
			slice:   []any{"1", "2", "3", "4", "5", "6", 7, "8", "9"},
			wantErr: true,
		},

		{
			name:    "index 8 bad type",
			slice:   []any{"1", "2", "3", "4", "5", "6", "7", 8, "9"},
			wantErr: true,
		},

		{
			name:    "index 9 bad type",
			slice:   []any{"1", "2", "3", "4", "5", "6", "7", "8", 9},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromSlice9[string, string, string, string, string, string, string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New9("1", "2", "3", "4", "5", "6", "7", "8", "9"), tup)
		})
	}
}
