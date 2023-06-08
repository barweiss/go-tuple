package tuple

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

{{/* These variables can be used when the context of dot changes. */}}
{{$indexes := .Indexes}}
{{$len := .Len}}
{{$stringOverload := buildSingleTypedOverload $indexes "string"}}

func TestT{{.Len}}_New(t *testing.T) {
	tup := New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, {{$stringOverload}}{
		{{range .Indexes -}}
		V{{.}}: {{. | quote}},
		{{end}}
	}, tup)
}

func TestT{{.Len}}_Len(t *testing.T) {
	tup := New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, {{.Len}}, tup.Len())
}

func TestT{{.Len}}_Values(t *testing.T) {
	tup := New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}})
	{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}v{{$index}}{{end}} := tup.Values()
	{{range .Indexes -}}
	require.Equal(t, {{. | quote}}, v{{.}})
	{{end -}}
}

func TestT{{.Len}}_Compare(t *testing.T) {
	lesser := New{{.Len}}({{range .Indexes}}{{.}},{{end}})
	greater := New{{.Len}}({{range .Indexes}}{{. | inc}},{{end}})

	tests := []struct{
		name string
		host, guest T{{.Len}}[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}int{{end}}]
		want OrderedComparisonResult
		wantEQ bool
		wantLT bool
		wantLE bool
		wantGT bool
		wantGE bool
	}{
		{
			name: "less than",
			host: lesser,
			guest: greater,
			want: -1,
			wantLT: true,
			wantLE: true,
		},
		{
			name: "greater than",
			host: greater,
			guest: lesser,
			want: 1,
			wantGT: true,
			wantGE: true,
		},
		{
			name: "equal",
			host: lesser,
			guest: lesser,
			want: 0,
			wantEQ: true,
			wantLE: true,
			wantGE: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			got := Compare{{.Len}}(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal{{.Len}}(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan{{.Len}}(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual{{.Len}}(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan{{.Len}}(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual{{.Len}}(tt.host, tt.guest))
		})
	}
}

func TestT{{.Len}}_Compare_Approx(t *testing.T) {
	lesser := New{{.Len}}({{range .Indexes}}approximationHelper({{. | quote}}),{{end}})
	greater := New{{.Len}}({{range .Indexes}}approximationHelper({{. | inc | quote}}),{{end}})

	tests := []struct{
		name string
		host, guest T{{.Len}}[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}approximationHelper{{end}}]
		want OrderedComparisonResult
		wantEQ bool
		wantLT bool
		wantLE bool
		wantGT bool
		wantGE bool
	}{
		{
			name: "less than",
			host: lesser,
			guest: greater,
			want: -1,
			wantLT: true,
			wantLE: true,
		},
		{
			name: "greater than",
			host: greater,
			guest: lesser,
			want: 1,
			wantGT: true,
			wantGE: true,
		},
		{
			name: "equal",
			host: lesser,
			guest: lesser,
			want: 0,
			wantEQ: true,
			wantLE: true,
			wantGE: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			got := Compare{{.Len}}(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal{{.Len}}(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan{{.Len}}(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual{{.Len}}(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan{{.Len}}(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual{{.Len}}(tt.host, tt.guest))
		})
	}
}

func TestT{{.Len}}_CompareC(t *testing.T) {
	lesser := New{{.Len}}({{range .Indexes}}stringComparable({{. | quote}}),{{end}})
	greater := New{{.Len}}({{range .Indexes}}stringComparable({{. | inc | quote}}),{{end}})

	tests := []struct{
		name string
		host, guest T{{.Len}}[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}stringComparable{{end}}]
		want OrderedComparisonResult
		wantEQ bool
		wantLT bool
		wantLE bool
		wantGT bool
		wantGE bool
	}{
		{
			name: "less than",
			host: lesser,
			guest: greater,
			want: -1,
			wantLT: true,
			wantLE: true,
		},
		{
			name: "greater than",
			host: greater,
			guest: lesser,
			want: 1,
			wantGT: true,
			wantGE: true,
		},
		{
			name: "equal",
			host: lesser,
			guest: lesser,
			want: 0,
			wantEQ: true,
			wantLE: true,
			wantGE: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			got := Compare{{.Len}}C(tt.host, tt.guest)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantEQ, got.EQ())
			require.Equal(t, tt.wantLT, got.LT())
			require.Equal(t, tt.wantLE, got.LE())
			require.Equal(t, tt.wantGT, got.GT())
			require.Equal(t, tt.wantGE, got.GE())

			require.Equal(t, tt.wantEQ, Equal{{.Len}}C(tt.host, tt.guest))
			require.Equal(t, tt.wantLT, LessThan{{.Len}}C(tt.host, tt.guest))
			require.Equal(t, tt.wantLE, LessOrEqual{{.Len}}C(tt.host, tt.guest))
			require.Equal(t, tt.wantGT, GreaterThan{{.Len}}C(tt.host, tt.guest))
			require.Equal(t, tt.wantGE, GreaterOrEqual{{.Len}}C(tt.host, tt.guest))
		})
	}
}

func TestT{{.Len}}_EqualE(t *testing.T) {
	a := New{{.Len}}({{range .Indexes}}intEqualable({{.}}),{{end}})
	b := New{{.Len}}({{range .Indexes}}intEqualable({{. | inc}}),{{end}})

	require.False(t, Equal{{.Len}}E(a, b))
	require.True(t, Equal{{.Len}}E(a, a))
}

func TestT{{.Len}}_String(t *testing.T) {
	tup := New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, `[{{range $i, $index := .Indexes}}{{if gt $i 0}} {{end}}{{. | quote}}{{end}}]`, tup.String())
}

