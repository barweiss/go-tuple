package tuple

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT1_New(t *testing.T) {
	tup := New1("1")
	require.Equal(t, T1[string]{
		V1: "1",
	}, tup)
}

func TestT1_Len(t *testing.T) {
	tup := New1("1")
	require.Equal(t, 1, tup.Len())
}

func TestT1_Values(t *testing.T) {
	tup := New1("1")
	v1 := tup.Values()
	require.Equal(t, "1", v1)
}

func TestT1_Compare(t *testing.T) {
	lesser := New1(1)
	greater := New1(2)

	tests := []struct {
		name        string
		host, guest T1[int]
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
			got := Compare1(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal1(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan1(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual1(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan1(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual1(tt.host, tt.guest))
		})
	}
}

func TestT1_Compare_Approx(t *testing.T) {
	lesser := New1(approximationHelper("1"))
	greater := New1(approximationHelper("2"))

	tests := []struct {
		name        string
		host, guest T1[approximationHelper]
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
			got := Compare1(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal1(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan1(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual1(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan1(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual1(tt.host, tt.guest))
		})
	}
}

func TestT1_CompareC(t *testing.T) {
	lesser := New1(stringComparable("1"))
	greater := New1(stringComparable("2"))

	tests := []struct {
		name        string
		host, guest T1[stringComparable]
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
			got := Compare1C(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal1C(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan1C(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual1C(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan1C(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual1C(tt.host, tt.guest))
		})
	}
}

func TestT1_EqualE(t *testing.T) {
	a := New1(intEqualable(1))
	b := New1(intEqualable(2))

	require.False(t, Equal1E(a, b))
	require.True(t, Equal1E(a, a))
}

func TestT1_String(t *testing.T) {
	tup := New1("1")
	require.Equal(t, `["1"]`, tup.String())
}

func TestT1_GoString(t *testing.T) {
	tup := New1("1")
	require.Equal(t, `tuple.T1[string]{V1: "1"}`, tup.GoString())
}

func TestT1_ToArray(t *testing.T) {
	tup := New1("1")
	require.Equal(t, [1]any{
		"1",
	}, tup.Array())
}

func TestT1_ToSlice(t *testing.T) {
	tup := New1("1")
	require.Equal(t, []any{
		"1",
	}, tup.Slice())
}

func TestT1_FromArrayX(t *testing.T) {
	tests := []struct {
		name      string
		array     [1]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [1]any{
				"1",
			},
			wantPanic: false,
		},

		{
			name:      "index 1 bad type",
			array:     [1]any{0},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T1[string] {
				return FromArray1X[string](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New1("1"), do())
		})
	}
}

func TestT1_FromArray(t *testing.T) {
	tests := []struct {
		name    string
		array   [1]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [1]any{
				"1",
			},
			wantErr: false,
		},

		{
			name:    "index 1 bad type",
			array:   [1]any{1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromArray1[string](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New1("1"), tup)
		})
	}
}

func TestT1_FromSliceX(t *testing.T) {
	tests := []struct {
		name      string
		slice     []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1",
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
				"1",
				"extra",
			},
			wantPanic: true,
		},

		{
			name:      "index 1 bad type",
			slice:     []any{0},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func() T1[string] {
				return FromSlice1X[string](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func() {
					_ = do()
				})
				return
			}

			require.Equal(t, New1("1"), do())
		})
	}
}

func TestT1_FromSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				"1",
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
				"1",
				"extra",
			},
			wantErr: true,
		},

		{
			name:    "index 1 bad type",
			slice:   []any{1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromSlice1[string](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New1("1"), tup)
		})
	}
}

func TestT1_MarshalJSON(t *testing.T) {
	tup := New1("1")

	got, err := json.Marshal(tup)
	require.NoError(t, err)
	require.Equal(t, got, []byte(`["1"]`))
}

func TestT1_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    T1[string]
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
			data:    []byte(`[1]`),
			wantErr: true,
		},

		{
			name:    "json array of valid types",
			data:    []byte(`["1"]`),
			want:    New1("1"),
			wantErr: false,
		},
		{
			name:    "json object of valid types",
			data:    []byte(`{"V1": "1"}`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got T1[string]
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

func TestT1_Marshal_Unmarshal(t *testing.T) {
	tup := New1("1")

	marshalled, err := json.Marshal(tup)
	require.NoError(t, err)

	var unmarshalled T1[string]
	err = json.Unmarshal(marshalled, &unmarshalled)

	require.NoError(t, err)
	require.Equal(t, tup, unmarshalled)
}
