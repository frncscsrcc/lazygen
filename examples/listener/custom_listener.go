package listener

import (
	"github.com/frncscsrcc/lazygen/examples/api"
)

func RegisterCustomRouters(router *api.Router) *api.Router {
	// Register here your custom APIs
	/* Eg:
	    router.AddRoute(
			api.NewRouteConfig(
				"/custom_end_point/:id",
				user.ReadUser,
				http.MethodPost,
				"Admin",
			),
		)
	*/

	return router
}
