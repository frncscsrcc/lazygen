// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!

package listener

import (
	"github.com/frncscsrcc/lazygen/examples/api"
	"github.com/frncscsrcc/lazygen/examples/api/role"
	"github.com/frncscsrcc/lazygen/examples/api/user"
	"log"
	"net/http"
)

func GetRouter() *api.Router {
	router := api.NewRouter()

	// ------------------------------------
	// User
	// ------------------------------------
	router.AddRoute(
		api.NewRouteConfig(
			"/api/v1/user/:id",
			user.ReadUser,
			http.MethodGet,
			"User",
			"Guest",
		),
	)
	router.AddRoute(
		api.NewRouteConfig(
			"/api/v1/user/:id",
			user.UpdateUser,
			http.MethodPut,
			"User",
		),
	)
	router.AddRoute(
		api.NewRouteConfig(
			"/api/v1/user/:id",
			user.DeleteUser,
			http.MethodDelete,
			"Admin",
		),
	)

	// ------------------------------------
	// Role
	// ------------------------------------
	router.AddRoute(
		api.NewRouteConfig(
			"/api/v1/role",
			role.CreateRole,
			http.MethodPost,
			"Admin",
		),
	)
	router.AddRoute(
		api.NewRouteConfig(
			"/api/v1/role/:id",
			role.ReadRole,
			http.MethodGet,
			"User",
			"Guest",
		),
	)
	router.AddRoute(
		api.NewRouteConfig(
			"/api/v1/role/:id",
			role.UpdateRole,
			http.MethodPut,
			"User",
		),
	)
	router.AddRoute(
		api.NewRouteConfig(
			"/api/v1/role/:id",
			role.DeleteRole,
			http.MethodDelete,
			"Admin",
		),
	)

	// Load custom endpoints
	RegisterCustomRouters(router)

	return router
}

func StartListener(host string, port string) {
	mainRouter := GetRouter()
	log.Printf("Listening on %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, mainRouter))
}
