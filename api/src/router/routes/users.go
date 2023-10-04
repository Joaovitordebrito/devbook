package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:      "/user",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		Auth:     false,
	},
	{
		URI:      "/user",
		Method:   http.MethodGet,
		Function: controllers.GetUsers,
		Auth:     false,
	},
	{
		URI:      "/user/{userID}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		Auth:     false,
	},
	{
		URI:      "/user/{userID}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		Auth:     false,
	},
	{
		URI:      "/user/{userID}",
		Method:   http.MethodDelete,
		Function: controllers.DeletUser,
		Auth:     false,
	},
}
