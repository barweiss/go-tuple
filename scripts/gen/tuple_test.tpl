package tuple

import (
	"testing"

	"github.com/stretchr/testify/require"
)

{{/* These variables can be used when the context of dot changes. */}}
{{$indexes := .Indexes}}
{{$len := .Len}}

func TestT{{.Len}}_New(t *testing.T) {
	tup := New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, T{{.Len}}[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}string{{end}}]{
		{{range .Indexes -}}
		V{{.}}: {{. | quote}},
		{{end}}
	}, tup)
}

func TestT{{.Len}}_Len(t *testing.T) {
	tup := New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, {{.Len}}, tup.Len())
}

func TestT{{.Len}}_Values(t *testing.T) {
	tup := New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}})
	{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}v{{$index}}{{end}} := tup.Values()
	{{range .Indexes -}}
	require.Equal(t, {{. | quote}}, v{{.}})
	{{end -}}
}

func TestT{{.Len}}_String(t *testing.T) {
	tup := New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, `[{{range $i, $index := .Indexes}}{{if gt $i 0}} {{end}}{{. | quote}}{{end}}]`, tup.String())
}

func TestT{{.Len}}_GoString(t *testing.T) {
	tup := New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, `tuple.T{{.Len}}[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}string{{end}}]{
		{{- range $i, $index := .Indexes -}}
		{{- if gt $i 0}}, {{end -}}
		V{{$index}}: {{$index | quote}}
		{{- end -}}
	}`, tup.GoString())
}

func TestT{{.Len}}_ToArray(t *testing.T) {
	tup := New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}})
	require.Equal(t, [{{.Len}}]any{
		{{range .Indexes -}}{{. | quote}},{{end}}
	}, tup.Array())
}

func TestT{{.Len}}_ToSlice(t *testing.T) {
	tup := New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}})
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
			do := func () T{{.Len}}[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}string{{end}}] {
				return FromArray{{.Len}}X[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}string{{end}}](tt.array)
			}

			if tt.wantPanic {
				require.Panics(t, func () {
					_ = do()
				})
				return
			}

			require.Equal(t, New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}}), do())
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
			require.Equal(t, New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}}), tup)
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
			do := func () T{{.Len}}[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}string{{end}}] {
				return FromSlice{{.Len}}X[{{range $i, $index := .Indexes}}{{if gt $i 0}}, {{end}}string{{end}}](tt.slice)
			}

			if tt.wantPanic {
				require.Panics(t, func () {
					_ = do()
				})
				return
			}

			require.Equal(t, New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}}), do())
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
			require.Equal(t, New{{len .Indexes}}({{range .Indexes}}{{. | quote}},{{end}}), tup)
		})
	}
}