func TestT{{.Len}}_GoString(t *testing.T) {
	tup := New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, `tuple.{{$stringOverload}}{
		{{- range $i, $index := .Indexes -}}
		{{- if gt $i 0}}, {{end -}}
		V{{$index}}: {{$index | quote}}
		{{- end -}}
	}`, tup.GoString())
}

func TestT{{.Len}}_ToArray(t *testing.T) {
	tup := New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, [{{.Len}}]any{
		{{range .Indexes -}}{{. | quote}},{{end}}
	}, tup.Array())
}

func TestT{{.Len}}_ToSlice(t *testing.T) {
	tup := New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, []any{
		{{range .Indexes -}}{{. | quote}},{{end}}
	}, tup.Slice())
}

func TestT{{.Len}}_FromArrayX(t *testing.T) {
	tests := []struct{
		name string
		array [{{.Len}}]any
		wantPanic bool
	}{
		{
			name: "all types match",
			array: [{{.Len}}]any{
				{{range .Indexes}}{{. | quote}},{{end}}
			},
			wantPanic: false,
		},
		{{range $testIndex, $index := .Indexes}}
		{
			name: "index {{$index}} bad type",
			array: [{{$len}}]any{
				{{- range $arrayIndex, $elemIndex := $indexes -}}
					{{- if eq $testIndex $arrayIndex -}}
						{{$arrayIndex}},
					{{- else -}}
						{{$arrayIndex | quote}},
					{{- end -}}
				{{- end -}}
			},
			wantPanic: true,
		},
		{{end}}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func () {{$stringOverload}} {
				return FromArray{{.Len}}X[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}string{{end}}](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func () {
					_ = do()
				})
				return
			}

			require.Equal(t, New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}}), do())
		})
	}
}

func TestT{{.Len}}_FromArray(t *testing.T) {
	tests := []struct{
		name string
		array [{{.Len}}]any
		wantErr bool
	}{
		{
			name: "all types match",
			array: [{{.Len}}]any{
				{{range .Indexes}}{{. | quote}},{{end}}
			},
			wantErr: false,
		},
		{{range $testIndex, $index := .Indexes}}
		{
			name: "index {{$index}} bad type",
			array: [{{$len}}]any{
				{{- range $arrayIndex, $elemIndex := $indexes -}}
					{{- if eq $testIndex $arrayIndex -}}
						{{$elemIndex}},
					{{- else -}}
						{{$elemIndex | quote}},
					{{- end -}}
				{{- end -}}
			},
			wantErr: true,
		},
		{{end}}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromArray{{.Len}}[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}string{{end}}](tt.array)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}}), tup)
		})
	}
}

func TestT{{.Len}}_FromSliceX(t *testing.T) {
	tests := []struct{
		name string
		slice []any
		wantPanic bool
	}{
		{
			name: "all types match",
			slice: []any{
				{{range .Indexes -}}{{. | quote}},{{end}}
			},
			wantPanic: false,
		},
		{
			name: "slice empty",
			slice: []any{},
			wantPanic: true,
		},
		{
			name: "slice too short",
			slice: []any{
				{{range $i, $unused := .Indexes}}{{if gt $i 1}}{{. | quote}},{{end}}{{end}}
			},
			wantPanic: true,
		},
		{
			name: "slice too long",
			slice: []any{
				{{range .Indexes}}{{. | quote}},{{end}}
				"extra",
			},
			wantPanic: true,
		},
		{{range $testIndex, $index := .Indexes}}
		{
			name: "index {{$index}} bad type",
			slice: []any{
				{{- range $arrayIndex, $elemIndex := $indexes -}}
					{{- if eq $testIndex $arrayIndex -}}
						{{$arrayIndex}},
					{{- else -}}
						{{$arrayIndex | quote}},
					{{- end -}}
				{{- end -}}
			},
			wantPanic: true,
		},
		{{end}}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do := func () {{$stringOverload}} {
				return FromSlice{{.Len}}X[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}string{{end}}](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func () {
					_ = do()
				})
				return
			}

			require.Equal(t, New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}}), do())
		})
	}
}

