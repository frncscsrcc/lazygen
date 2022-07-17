// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!

// Router
package api

import (
	"log"
	"fmt"
	"net/http"
	"strings"
)

type routingData struct {
	isAuth bool
	token  string
	roles  []string
}

type routableFactory func() Routable

type Routable interface {
	SetData(*routingData)
	GetData() *routingData
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type RouteConfig struct {
	Path    string
	Factory routableFactory
	Roles   []string
	Method  string
}

func NewRouteConfig(path string, factory routableFactory, method string, roles ...string) RouteConfig {
	rolesSlice := make([]string, 0)
	rolesSlice = append(rolesSlice, roles...)
	return RouteConfig{
		Path:    path,
		Factory: factory,
		Method:  method,
		Roles:   rolesSlice,
	}
}

type Router struct {
	routes map[string]RouteConfig
	data   *routingData
}

func NewRouter() *Router {
	r := &Router{
		routes: make(map[string]RouteConfig),
	}
	return r
}

// ---------------------------------------

func (router *Router) SetData(data *routingData) {
	router.data = data
}

func (router *Router) GetData() *routingData {
	return router.data
}

func (router *Router) AddRoute(config RouteConfig) *Router {
	log.Printf("Loaded API Endpoint: %s (Method=%s Roles=%+v)\n", config.Path, config.Method, config.Roles)
	router.routes[config.Path] = config
	return router
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.authenticate(r)

	url := r.URL.Path

	if routeConfig, defined := router.routes[url]; defined {

		if len(routeConfig.Roles) > 0 {
			if !router.data.isAuth {
				router.Forbidden(w)
				return
			}

			hasRole := false
			for _, userRole := range router.data.roles {
				for _, requestedRoles := range routeConfig.Roles {
					if strings.EqualFold(userRole, requestedRoles) {
						hasRole = true
						break
					}
					if hasRole {
						break
					}
				}
			}

			if !hasRole {
				router.Forbidden(w)
				return
			}
		}

		next := routeConfig.Factory()
		next.SetData(router.data)
		next.ServeHTTP(w, r)
		return
	}

	router.NotFound(w)
}

// ---------------------------------------

func (router *Router) authenticate(r *http.Request) {
	data := &routingData{
		roles: make([]string, 0),
	}

	// Check for an Authorization token
	if token := r.Header.Get("Authorization"); token != "" {
		data.token = token
		data.isAuth = true
	}

	// Check for a cookie
	if cookie, err := r.Cookie("sessionID"); err == nil {
		_ = cookie.Value
		data.isAuth = true
	}

	data.roles = append(data.roles, "ADMIN") // TODO EXTRACT ROLES

	router.data = data
}

func (router *Router) Forbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprint(w, "FORBIDDEN")
}

func (router *Router) NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "NOT FOUND")
}
