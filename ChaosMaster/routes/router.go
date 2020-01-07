package routes

import (
	"net/http"

	"mallekoppie/ChaosGenerator/ChaosMaster/middleware"
	"mallekoppie/ChaosGenerator/ChaosMaster/service"

	"github.com/gorilla/mux"
)

const (
	AllowedContentType string = "application/json"
)

type Route struct {
	Path               string
	Method             string
	HandlerFunc        http.HandlerFunc
	SlaMs              int64
	RolesRequired      []string
	AllowedContentType string
}

type Routes []Route

// Add longest paths first
var serviceRoutes = Routes{
	Route{
		Path:          "/testgroup",
		Method:        http.MethodGet,
		HandlerFunc:   service.GetAllTestGroups,
		SlaMs:         100,
		RolesRequired: []string{"user"},
	},
	Route{
		Path:               "/testgroup",
		Method:             http.MethodPost,
		HandlerFunc:        service.AddTestGroup,
		SlaMs:              100,
		RolesRequired:      []string{"user"},
		AllowedContentType: AllowedContentType,
	},
	Route{
		Path:          "/testgroup",
		Method:        http.MethodDelete,
		HandlerFunc:   service.DeleteTestGroup,
		SlaMs:         100,
		RolesRequired: []string{"user"},
	},
	Route{
		Path:          "/agents",
		Method:        http.MethodGet,
		HandlerFunc:   service.GetAllAgents,
		SlaMs:         100,
		RolesRequired: []string{"user"},
	},
	Route{
		Path:               "/agents",
		Method:             http.MethodPut,
		HandlerFunc:        service.UpdateAgent,
		SlaMs:              100,
		RolesRequired:      []string{"user"},
		AllowedContentType: AllowedContentType,
	},
	Route{
		Path:               "/agents",
		Method:             http.MethodDelete,
		HandlerFunc:        service.DeleteAgent,
		SlaMs:              100,
		RolesRequired:      []string{"user"},
		AllowedContentType: AllowedContentType,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for index := range serviceRoutes {
		route := serviceRoutes[index]
		var handler http.Handler
		handler = route.HandlerFunc

		// Add the middleware components. The are executed from the bottom up
		handler = middleware.AllowedContentType(handler, route.AllowedContentType)
		handler = middleware.AllowCors(handler)
		//handler = middleware.UseOAuth2(handler, route.RolesRequired) // Disabled during development
		handler = middleware.TrackServiceMethodSla(handler, route.SlaMs)

		router.
			Path(route.Path).
			Methods(route.Method).
			Handler(handler)

	}

	return router
}