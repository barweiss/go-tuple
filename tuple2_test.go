package tuple

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT2_New(t *testing.T) {
	tup := New2("1", "2")
	require.Equal(t, T2[string, string]{
		V1: "1",
		V2: "2",
	}, tup)
}

func TestT2_Len(t *testing.T) {
	tup := New2("1", "2")
	require.Equal(t, 2, tup.Len())
}

func TestT2_Values(t *testing.T) {
	tup := New2("1", "2")
	v1, v2 := tup.Values()
	require.Equal(t, "1", v1)
	require.Equal(t, "2", v2)
}

func TestT2_Compare(t *testing.T) {
	lesser := New2(1, 2)
	greater := New2(2, 3)

	tests := []struct {
		name        string
		host, guest T2[int, int]
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
			got := Compare2(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal2(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan2(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual2(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan2(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual2(tt.host, tt.guest))
		})
	}
}

func TestT2_Compare_Approx(t *testing.T) {
	lesser := New2(approximationHelper("1"), approximationHelper("2"))
	greater := New2(approximationHelper("2"), approximationHelper("3"))

	tests := []struct {
		name        string
		host, guest T2[approximationHelper, approximationHelper]
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
			got := Compare2(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal2(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan2(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual2(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan2(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual2(tt.host, tt.guest))
		})
	}
}

func TestT2_CompareC(t *testing.T) {
	lesser := New2(stringComparable("1"), stringComparable("2"))
	greater := New2(stringComparable("2"), stringComparable("3"))

	tests := []struct {
		name        string
		host, guest T2[stringComparable, stringComparable]
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
			got := Compare2C(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal2C(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan2C(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual2C(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan2C(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual2C(tt.host, tt.guest))
		})
	}
}

func TestT2_EqualE(t *testing.T) {
	a := New2(intEqualable(1), intEqualable(2))
	b := New2(intEqualable(2), intEqualable(3))

	require.False(t, Equal2E(a, b))
	require.True(t, Equal2E(a, a))
}

func TestT2_String(t *testing.T) {
	tup := New2("1", "2")
	require.Equal(t, `["1" "2"]`, tup.String())
}

func TestT2_GoString(t *testing.T) {
	tup := New2("1", "2")
	require.Equal(t, `tuple.T2[string, string]{V1: "1", V2: "2"}`, tup.GoString())
}

func TestT2_ToArray(t *testing.T) {
	tup := New2("1", "2")
	require.Equal(t, [2]any{
		"1", "2",
	}, tup.Array())
}

func TestT2_ToSlice(t *testing.T) {
	tup := New2("1", "2")
	require.Equal(t, []any{
		"1", "2",
	}, tup.Slice())
}

func TestT2_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [2]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [2]any{
				"1", "2",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [2]any{0, "1"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			array:     [2]any{"0", 1},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T2[string, string] {
				return FromArray2X[string, string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New2("1", "2"), do())
		})
	}
}

func TestT2_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [2]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [2]any{
				"1", "2",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [2]any{1, "2"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			array:   [2]any{"1", 2},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromArray2[string, string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New2("1", "2"), tup)
		})
	}
}

func TestT2_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2",
			},
			wantPanic: false,
		},
		{
			name:      "slice empty",
			slice:     []any{},
			wantPanic: true,
		},
		{
			name:      "slice too short",
			slice:     []any{},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0, "1"},
			wantPanic: true,
		},

		{
			name:      "index 2 bad type",
			slice:     []any{"0", 1},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T2[string, string] {
				return FromSlice2X[string, string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New2("1", "2"), do())
		})
	}
}

func TestT2_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1", "2",
			},
			wantErr: false,
		},
		{
			name:    "slice empty",
			slice:   []any{},
			wantErr: true,
		},
		{
			name:    "slice too short",
			slice:   []any{},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				"1", "2",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1, "2"},
			wantErr: true,
		},

		{
			name:    "index 2 bad type",
			slice:   []any{"1", 2},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromSlice2[string, string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New2("1", "2"), tup)
		})
	}
}

func TestT2_MarshalJSON(t *testing.T) {
	tup := New2("1", "2")

	got, err := json.Marshal(tup)
	require.NoError(t, err)
	require.Equal(t, got, []byte(`["1","2"]`))
}

func TestT2_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    T2[string, string]
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
			data:    []byte(`[1,2]`),
			wantErr: true,
		},
		{
			name:    "json array with 1 invalid type",
			data:    []byte(`[1,"2"]`),
			wantErr: true,
		},
		{
			name:    "json array of valid types",
			data:    []byte(`["1","2"]`),
			want:    New2("1", "2"),
			wantErr: false,
		},
		{
			name:    "json object of valid types",
			data:    []byte(`{"V1": "1","V2": "2"}`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got T2[string, string]
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

func TestT2_Unmarshal_CustomStruct(t *testing.T) {
	type Custom struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	want := New2(Custom{Name: "1", Age: 1}, Custom{Name: "2", Age: 2})
	var got T2[Custom, Custom]
	err := json.Unmarshal([]byte(`[
		{ "name": "1", "age": 1 },
		{ "name": "2", "age": 2 }
	]`), &got)

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestT2_Marshal_Unmarshal(t *testing.T) {
	tup := New2("1", "2")

	marshalled, err := json.Marshal(tup)
	require.NoError(t, err)

	var unmarshalled T2[string, string]
	err = json.Unmarshal(marshalled, &unmarshalled)

	require.NoError(t, err)
	require.Equal(t, tup, unmarshalled)
}
