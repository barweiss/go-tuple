package tuple

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT3_New(t *testing.T) {
	tup := New3("1", "2", "3")
	require.Equal(t, T3[string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
	}, tup)
}

func TestT3_Len(t *testing.T) {
	tup := New3("1", "2", "3")
	require.Equal(t, 3, tup.Len())
}

func TestT3_Values(t *testing.T) {
	tup := New3("1", "2", "3")
	v1, v2, v3 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
}

func TestT3_Compare(t *testing.T) {
	lesser := New3(1, 2, 3)
	greater := New3(2, 3, 4)

	tests := []struct {
		name        string
		host, guest T3[int, int, int]
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
			got := Compare3(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal3(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan3(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual3(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan3(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual3(tt.host, tt.guest))
		})
	}
}

func TestT3_Compare_Approx(t *testing.T) {
	lesser := New3(approximationHelper("1"), approximationHelper("2"), approximationHelper("3"))
	greater := New3(approximationHelper("2"), approximationHelper("3"), approximationHelper("4"))

	tests := []struct {
		name        string
		host, guest T3[approximationHelper, approximationHelper, approximationHelper]
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
			got := Compare3(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal3(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan3(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual3(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan3(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual3(tt.host, tt.guest))
		})
	}
}

func TestT3_CompareC(t *testing.T) {
	lesser := New3(stringComparable("1"), stringComparable("2"), stringComparable("3"))
	greater := New3(stringComparable("2"), stringComparable("3"), stringComparable("4"))

	tests := []struct {
		name        string
		host, guest T3[stringComparable, stringComparable, stringComparable]
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
			got := Compare3C(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal3C(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan3C(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual3C(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan3C(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual3C(tt.host, tt.guest))
		})
	}
}

func TestT3_EqualE(t *testing.T) {
	a := New3(intEqualable(1), intEqualable(2), intEqualable(3))
	b := New3(intEqualable(2), intEqualable(3), intEqualable(4))

	require.False(t, Equal3E(a, b))
	require.True(t, Equal3E(a, a))
}

func TestT3_String(t *testing.T) {
	tup := New3("1", "2", "3")
	require.Equal(t, `["1" "2" "3"]`, tup.String())
}

func TestT3_GoString(t *testing.T) {
	tup := New3("1", "2", "3")
	require.Equal(t, `tuple.T3[string, string, string]{V1: "1", V2: "2", V3: "3"}`, tup.GoString())
}

func TestT3_ToArray(t *testing.T) {
	tup := New3("1", "2", "3")
	require.Equal(t, [3]any{
		"1", "2", "3",
	}, tup.Array())
}

func TestT3_ToSlice(t *testing.T) {
	tup := New3("1", "2", "3")
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
			do := func() T3[string, string, string] {
				return FromArray3X[string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New3("1", "2", "3"), do())
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
			tup, err := FromArray3[string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New3("1", "2", "3"), tup)
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
			do := func() T3[string, string, string] {
				return FromSlice3X[string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New3("1", "2", "3"), do())
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
			tup, err := FromSlice3[string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New3("1", "2", "3"), tup)
		})
	}
}
