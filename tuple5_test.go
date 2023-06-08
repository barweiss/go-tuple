package tuple

import (
	"encoding/json"
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

func TestT5_Compare(t *testing.T) {
	lesser := New5(1, 2, 3, 4, 5)
	greater := New5(2, 3, 4, 5, 6)

	tests := []struct {
		name        string
		host, guest T5[int, int, int, int, int]
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
			got := Compare5(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal5(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan5(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual5(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan5(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual5(tt.host, tt.guest))
		})
	}
}

func TestT5_Compare_Approx(t *testing.T) {
	lesser := New5(approximationHelper("1"), approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"))
	greater := New5(approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"), approximationHelper("6"))

	tests := []struct {
		name        string
		host, guest T5[approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper]
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
			got := Compare5(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal5(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan5(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual5(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan5(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual5(tt.host, tt.guest))
		})
	}
}

func TestT5_CompareC(t *testing.T) {
	lesser := New5(stringComparable("1"), stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"))
	greater := New5(stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"), stringComparable("6"))

	tests := []struct {
		name        string
		host, guest T5[stringComparable, stringComparable, stringComparable, stringComparable, stringComparable]
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
			got := Compare5C(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal5C(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan5C(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual5C(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan5C(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual5C(tt.host, tt.guest))
		})
	}
}

func TestT5_EqualE(t *testing.T) {
	a := New5(intEqualable(1), intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5))
	b := New5(intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5), intEqualable(6))

	require.False(t, Equal5E(a, b))
	require.True(t, Equal5E(a, a))
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

func TestT5_MarshalJSON(t *testing.T) {
	tup := New5("1", "2", "3", "4", "5")

	got, err := json.Marshal(tup)
	require.NoError(t, err)
	require.Equal(t, got, []byte(`["1","2","3","4","5"]`))
}

func TestT5_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    T5[string, string, string, string, string]
		wantErr bool
	}{
		{
			name:    "nil data",
			data:    nil,
			wantErr: true,
		},
		{
			name:    "empty data",
			data:    []byte{},
			wantErr: true,
		},
		{
			name:    "string data",
			data:    []byte(`"hi"`),
			wantErr: true,
		},
		{
			name:    "empty json array",
			data:    []byte(`[]`),
			wantErr: true,
		},
		{
			name:    "longer json array",
			data:    []byte(`["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"]`),
			wantErr: true,
		},
		{
			name:    "json array of invalid types",
			data:    []byte(`[1,2,3,4,5]`),
			wantErr: true,
		},
		{
			name:    "json array with 1 invalid type",
			data:    []byte(`[1,"2","3","4","5"]`),
			wantErr: true,
		},
		{
			name:    "json array of valid types",
			data:    []byte(`["1","2","3","4","5"]`),
			want:    New5("1", "2", "3", "4", "5"),
			wantErr: false,
		},
		{
			name:    "json object of valid types",
			data:    []byte(`{"V1": "1","V2": "2","V3": "3","V4": "4","V5": "5"}`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got T5[string, string, string, string, string]
			err := json.Unmarshal(tt.data, &got)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestT5_Marshal_Unmarshal(t *testing.T) {
	tup := New5("1", "2", "3", "4", "5")

	marshalled, err := json.Marshal(tup)
	require.NoError(t, err)

	var unmarshalled T5[string, string, string, string, string]
	err = json.Unmarshal(marshalled, &unmarshalled)

	require.NoError(t, err)
	require.Equal(t, tup, unmarshalled)
}
