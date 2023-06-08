package tuple

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT7_New(t *testing.T) {
	tup := New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, T7[string, string, string, string, string, string, string]{
		V1: "1",
		V2: "2",
		V3: "3",
		V4: "4",
		V5: "5",
		V6: "6",
		V7: "7",
	}, tup)
}

func TestT7_Len(t *testing.T) {
	tup := New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, 7, tup.Len())
}

func TestT7_Values(t *testing.T) {
	tup := New7("1", "2", "3", "4", "5", "6", "7")
	v1, v2, v3, v4, v5, v6, v7 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
	require.Equal(t, "3", v3)
	require.Equal(t, "4", v4)
	require.Equal(t, "5", v5)
	require.Equal(t, "6", v6)
	require.Equal(t, "7", v7)
}

func TestT7_Compare(t *testing.T) {
	lesser := New7(1, 2, 3, 4, 5, 6, 7)
	greater := New7(2, 3, 4, 5, 6, 7, 8)

	tests := []struct {
		name        string
		host, guest T7[int, int, int, int, int, int, int]
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
			got := Compare7(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal7(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan7(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual7(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan7(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual7(tt.host, tt.guest))
		})
	}
}

func TestT7_Compare_Approx(t *testing.T) {
	lesser := New7(approximationHelper("1"), approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"), approximationHelper("6"), approximationHelper("7"))
	greater := New7(approximationHelper("2"), approximationHelper("3"), approximationHelper("4"), approximationHelper("5"), approximationHelper("6"), approximationHelper("7"), approximationHelper("8"))

	tests := []struct {
		name        string
		host, guest T7[approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper, approximationHelper]
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
			got := Compare7(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal7(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan7(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual7(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan7(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual7(tt.host, tt.guest))
		})
	}
}

func TestT7_CompareC(t *testing.T) {
	lesser := New7(stringComparable("1"), stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"), stringComparable("6"), stringComparable("7"))
	greater := New7(stringComparable("2"), stringComparable("3"), stringComparable("4"), stringComparable("5"), stringComparable("6"), stringComparable("7"), stringComparable("8"))

	tests := []struct {
		name        string
		host, guest T7[stringComparable, stringComparable, stringComparable, stringComparable, stringComparable, stringComparable, stringComparable]
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
			got := Compare7C(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal7C(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan7C(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual7C(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan7C(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual7C(tt.host, tt.guest))
		})
	}
}

func TestT7_EqualE(t *testing.T) {
	a := New7(intEqualable(1), intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5), intEqualable(6), intEqualable(7))
	b := New7(intEqualable(2), intEqualable(3), intEqualable(4), intEqualable(5), intEqualable(6), intEqualable(7), intEqualable(8))

	require.False(t, Equal7E(a, b))
	require.True(t, Equal7E(a, a))
}

func TestT7_String(t *testing.T) {
	tup := New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, `["1" "2" "3" "4" "5" "6" "7"]`, tup.String())
}

func TestT7_GoString(t *testing.T) {
	tup := New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, `tuple.T7[string, string, string, string, string, string, string]{V1: "1", V2: "2", V3: "3", V4: "4", V5: "5", V6: "6", V7: "7"}`, tup.GoString())
}

func TestT7_ToArray(t *testing.T) {
	tup := New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, [7]any{
		"1", "2", "3", "4", "5", "6", "7",
	}, tup.Array())
}

func TestT7_ToSlice(t *testing.T) {
	tup := New7("1", "2", "3", "4", "5", "6", "7")
	require.Equal(t, []any{
		"1", "2", "3", "4", "5", "6", "7",
	}, tup.Slice())
}

func TestT7_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [7]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [7]any{
				"1", "2", "3", "4", "5", "6", "7",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [7]any{0, "1", "2", "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			array:     [7]any{"0", 1, "2", "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			array:     [7]any{"0", "1", 2, "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			array:     [7]any{"0", "1", "2", 3, "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			array:     [7]any{"0", "1", "2", "3", 4, "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 6 bad type",
			array:     [7]any{"0", "1", "2", "3", "4", 5, "6"},
			wantPanic: true,
		},

		{
			name:      "index 7 bad type",
			array:     [7]any{"0", "1", "2", "3", "4", "5", 6},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T7[string, string, string, string, string, string, string] {
				return FromArray7X[string, string, string, string, string, string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New7("1", "2", "3", "4", "5", "6", "7"), do())
		})
	}
}

func TestT7_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [7]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [7]any{
				"1", "2", "3", "4", "5", "6", "7",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [7]any{1, "2", "3", "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			array:   [7]any{"1", 2, "3", "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			array:   [7]any{"1", "2", 3, "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			array:   [7]any{"1", "2", "3", 4, "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			array:   [7]any{"1", "2", "3", "4", 5, "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 6 bad type",
			array:   [7]any{"1", "2", "3", "4", "5", 6, "7"},
			wantErr: true,
		},

		{
			name:    "index 7 bad type",
			array:   [7]any{"1", "2", "3", "4", "5", "6", 7},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromArray7[string, string, string, string, string, string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New7("1", "2", "3", "4", "5", "6", "7"), tup)
		})
	}
}

func TestT7_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7",
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
				"3", "4", "5", "6", "7",
			},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0, "1", "2", "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			slice:     []any{"0", 1, "2", "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 3 bad type",
			slice:     []any{"0", "1", 2, "3", "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 4 bad type",
			slice:     []any{"0", "1", "2", 3, "4", "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 5 bad type",
			slice:     []any{"0", "1", "2", "3", 4, "5", "6"},
			wantPanic: true,
		},

		{
			name:      "index 6 bad type",
			slice:     []any{"0", "1", "2", "3", "4", 5, "6"},
			wantPanic: true,
		},

		{
			name:      "index 7 bad type",
			slice:     []any{"0", "1", "2", "3", "4", "5", 6},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T7[string, string, string, string, string, string, string] {
				return FromSlice7X[string, string, string, string, string, string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New7("1", "2", "3", "4", "5", "6", "7"), do())
		})
	}
}

func TestT7_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7",
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
				"3", "4", "5", "6", "7",
			},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2", "3", "4", "5", "6", "7",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1, "2", "3", "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			slice:   []any{"1", 2, "3", "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 3 bad type",
			slice:   []any{"1", "2", 3, "4", "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 4 bad type",
			slice:   []any{"1", "2", "3", 4, "5", "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 5 bad type",
			slice:   []any{"1", "2", "3", "4", 5, "6", "7"},
			wantErr: true,
		},

		{
			name:    "index 6 bad type",
			slice:   []any{"1", "2", "3", "4", "5", 6, "7"},
			wantErr: true,
		},

		{
			name:    "index 7 bad type",
			slice:   []any{"1", "2", "3", "4", "5", "6", 7},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromSlice7[string, string, string, string, string, string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New7("1", "2", "3", "4", "5", "6", "7"), tup)
		})
	}
}

func TestT7_MarshalJSON(t *testing.T) {
	tup := New7("1", "2", "3", "4", "5", "6", "7")

	got, err := json.Marshal(tup)
	require.NoError(t, err)
	require.Equal(t, got, []byte(`["1","2","3","4","5","6","7"]`))
}

func TestT7_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    T7[string, string, string, string, string, string, string]
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
			name:    "json array with invalid type at index 0",
			data:    []byte(`[1,"2","3","4","5","6","7"]`),
			wantErr: true,
		},
		{
			name:    "json array with invalid type at index 1",
			data:    []byte(`["1",2,"3","4","5","6","7"]`),
			wantErr: true,
		},
		{
			name:    "json array with invalid type at index 2",
			data:    []byte(`["1","2",3,"4","5","6","7"]`),
			wantErr: true,
		},
		{
			name:    "json array with invalid type at index 3",
			data:    []byte(`["1","2","3",4,"5","6","7"]`),
			wantErr: true,
		},
		{
			name:    "json array with invalid type at index 4",
			data:    []byte(`["1","2","3","4",5,"6","7"]`),
			wantErr: true,
		},
		{
			name:    "json array with invalid type at index 5",
			data:    []byte(`["1","2","3","4","5",6,"7"]`),
			wantErr: true,
		},
		{
			name:    "json array with invalid type at index 6",
			data:    []byte(`["1","2","3","4","5","6",7]`),
			wantErr: true,
		},
		{
			name:    "json array of valid types",
			data:    []byte(`["1","2","3","4","5","6","7"]`),
			want:    New7("1", "2", "3", "4", "5", "6", "7"),
			wantErr: false,
		},
		{
			name:    "json object of valid types",
			data:    []byte(`{"V1": "1","V2": "2","V3": "3","V4": "4","V5": "5","V6": "6","V7": "7"}`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got T7[string, string, string, string, string, string, string]
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

func TestT7_Unmarshal_CustomStruct(t *testing.T) {
	type Custom struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	want := New7(Custom{Name: "1", Age: 1}, Custom{Name: "2", Age: 2}, Custom{Name: "3", Age: 3}, Custom{Name: "4", Age: 4}, Custom{Name: "5", Age: 5}, Custom{Name: "6", Age: 6}, Custom{Name: "7", Age: 7})
	var got T7[Custom, Custom, Custom, Custom, Custom, Custom, Custom]
	err := json.Unmarshal([]byte(`[
		{ "name": "1", "age": 1 },
		{ "name": "2", "age": 2 },
		{ "name": "3", "age": 3 },
		{ "name": "4", "age": 4 },
		{ "name": "5", "age": 5 },
		{ "name": "6", "age": 6 },
		{ "name": "7", "age": 7 }
	]`), &got)

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestT7_Marshal_Unmarshal(t *testing.T) {
	tup := New7("1", "2", "3", "4", "5", "6", "7")

	marshalled, err := json.Marshal(tup)
	require.NoError(t, err)

	var unmarshalled T7[string, string, string, string, string, string, string]
	err = json.Unmarshal(marshalled, &unmarshalled)

	require.NoError(t, err)
	require.Equal(t, tup, unmarshalled)
}
