// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!

package listener

import (
    "log"
    "net/http"
	"github.com/frncscsrcc/lazygen/examples/api"

    {{- range .Entities }}
    {{- if or .API.Create.Enabled .API.Read.Enabled .API.Update.Enabled .API.Delete.Enabled }}
    "{{ $.Module }}/{{ $.FolderAPI }}/{{ .Name | lc }}"
    {{- end }}
    {{- end }}
)

func GetRouter() *api.Router {
	router := api.NewRouter()

	{{- range .Entities }}
    {{- if or .API.Create.Enabled .API.Read.Enabled .API.Update.Enabled .API.Delete.Enabled }}
    
    // ------------------------------------
    // {{ .Name | fc }}
    // ------------------------------------
    
    {{- if .API.Create.Enabled }}
    router.AddRoute(
        api.NewRouteConfig(
            "{{ $.BasePathAPI }}/{{ .Name | lc }}",
            {{ .Name | lc }}.Create{{ .Name | fc }},
            http.MethodPost,
            {{- range .API.Create.Roles }}
            "{{ . }}",
            {{- end }}
        ),
    )
    {{- end }}
    {{- if .API.Read.Enabled }}
    router.AddRoute(
        api.NewRouteConfig(
            "{{ $.BasePathAPI }}/{{ .Name | lc }}/:id",
            {{ .Name | lc }}.Read{{ .Name | fc }},
            http.MethodGet,
            {{- range .API.Read.Roles }}
            "{{ . }}",
            {{- end }}
        ),
    )
    {{- end }}
    {{- if .API.Update.Enabled }}
    router.AddRoute(
        api.NewRouteConfig(
            "{{ $.BasePathAPI }}/{{ .Name | lc }}/:id",
            {{ .Name | lc }}.Update{{ .Name | fc }},
            http.MethodPut,
            {{- range .API.Update.Roles }}
            "{{ . }}",
            {{- end }}
        ),
    )
    {{- end }}
    {{- if .API.Delete.Enabled }}
    router.AddRoute(
        api.NewRouteConfig(
            "{{ $.BasePathAPI }}/{{ .Name | lc }}/:id",
            {{ .Name | lc }}.Delete{{ .Name | fc }},
            http.MethodDelete,
            {{- range .API.Delete.Roles }}
            "{{ . }}",
            {{- end }}
        ),
    )
    {{- end }}

	{{- end }}
	{{- end }}

    // Load custom endpoints
    RegisterCustomRouters(router)

	return router
}

func StartListener(host string, port string) {
    mainRouter := GetRouter()
    log.Printf("Listening on %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, mainRouter))
}