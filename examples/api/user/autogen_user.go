// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `User`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `user`

// User: Base controllers for the User entity
package user

import (
	"net/http"

	"github.com/frncscsrcc/lazygen/examples/api"
)

// -------------------------------

type ReadUserRouterBase struct{ api.Router }

type ReadUserRouter struct{ ReadUserRouterBase }

func ReadUser() api.Routable { return &ReadUserRouter{} }

func (router *ReadUserRouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}

// -------------------------------

type UpdateUserRouterBase struct{ api.Router }

type UpdateUserRouter struct{ UpdateUserRouterBase }

func UpdateUser() api.Routable { return &UpdateUserRouter{} }

func (router *UpdateUserRouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}

// -------------------------------

type DeleteUserRouterBase struct{ api.Router }

type DeleteUserRouter struct{ DeleteUserRouterBase }

func DeleteUser() api.Routable { return &DeleteUserRouter{} }

func (router *DeleteUserRouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.NotFound(w)
}
