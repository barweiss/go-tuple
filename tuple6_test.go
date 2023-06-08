package tuple

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT6_New(t *testing.T) {
	tup := New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, T6[string, string, string, string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
		V4: "4",
		V5: "5",
		V6: "6",
	}, tup)
}

func TestT6_Len(t *testing.T) {
	tup := New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, 6, tup.Len())
}

func TestT6_Values(t *testing.T) {
	tup := New6("1", "2", "3", "4", "5", "6")
	v1, v2, v3, v4, v5, v6 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
	require.Equal(t, "4", v4)
	require.Equal(t, "5", v5)
	require.Equal(t, "6", v6)
}

func TestT6_Compare(t *testing.T) {
	lesser := New6(1, 2, 3, 4, 5, 6)
	greater := New6(2, 3, 4, 5, 6, 7)

	tests := []struct {
		name        string
		host, guest T6[int, int, int, int, int, int]
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
			got := Compare6(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal6(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan6(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual6(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan6(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual6(tt.host, tt.guest))
		})
	}
}

func TestT6_Compare_Approx(t *testing.T) {
	lesser := New6(approximationHelper("1"), approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"), approximationHelper("6"))
	greater := New6(approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"), approximationHelper("6"), approximationHelper("7"))

	tests := []struct {
		name        string
		host, guest T6[approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper]
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
			got := Compare6(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal6(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan6(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual6(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan6(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual6(tt.host, tt.guest))
		})
	}
}

func TestT6_CompareC(t *testing.T) {
	lesser := New6(stringComparable("1"), stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"), stringComparable("6"))
	greater := New6(stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"), stringComparable("6"), stringComparable("7"))

	tests := []struct {
		name        string
		host, guest T6[stringComparable, stringComparable, stringComparable, stringComparable, stringComparable, stringComparable]
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
			got := Compare6C(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal6C(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan6C(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual6C(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan6C(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual6C(tt.host, tt.guest))
		})
	}
}

func TestT6_EqualE(t *testing.T) {
	a := New6(intEqualable(1), intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5), intEqualable(6))
	b := New6(intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5), intEqualable(6), intEqualable(7))

	require.False(t, Equal6E(a, b))
	require.True(t, Equal6E(a, a))
}

func TestT6_String(t *testing.T) {
	tup := New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, `["1" "2" "3" "4" "5" "6"]`, tup.String())
}

func TestT6_GoString(t *testing.T) {
	tup := New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, `tuple.T6[string, string, string, string, string, string]{V1: "1", V2: "2", V3: "3", V4: "4", V5: "5", V6: "6"}`, tup.GoString())
}

func TestT6_ToArray(t *testing.T) {
	tup := New6("1", "2", "3", "4", "5", "6")
	require.Equal(t, [6]any{
		"1", "2", "3", "4", "5", "6",
	}, tup.Array())
}

func TestT6_ToSlice(t *testing.T) {
	tup := New6("1", "2", "3", "4", "5", "6")
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
			do := func() T6[string, string, string, string, string, string] {
				return FromArray6X[string, string, string, string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New6("1", "2", "3", "4", "5", "6"), do())
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
			tup, err := FromArray6[string, string, string, string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New6("1", "2", "3", "4", "5", "6"), tup)
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
			do := func() T6[string, string, string, string, string, string] {
				return FromSlice6X[string, string, string, string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New6("1", "2", "3", "4", "5", "6"), do())
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
			tup, err := FromSlice6[string, string, string, string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New6("1", "2", "3", "4", "5", "6"), tup)
		})
	}
}

func TestT6_MarshalJSON(t *testing.T) {
	tup := New6("1", "2", "3", "4", "5", "6")

	got, err := json.Marshal(tup)
	require.NoError(t, err)
	require.Equal(t, got, []byte(`["1","2","3","4","5","6"]`))
}

func TestT6_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    T6[string, string, string, string, string, string]
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
			data:    []byte(`[1,2,3,4,5,6]`),
			wantErr: true,
		},
		{
			name:    "json array with 1 invalid type",
			data:    []byte(`[1,"2","3","4","5","6"]`),
			wantErr: true,
		},
		{
			name:    "json array of valid types",
			data:    []byte(`["1","2","3","4","5","6"]`),
			want:    New6("1", "2", "3", "4", "5", "6"),
			wantErr: false,
		},
		{
			name:    "json object of valid types",
			data:    []byte(`{"V1": "1","V2": "2","V3": "3","V4": "4","V5": "5","V6": "6"}`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got T6[string, string, string, string, string, string]
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

func TestT6_Marshal_Unmarshal(t *testing.T) {
	tup := New6("1", "2", "3", "4", "5", "6")

	marshalled, err := json.Marshal(tup)
	require.NoError(t, err)

	var unmarshalled T6[string, string, string, string, string, string]
	err = json.Unmarshal(marshalled, &unmarshalled)

	require.NoError(t, err)
	require.Equal(t, tup, unmarshalled)
}
