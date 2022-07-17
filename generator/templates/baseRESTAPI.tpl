// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `{{ .Name | fc }}`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `{{ .Name | lc }}`

// {{ .Name | fc }}: Base controllers for the {{ .Name | fc }} entity
package {{ .Name | lc }}

import (
	"net/http"

	"{{ .RouterPackage }}"
)

{{- if .API.Create.Enabled }}
type Create{{ .Name | fc }}RouterBase struct { api.Router }

type Create{{ .Name | fc }}Router struct { Create{{ .Name | fc }}RouterBase }

func Create{{ .Name | fc }}() api.Routable { return &Create{{ .Name | fc }}Router{} }

func (router *Create{{ .Name | fc }}RouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}
{{- end }}

{{- if .API.Read.Enabled }}
// -------------------------------

type Read{{ .Name | fc }}RouterBase struct { api.Router }

type Read{{ .Name | fc }}Router struct { Read{{ .Name | fc }}RouterBase }

func Read{{ .Name | fc }}() api.Routable { return &Read{{ .Name | fc }}Router{} }

func (router *Read{{ .Name | fc }}RouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}
{{- end }}

{{- if .API.Update.Enabled }}
// -------------------------------

type Update{{ .Name | fc }}RouterBase struct { api.Router }

type Update{{ .Name | fc }}Router struct { Update{{ .Name | fc }}RouterBase }

func Update{{ .Name | fc }}() api.Routable { return &Update{{ .Name | fc }}Router{} }

func (router *Update{{ .Name | fc }}RouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}
{{- end }}

{{- if .API.Delete.Enabled }}
// -------------------------------

type Delete{{ .Name | fc }}RouterBase struct { api.Router }

type Delete{{ .Name | fc }}Router struct { Delete{{ .Name | fc }}RouterBase }

func Delete{{ .Name | fc }}() api.Routable { return &Delete{{ .Name | fc }}Router{} }

func (router *Delete{{ .Name | fc }}RouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}
{{- end }}