func TestT{{.Len}}_FromSlice(t *testing.T) {
	tests := []struct{
		name string
		slice []any
		wantErr bool
	}{
		{
			name: "all types match",
			slice: []any{
				{{range .Indexes}}{{. | quote}},{{end}}
			},
			wantErr: false,
		},
		{
			name: "slice empty",
			slice: []any{},
			wantErr: true,
		},
		{
			name: "slice too short",
			slice: []any{
				{{range $i, $unused := .Indexes}}{{if gt $i 1}}{{. | quote}},{{end}}{{end}}
			},
			wantErr: true,
		},
		{
			name: "slice too long",
			slice: []any{
				{{range .Indexes}}{{. | quote}},{{end}}
				"extra",
			},
			wantErr: true,
		},
		{{range $testIndex, $index := .Indexes}}
		{
			name: "index {{$index}} bad type",
			slice: []any{
				{{- range $arrayIndex, $elemIndex := $indexes -}}
					{{- if eq $testIndex $arrayIndex -}}
						{{$elemIndex}},
					{{- else -}}
						{{$elemIndex | quote}},
					{{- end -}}
				{{- end -}}
			},
			wantErr: true,
		},
		{{end}}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tup, err := FromSlice{{.Len}}[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}string{{end}}](tt.slice)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}}), tup)
		})
	}
}

func TestT{{.Len}}_MarshalJSON(t *testing.T) {
	tup := New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}})

	got, err := json.Marshal(tup)
	require.NoError(t, err)
	require.Equal(t, got, []byte(`[{{range .Indexes}}{{if ne . 1}},{{end}}{{. | quote}}{{end}}]`))
}

func TestT{{.Len}}_UnmarshalJSON(t *testing.T) {
	tests := []struct{
		name string
		data []byte
		want {{$stringOverload}}
		wantErr bool
	}{
		{
			name: "nil data",
			data: nil,
			wantErr: true,
		},
		{
			name: "empty data",
			data: []byte{},
			wantErr: true,
		},
		{
			name: "string data",
			data: []byte(`"hi"`),
			wantErr: true,
		},
		{
			name: "empty json array",
			data: []byte(`[]`),
			wantErr: true,
		},
		{
			name: "longer json array",
			data: []byte(`["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"]`),
			wantErr: true,
		},
		{{range $invalidIndex, $_ := .Indexes}}
		{
			name: "json array with invalid type at index {{$invalidIndex}}",
			data: []byte(`[{{range $currentIndex, $num := $.Indexes}}{{if ne $currentIndex 0}},{{end}}{{if eq $currentIndex $invalidIndex}}{{$num}}{{else}}{{$num | quote}}{{end}}{{end}}]`),
			wantErr: true,
		},
		{{- end}}
		{
			name: "json array of valid types",
			data: []byte(`[{{range $.Indexes}}{{if ne . 1}},{{end}}{{. | quote}}{{end}}]`),
			want: New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}}),
			wantErr: false,
		},
		{
			name: "json object of valid types",
			data: []byte(`{{"{"}}{{range $.Indexes}}{{if ne . 1}},{{end}}"V{{.}}": {{. | quote}}{{end}}{{"}"}}`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got {{$stringOverload}}
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

func TestT{{.Len}}_Unmarshal_CustomStruct(t *testing.T) {
	type Custom struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}

	want := New{{.Len}}({{range .Indexes}}Custom{ Name: {{. | quote}}, Age: {{.}} },{{end}})
	var got T{{.Len}}[{{range .Indexes}}Custom,{{end}}]
	err := json.Unmarshal([]byte(`[
		{{- range .Indexes -}}
		{{- if ne . 1}},{{end}}
		{ "name": {{. | quote}}, "age": {{.}} }
		{{- end}}
	]`), &got)

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestT{{.Len}}_Marshal_Unmarshal(t *testing.T) {
	tup := New{{.Len}}({{range .Indexes}}{{. | quote}},{{end}})

	marshalled, err := json.Marshal(tup)
	require.NoError(t, err)

	var unmarshalled {{$stringOverload}}
	err = json.Unmarshal(marshalled, &unmarshalled)

	require.NoError(t, err)
	require.Equal(t, tup, unmarshalled)
}
