package tuple

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT5_New(t *testing.T) {
	tup := New5("1", "2", "3", "4", "5")
	require.Equal(t, T5[string, string, string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
		V4: "4",
		V5: "5",
	}, tup)
}

func TestT5_Len(t *testing.T) {
	tup := New5("1", "2", "3", "4", "5")
	require.Equal(t, 5, tup.Len())
}

func TestT5_Values(t *testing.T) {
	tup := New5("1", "2", "3", "4", "5")
	v1, v2, v3, v4, v5 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
	require.Equal(t, "4", v4)
	require.Equal(t, "5", v5)
}

func TestT5_String(t *testing.T) {
	tup := New5("1", "2", "3", "4", "5")
	require.Equal(t, `["1" "2" "3" "4" "5"]`, tup.String())
}

func TestT5_GoString(t *testing.T) {
	tup := New5("1", "2", "3", "4", "5")
	require.Equal(t, `tuple.T5[string, string, string, string, string]{V1: "1", V2: "2", V3: "3", V4: "4", V5: "5"}`, tup.GoString())
}

func TestT5_ToArray(t *testing.T) {
	tup := New5("1", "2", "3", "4", "5")
	require.Equal(t, [5]any{
		"1", "2", "3", "4", "5",
	}, tup.Array())
}

func TestT5_ToSlice(t *testing.T) {
	tup := New5("1", "2", "3", "4", "5")
	require.Equal(t, []any{
		"1", "2", "3", "4", "5",
	}, tup.Slice())
}

func TestT5_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [5]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [5]any{
				"1", "2", "3", "4", "5",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [5]any{0, "1", "2", "3", "4"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			array:     [5]any{"0", 1, "2", "3", "4"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			array:     [5]any{"0", "1", 2, "3", "4"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			array:     [5]any{"0", "1", "2", 3, "4"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			array:     [5]any{"0", "1", "2", "3", 4},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T5[string, string, string, string, string] {
				return FromArray5X[string, string, string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New5("1", "2", "3", "4", "5"), do())
		})
	}
}

func TestT5_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [5]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [5]any{
				"1", "2", "3", "4", "5",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [5]any{1, "2", "3", "4", "5"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			array:   [5]any{"1", 2, "3", "4", "5"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			array:   [5]any{"1", "2", 3, "4", "5"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			array:   [5]any{"1", "2", "3", 4, "5"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			array:   [5]any{"1", "2", "3", "4", 5},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromArray5[string, string, string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New5("1", "2", "3", "4", "5"), tup)
		})
	}
}

func TestT5_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5",
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
				"3", "4", "5",
			},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0, "1", "2", "3", "4"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			slice:     []any{"0", 1, "2", "3", "4"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			slice:     []any{"0", "1", 2, "3", "4"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			slice:     []any{"0", "1", "2", 3, "4"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			slice:     []any{"0", "1", "2", "3", 4},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T5[string, string, string, string, string] {
				return FromSlice5X[string, string, string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New5("1", "2", "3", "4", "5"), do())
		})
	}
}

func TestT5_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5",
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
				"3", "4", "5",
			},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1, "2", "3", "4", "5"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			slice:   []any{"1", 2, "3", "4", "5"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			slice:   []any{"1", "2", 3, "4", "5"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			slice:   []any{"1", "2", "3", 4, "5"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			slice:   []any{"1", "2", "3", "4", 5},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromSlice5[string, string, string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New5("1", "2", "3", "4", "5"), tup)
		})
	}
}
