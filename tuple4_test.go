package tuple

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT4_New(t *testing.T) {
	tup := New4("1", "2", "3", "4")
	require.Equal(t, T4[string, string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
		V4: "4",
	}, tup)
}

func TestT4_Len(t *testing.T) {
	tup := New4("1", "2", "3", "4")
	require.Equal(t, 4, tup.Len())
}

func TestT4_Values(t *testing.T) {
	tup := New4("1", "2", "3", "4")
	v1, v2, v3, v4 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
	require.Equal(t, "4", v4)
}

func TestT4_Compare(t *testing.T) {
	lesser := New4(1, 2, 3, 4)
	greater := New4(2, 3, 4, 5)

	tests := []struct {
		name        string
		host, guest T4[int, int, int, int]
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
			got := Compare4(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal4(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan4(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual4(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan4(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual4(tt.host, tt.guest))
		})
	}
}

func TestT4_Compare_Approx(t *testing.T) {
	lesser := New4(approximationHelper("1"), approximationHelper("2"), approximationHelper("3"), approximationHelper("4"))
	greater := New4(approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"))

	tests := []struct {
		name        string
		host, guest T4[approximationHelper, approximationHelper, approximationHelper, approximationHelper]
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
			got := Compare4(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal4(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan4(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual4(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan4(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual4(tt.host, tt.guest))
		})
	}
}

func TestT4_CompareC(t *testing.T) {
	lesser := New4(stringComparable("1"), stringComparable("2"), stringComparable("3"), stringComparable("4"))
	greater := New4(stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"))

	tests := []struct {
		name        string
		host, guest T4[stringComparable, stringComparable, stringComparable, stringComparable]
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
			got := Compare4C(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal4C(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan4C(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual4C(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan4C(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual4C(tt.host, tt.guest))
		})
	}
}

func TestT4_EqualE(t *testing.T) {
	a := New4(intEqualable(1), intEqualable(2), intEqualable(3), intEqualable(4))
	b := New4(intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5))

	require.False(t, Equal4E(a, b))
	require.True(t, Equal4E(a, a))
}

func TestT4_String(t *testing.T) {
	tup := New4("1", "2", "3", "4")
	require.Equal(t, `["1" "2" "3" "4"]`, tup.String())
}

func TestT4_GoString(t *testing.T) {
	tup := New4("1", "2", "3", "4")
	require.Equal(t, `tuple.T4[string, string, string, string]{V1: "1", V2: "2", V3: "3", V4: "4"}`, tup.GoString())
}

func TestT4_ToArray(t *testing.T) {
	tup := New4("1", "2", "3", "4")
	require.Equal(t, [4]any{
		"1", "2", "3", "4",
	}, tup.Array())
}

func TestT4_ToSlice(t *testing.T) {
	tup := New4("1", "2", "3", "4")
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
			do := func() T4[string, string, string, string] {
				return FromArray4X[string, string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New4("1", "2", "3", "4"), do())
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
			tup, err := FromArray4[string, string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New4("1", "2", "3", "4"), tup)
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
			do := func() T4[string, string, string, string] {
				return FromSlice4X[string, string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New4("1", "2", "3", "4"), do())
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
			tup, err := FromSlice4[string, string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New4("1", "2", "3", "4"), tup)
		})
	}
}
