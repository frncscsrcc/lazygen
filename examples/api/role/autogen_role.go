// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `Role`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `role`

// Role: Base controllers for the Role entity
package role

import (
	"net/http"

	"github.com/frncscsrcc/lazygen/examples/api"
)

type CreateRoleRouterBase struct{ api.Router }

type CreateRoleRouter struct{ CreateRoleRouterBase }

func CreateRole() api.Routable { return &CreateRoleRouter{} }

func (router *CreateRoleRouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}

// -------------------------------

type ReadRoleRouterBase struct{ api.Router }

type ReadRoleRouter struct{ ReadRoleRouterBase }

func ReadRole() api.Routable { return &ReadRoleRouter{} }

func (router *ReadRoleRouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}

// -------------------------------

type UpdateRoleRouterBase struct{ api.Router }

type UpdateRoleRouter struct{ UpdateRoleRouterBase }

func UpdateRole() api.Routable { return &UpdateRoleRouter{} }

func (router *UpdateRoleRouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}

// -------------------------------

type DeleteRoleRouterBase struct{ api.Router }

type DeleteRoleRouter struct{ DeleteRoleRouterBase }

func DeleteRole() api.Routable { return &DeleteRoleRouter{} }

func (router *DeleteRoleRouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}
