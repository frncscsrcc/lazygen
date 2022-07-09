// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `{{ .Name | fc }}`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `{{ .Name | lc }}`

// {{ .Name | fc }}: {{ .Description }}
package {{ .Name | lc }}

import (
	"encoding/json"
    "errors"
	{{ range .Requires }}"{{ . }}"{{ end }}
)

// DTO to marshal in JSON
type DTO struct {
	{{- range .Fields }}
    {{ .Name | fc }} {{if .Multiple }}[]{{ end }}{{ .Type | getDTO }} `json:"{{ .Name | lc }}"`
    {{- end }}
}

// -------------------------------------------

func (x *{{ $.Name | fc }}Base) ToDTO() *DTO {
	dto := &DTO{}
	{{- range .Fields }}
	{{- if .Exposed }}
	dto.{{ .Name | fc }} = x.{{ .Name | lc }}{{- if .Custom }}.ToDTO(){{- end }}
	{{- end }}
	{{- end }}
	return dto
}

func (dto *DTO) ToType() *{{ $.Name | fc }} {
    {{ $.Name | lc }} := New{{ $.Name | fc }}()
	{{- range .Fields }}
	{{- if .Exposed }}
	{{ $.Name | lc }}.{{ .Name | lc }} = dto.{{ .Name | fc }}{{- if .Custom }}.ToType(){{- end }}
	{{- end }}
	{{- end }}
	return {{ $.Name | lc }}
}


// -------------------------------------------

// FromJSON generate a *{{ .Name | fc }} from a valid json (bytes)

func FromJSON (jsonBytes []byte) (*{{ .Name | fc }}, error) {
	{{ .Name | lc }} := &{{ .Name | fc }}{}
	dto := &DTO{}
	err := json.Unmarshal(jsonBytes, dto)
	if err != nil {
        return New{{ .Name | fc }}(), errors.New("can not unmarshal {{ .Name | fc }}DTO")
	}
	{{- range .Fields }}
	{{ $.Name | lc }}.{{ .Name | lc }} = dto.{{ .Name | fc }}{{- if .Custom }}.ToType(){{- end }}
	{{- end }}
	return {{ .Name | lc }}, nil
}


// ToJSON generate a valid JSON (bytes) from *{{ $.Name | fc }}
func (x *{{ $.Name | fc }}Base) ToJSON () (string, error) {
	dto := DTO{}
	{{- range .Fields }}
	{{- if .Exposed }}
	dto.{{ .Name | fc }} = x.{{ .Name | lc }}{{ if .Custom }}.ToDTO(){{ end }}
	{{- end }}
	{{- end }}

	bytes, err := json.Marshal(dto)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

