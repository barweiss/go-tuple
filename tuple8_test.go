package tuple

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT8_New(t *testing.T) {
	tup := New8("1", "2", "3", "4", "5", "6", "7", "8")
	require.Equal(t, T8[string, string, string, string, string, string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
		V4: "4",
		V5: "5",
		V6: "6",
		V7: "7",
		V8: "8",
	}, tup)
}

func TestT8_Len(t *testing.T) {
	tup := New8("1", "2", "3", "4", "5", "6", "7", "8")
	require.Equal(t, 8, tup.Len())
}

func TestT8_Values(t *testing.T) {
	tup := New8("1", "2", "3", "4", "5", "6", "7", "8")
	v1, v2, v3, v4, v5, v6, v7, v8 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
	require.Equal(t, "4", v4)
	require.Equal(t, "5", v5)
	require.Equal(t, "6", v6)
	require.Equal(t, "7", v7)
	require.Equal(t, "8", v8)
}

func TestT8_Compare(t *testing.T) {
	lesser := New8(1, 2, 3, 4, 5, 6, 7, 8)
	greater := New8(2, 3, 4, 5, 6, 7, 8, 9)

	tests := []struct {
		name        string
		host, guest T8[int, int, int, int, int, int, int, int]
		want        OrderedComparisonResult
		wantEQ      bool
		wantLT      bool
		wantLE      bool
		wantGT      bool
		wantGE      bool
	}{
		{
			name:   "less than",
			host:   lesser,
			guest:  greater,
			want:   -1,
			wantLT: true,
			wantLE: true,
		},
		{
			name:   "greater than",
			host:   greater,
			guest:  lesser,
			want:   1,
			wantGT: true,
			wantGE: true,
		},
		{
			name:   "equal",
			host:   lesser,
			guest:  lesser,
			want:   0,
			wantEQ: true,
			wantLE: true,
			wantGE: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compare8(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal8(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan8(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual8(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan8(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual8(tt.host, tt.guest))
		})
	}
}

func TestT8_Compare_Approx(t *testing.T) {
	lesser := New8(approximationHelper("1"), approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"), approximationHelper("6"), approximationHelper("7"), approximationHelper("8"))
	greater := New8(approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"), approximationHelper("6"), approximationHelper("7"), approximationHelper("8"), approximationHelper("9"))

	tests := []struct {
		name        string
		host, guest T8[approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper]
		want        OrderedComparisonResult
		wantEQ      bool
		wantLT      bool
		wantLE      bool
		wantGT      bool
		wantGE      bool
	}{
		{
			name:   "less than",
			host:   lesser,
			guest:  greater,
			want:   -1,
			wantLT: true,
			wantLE: true,
		},
		{
			name:   "greater than",
			host:   greater,
			guest:  lesser,
			want:   1,
			wantGT: true,
			wantGE: true,
		},
		{
			name:   "equal",
			host:   lesser,
			guest:  lesser,
			want:   0,
			wantEQ: true,
			wantLE: true,
			wantGE: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compare8(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal8(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan8(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual8(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan8(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual8(tt.host, tt.guest))
		})
	}
}

func TestT8_CompareC(t *testing.T) {
	lesser := New8(stringComparable("1"), stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"), stringComparable("6"), stringComparable("7"), stringComparable("8"))
	greater := New8(stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"), stringComparable("6"), stringComparable("7"), stringComparable("8"), stringComparable("9"))

	tests := []struct {
		name        string
		host, guest T8[stringComparable, stringComparable, stringComparable, stringComparable, stringComparable, stringComparable, stringComparable, stringComparable]
		want        OrderedComparisonResult
		wantEQ      bool
		wantLT      bool
		wantLE      bool
		wantGT      bool
		wantGE      bool
	}{
		{
			name:   "less than",
			host:   lesser,
			guest:  greater,
			want:   -1,
			wantLT: true,
			wantLE: true,
		},
		{
			name:   "greater than",
			host:   greater,
			guest:  lesser,
			want:   1,
			wantGT: true,
			wantGE: true,
		},
		{
			name:   "equal",
			host:   lesser,
			guest:  lesser,
			want:   0,
			wantEQ: true,
			wantLE: true,
			wantGE: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compare8C(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal8C(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan8C(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual8C(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan8C(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual8C(tt.host, tt.guest))
		})
	}
}

func TestT8_EqualE(t *testing.T) {
	a := New8(intEqualable(1), intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5), intEqualable(6), intEqualable(7), intEqualable(8))
	b := New8(intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5), intEqualable(6), intEqualable(7), intEqualable(8), intEqualable(9))

	require.False(t, Equal8E(a, b))
	require.True(t, Equal8E(a, a))
}

func TestT8_String(t *testing.T) {
	tup := New8("1", "2", "3", "4", "5", "6", "7", "8")
	require.Equal(t, `["1" "2" "3" "4" "5" "6" "7" "8"]`, tup.String())
}

func TestT8_GoString(t *testing.T) {
	tup := New8("1", "2", "3", "4", "5", "6", "7", "8")
	require.Equal(t, `tuple.T8[string, string, string, string, string, string, string, string]{V1: "1", V2: "2", V3: "3", V4: "4", V5: "5", V6: "6", V7: "7", V8: "8"}`, tup.GoString())
}

func TestT8_ToArray(t *testing.T) {
	tup := New8("1", "2", "3", "4", "5", "6", "7", "8")
	require.Equal(t, [8]any{
		"1", "2", "3", "4", "5", "6", "7", "8",
	}, tup.Array())
}

func TestT8_ToSlice(t *testing.T) {
	tup := New8("1", "2", "3", "4", "5", "6", "7", "8")
	require.Equal(t, []any{
		"1", "2", "3", "4", "5", "6", "7", "8",
	}, tup.Slice())
}

func TestT8_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [8]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [8]any{
				"1", "2", "3", "4", "5", "6", "7", "8",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [8]any{0, "1", "2", "3", "4", "5", "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			array:     [8]any{"0", 1, "2", "3", "4", "5", "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			array:     [8]any{"0", "1", 2, "3", "4", "5", "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			array:     [8]any{"0", "1", "2", 3, "4", "5", "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			array:     [8]any{"0", "1", "2", "3", 4, "5", "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 6 bad type",
			array:     [8]any{"0", "1", "2", "3", "4", 5, "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 7 bad type",
			array:     [8]any{"0", "1", "2", "3", "4", "5", 6, "7"},
			wantPanic: true,
		},

		{
			name:      "index 8 bad type",
			array:     [8]any{"0", "1", "2", "3", "4", "5", "6", 7},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T8[string, string, string, string, string, string, string, string] {
				return FromArray8X[string, string, string, string, string, string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New8("1", "2", "3", "4", "5", "6", "7", "8"), do())
		})
	}
}

func TestT8_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [8]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [8]any{
				"1", "2", "3", "4", "5", "6", "7", "8",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [8]any{1, "2", "3", "4", "5", "6", "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			array:   [8]any{"1", 2, "3", "4", "5", "6", "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			array:   [8]any{"1", "2", 3, "4", "5", "6", "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			array:   [8]any{"1", "2", "3", 4, "5", "6", "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			array:   [8]any{"1", "2", "3", "4", 5, "6", "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 6 bad type",
			array:   [8]any{"1", "2", "3", "4", "5", 6, "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 7 bad type",
			array:   [8]any{"1", "2", "3", "4", "5", "6", 7, "8"},
			wantErr: true,
		},

		{
			name:    "index 8 bad type",
			array:   [8]any{"1", "2", "3", "4", "5", "6", "7", 8},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromArray8[string, string, string, string, string, string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New8("1", "2", "3", "4", "5", "6", "7", "8"), tup)
		})
	}
}

func TestT8_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7", "8",
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
				"3", "4", "5", "6", "7", "8",
			},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7", "8",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0, "1", "2", "3", "4", "5", "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			slice:     []any{"0", 1, "2", "3", "4", "5", "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			slice:     []any{"0", "1", 2, "3", "4", "5", "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			slice:     []any{"0", "1", "2", 3, "4", "5", "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			slice:     []any{"0", "1", "2", "3", 4, "5", "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 6 bad type",
			slice:     []any{"0", "1", "2", "3", "4", 5, "6", "7"},
			wantPanic: true,
		},

		{
			name:      "index 7 bad type",
			slice:     []any{"0", "1", "2", "3", "4", "5", 6, "7"},
			wantPanic: true,
		},

		{
			name:      "index 8 bad type",
			slice:     []any{"0", "1", "2", "3", "4", "5", "6", 7},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T8[string, string, string, string, string, string, string, string] {
				return FromSlice8X[string, string, string, string, string, string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New8("1", "2", "3", "4", "5", "6", "7", "8"), do())
		})
	}
}

func TestT8_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7", "8",
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
				"3", "4", "5", "6", "7", "8",
			},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7", "8",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1, "2", "3", "4", "5", "6", "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			slice:   []any{"1", 2, "3", "4", "5", "6", "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			slice:   []any{"1", "2", 3, "4", "5", "6", "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			slice:   []any{"1", "2", "3", 4, "5", "6", "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			slice:   []any{"1", "2", "3", "4", 5, "6", "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 6 bad type",
			slice:   []any{"1", "2", "3", "4", "5", 6, "7", "8"},
			wantErr: true,
		},

		{
			name:    "index 7 bad type",
			slice:   []any{"1", "2", "3", "4", "5", "6", 7, "8"},
			wantErr: true,
		},

		{
			name:    "index 8 bad type",
			slice:   []any{"1", "2", "3", "4", "5", "6", "7", 8},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromSlice8[string, string, string, string, string, string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New8("1", "2", "3", "4", "5", "6", "7", "8"), tup)
		})
	}
}
