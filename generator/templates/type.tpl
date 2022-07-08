// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `{{ .Name | fc }}`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `{{ .Name | lc }}`

// {{ .Name | fc }}: {{ .Description }}
package {{ .Name | lc }}

import (
	"encoding/json"
)

// {{ .Name | fc }}Base is the base struct with the default getters and setters
type {{ .Name | fc }}Base struct {
	{{- range .Fields }}
	{{ .Name | lc }} {{ if .Multiple }}[]{{ end }}{{ .Type | typeConvert }} 
	{{- if .Description }} // {{ .Description }}{{ end }}
	{{- end }}

	error error
	self *{{ .Name | fc }}
}

// {{ .Name | lc }}DTO to marshal in JSON
type {{ .Name | lc }}DTO struct {
	{{- range .Fields }}
	{{ .Name | fc }} {{if .Multiple }}[]{{ end }}{{ .Type | typeConvert }} `json:"{{ .Name | lc }}"`
	{{- end }}
}

// -------------------------------------------

// New{{ .Name | fc }} returns a pointer to a new New{{ .Name | fc }} object
func New{{ .Name | fc }}() *{{ .Name | fc }} {
	{{ .Name | lc }} := &{{ .Name | fc }}{}
	{{ .Name | lc }}.self = {{ .Name | lc }}
	return {{ .Name | lc }}
}

// -------------------------------------------

// Error return the error associate to the object
func (x *{{ $.Name | fc }}Base) Error () error {
	return x.error
}

// -------------------------------------------

// Error return the error associate to the object
func (x *{{ $.Name | fc }}Base) Validate () error {
	x.error = nil
	return x.error
}

// -------------------------------------------

{{- range .Fields }}

{{- if .Multiple }}
// {{ .Name | fc }} returns []{{ $.Name | fc}}.{{ .Name }}
func (x *{{ $.Name | fc }}Base) {{ .Name | fc }} () []{{ .Type | typeConvert }} {
	return x.{{ .Name | lc }}
}

func (x *{{ $.Name | fc }}Base) {{ .Name | fc }}IsEmpty () bool {
	return len(x.{{ .Name | lc }}) == 0
}

// {{ .Name | fc }}Size returns the size of the array {{ $.Name | fc }}.{{ .Name | lc }}
func (x *{{ $.Name | fc }}Base) {{ .Name | fc }}Size () int {
	return len(x.{{ .Name | lc }})
}

// Set{{ .Name | fc }} sets {{ $.Name | fc}}.{{ .Name }}
func (x *{{ $.Name | fc }}Base) Set{{ .Name | fc }} (v []{{ .Type | typeConvert}}) *{{ $.Name | fc }} {
	x.{{ .Name | lc }} = v
	return x.self
}

// Append{{ .Name | fc }} append {{ $.Name | fc}}.{{ .Name }}
func (x *{{ $.Name | fc }}Base) Append{{ .Name | fc }} (v ...{{ .Type | typeConvert}}) *{{ $.Name | fc }} {
	x.{{ .Name | lc }} = append(x.{{ .Name | lc }}, v...)
	return x.self
}

{{ else }}
// {{ .Name | fc }} returns {{ $.Name | fc}}.{{ .Name }}
func (x *{{ $.Name | fc }}Base) {{ .Name | fc }} () {{if .Multiple }}[]{{ end }}{{ .Type | typeConvert }} {
	return x.{{ .Name | lc }}
}

// Set{{ .Name | fc }} sets {{ $.Name | fc}}.{{ .Name }}
func (x *{{ $.Name | fc }}Base) Set{{ .Name | fc }} (v {{ .Type | typeConvert}}) *{{ $.Name | fc }} {
	x.{{ .Name | lc }} = v
	return x.self
}
{{- end }}

{{- end }}

// -------------------------------------------

// Unmarshal generate a *{{ .Name | fc }} from a valid json (bytes)
func Unmarshal (jsonBytes []byte) *{{ .Name | fc }} {
	{{ .Name | lc }} := &{{ .Name | fc }}{}
	dto := &{{ .Name | lc }}DTO{}
	err := json.Unmarshal(jsonBytes, dto)
	if err != nil {
		panic("can not unmarshal {{ .Name | fc }}DTO")
	}
	{{- range .Fields }}
	{{ $.Name | lc }}.{{ .Name | lc }} = dto.{{ .Name | fc }}
	{{- end }}
	return {{ .Name | lc }}
}

// Mashal generate a valid JSON (bytes) from *{{ $.Name | fc }}
func (x *{{ $.Name | fc }}Base) Marshal () ([]byte, error) {
	dto := {{ .Name | lc }}DTO{}
	{{- range .Fields }}
	{{- if .Exposed }}
	dto.{{ .Name | fc }} = x.{{ .Name | lc }}
	{{- end }}
	{{- end }}

	return json.Marshal(dto)
}

// -------------------------------------------

// String implements Stringer
func (x *{{ $.Name | fc }}Base) String () string {
	jsonBytes, _ := x.Marshal()
	return "*{{ $.Name | fc }}" + string(jsonBytes)
}
