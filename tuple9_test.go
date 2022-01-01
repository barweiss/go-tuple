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

func TestT9_Compare(t *testing.T) {
	lesser := New9(1, 2, 3, 4, 5, 6, 7, 8, 9)
	greater := New9(2, 3, 4, 5, 6, 7, 8, 9, 10)

	tests := []struct {
		name        string
		host, guest T9[int, int, int, int, int, int, int, int, int]
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
			got := Compare9(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal9(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan9(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual9(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan9(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual9(tt.host, tt.guest))
		})
	}
}

func TestT9_Compare_Approx(t *testing.T) {
	lesser := New9(approximationHelper("1"), approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"), approximationHelper("6"), approximationHelper("7"), approximationHelper("8"), approximationHelper("9"))
	greater := New9(approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"), approximationHelper("6"), approximationHelper("7"), approximationHelper("8"), approximationHelper("9"), approximationHelper("10"))

	tests := []struct {
		name        string
		host, guest T9[approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper]
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
			got := Compare9(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal9(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan9(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual9(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan9(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual9(tt.host, tt.guest))
		})
	}
}

func TestT9_CompareC(t *testing.T) {
	lesser := New9(stringComparable("1"), stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"), stringComparable("6"), stringComparable("7"), stringComparable("8"), stringComparable("9"))
	greater := New9(stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"), stringComparable("6"), stringComparable("7"), stringComparable("8"), stringComparable("9"), stringComparable("10"))

	tests := []struct {
		name        string
		host, guest T9[stringComparable, stringComparable, stringComparable, stringComparable, stringComparable, stringComparable, stringComparable, stringComparable, stringComparable]
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
			got := Compare9C(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal9C(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan9C(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual9C(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan9C(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual9C(tt.host, tt.guest))
		})
	}
}

func TestT9_EqualE(t *testing.T) {
	a := New9(intEqualable(1), intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5), intEqualable(6), intEqualable(7), intEqualable(8), intEqualable(9))
	b := New9(intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5), intEqualable(6), intEqualable(7), intEqualable(8), intEqualable(9), intEqualable(10))

	require.False(t, Equal9E(a, b))
	require.True(t, Equal9E(a, a))
}

func TestT9_String(t *testing.T) {
	tup := New9("1", "2", "3", "4", "5", "6", "7", "8", "9")
	require.Equal(t, `["1" "2" "3" "4" "5" "6" "7" "8" "9"]`, tup.String())
}

func TestT9_GoString(t *testing.T) {
	tup := New9("1", "2", "3", "4", "5", "6", "7", "8", "9")
	require.Equal(t, `tuple.T9[string, string, string, string, string, string, string, string, string]{V1: "1", V2: "2", V3: "3", V4: "4", V5: "5", V6: "6", V7: "7", V8: "8", V9: "9"}`, tup.GoString())
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
