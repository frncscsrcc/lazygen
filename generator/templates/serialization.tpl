// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `{{ .Name | fc }}`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `{{ .Name | lc }}`

// {{ .Name | fc }}: {{ .Description }}
package {{ .Name | lc }}

import (
	"encoding/json"
	{{- range .Requires }}
	"{{ . }}"
	{{- end }}
)

// DTO to marshal in JSON
type DTO struct {
	{{- range .Fields }}
    {{ .Name | fc }} {{if .Multiple }}[]{{ end }}{{ .Type | getDTO }} `json:"{{ .Name | lc }}"`
    {{- end }}
}

// New{{ .Name | fc }} build a DTO empty structure
func NewDTO() *DTO {
	dto := &DTO{}
	{{- range .Fields }}
	{{- if .Multiple}}
	{{- if .Custom }}
	dto.{{ .Name | fc }} = make([]{{ .Type | package }}.DTO, 0)
	{{- else }}
	dto.{{ .Name | fc }} = make([]{{ .Type | typeConvert }}, 0)
	{{- end }}
	{{- end }}
	{{- end }}
	return dto
}

// -------------------------------------------

func (x *{{ $.Name | fc }}Base) ToDTO() *DTO {
	dto := NewDTO()
	{{- range .Fields }}
	{{- if .Exposed }}
	{{- if .Custom }}
	{{- if .Multiple}}
	dto.{{ .Name | fc }} = make([]*{{ .Type | lc }}.DTO, 0)
	for _, {{ .Type | lc }} := range(x.{{ .Name }}) {
		if {{ .Type | lc }} != nil {
			dto.{{ .Name | fc }} = append(dto.{{ .Name | fc }}, {{ .Type | lc }}.ToDTO())
		}
	}
	{{- else }}
	if x.{{ .Name | lc }} != nil {
		dto.{{ .Name | fc }} = x.{{ .Name | lc }}.ToDTO()
	}
	{{- end }}
	{{- else }}
	dto.{{ .Name | fc }} = x.{{ .Name | fc }}()
	{{- end }}
	{{- end }}
	{{- end }}
	return dto
}

func (dto *DTO) ToType() *{{ $.Name | fc }} {
    {{ $.Name | lc }} := New{{ $.Name | fc }}()
	{{- range .Fields }}
	{{- if .Exposed }}
	{{- if .Custom }}
	{{- if .Multiple }}
	for _, {{ .Type | lc }} := range(dto.{{ .Name | fc }} ){
		if {{ .Type | lc }} != nil {
			{{ $.Name | lc }}.Append{{ .Name | fc }}({{ .Type | lc }}.ToType())
		}
	}
	{{- else }}
	if dto.{{ .Name | fc }} != nil {
		{{ $.Name | lc }}.Set{{ .Name | fc }}(dto.{{ .Name | fc }}.ToType())
	}
	{{- end }}
	{{- else }}
	{{ $.Name | lc }}.Set{{ .Name | fc }}(dto.{{ .Name | fc }})
	{{- end }}
	{{- end }}
	{{- end }}
	return {{ $.Name | lc }}
}


// -------------------------------------------

// FromJSON generate a *{{ .Name | fc }} from a valid json (passed as string)

func FromJSON (jsonString string) (*{{ .Name | fc }}, error) {
	dto := &DTO{}
	err := json.Unmarshal([]byte(jsonString), dto)
	if err != nil {
        return New{{ .Name | fc }}(), err
	}
	{{ .Name | lc }} := dto.ToType()
	return {{ .Name | lc }}, nil
}


// ToJSON generate a valid JSON (as string) from *{{ $.Name | fc }}
func (x *{{ $.Name | fc }}Base) ToJSON () (string, error) {
	dto := x.ToDTO()

	bytes, err := json.Marshal(dto)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

