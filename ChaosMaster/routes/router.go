package routes

import (
	"net/http"

	"mallekoppie/ChaosGenerator/ChaosMaster/middleware"
	"mallekoppie/ChaosGenerator/ChaosMaster/service"

	"github.com/gorilla/mux"
)

const (
	allowedContentType       string = "application/json"
	allowedContentTypeDelete string = "application/json;charset=UTF-8"
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
		Path:          "/testgroups",
		Method:        http.MethodGet,
		HandlerFunc:   service.GetAllTestGroups,
		SlaMs:         100,
		RolesRequired: []string{"user"},
	},
	Route{
		Path:               "/testgroups",
		Method:             http.MethodPost,
		HandlerFunc:        service.AddTestGroup,
		SlaMs:              100,
		RolesRequired:      []string{"user"},
		AllowedContentType: allowedContentType,
	},
	Route{
		Path:               "/testgroups",
		Method:             http.MethodPut,
		HandlerFunc:        service.UpdateTestGroup,
		SlaMs:              100,
		RolesRequired:      []string{"user"},
		AllowedContentType: allowedContentType,
	},
	Route{
		Path:          "/testgroups/{id}",
		Method:        http.MethodDelete,
		HandlerFunc:   service.DeleteTestGroup,
		SlaMs:         100,
		RolesRequired: []string{"user"},
	},
	Route{
		Path:               "/testcollections",
		Method:             http.MethodPost,
		HandlerFunc:        service.AddTestCollection,
		SlaMs:              100,
		RolesRequired:      []string{"user"},
		AllowedContentType: allowedContentType,
	},
	Route{
		Path:               "/testcollections",
		Method:             http.MethodPut,
		HandlerFunc:        service.UpdateTestCollection,
		SlaMs:              100,
		RolesRequired:      []string{"user"},
		AllowedContentType: allowedContentType,
	},
	// Only during development
	Route{
		Path:          "/testcollections",
		Method:        http.MethodOptions,
		HandlerFunc:   service.TemplateToCopy,
		SlaMs:         100,
		RolesRequired: []string{"user"},
	},
	Route{
		Path:          "/testcollections/{id}",
		Method:        http.MethodDelete,
		HandlerFunc:   service.DeleteTestCollection,
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
		AllowedContentType: allowedContentType,
	},
	Route{
		Path:               "/agents",
		Method:             http.MethodDelete,
		HandlerFunc:        service.DeleteAgent,
		SlaMs:              100,
		RolesRequired:      []string{"user"},
		AllowedContentType: allowedContentTypeDelete,
	},
	// Only during development
	Route{
		Path:          "/agents",
		Method:        http.MethodOptions,
		HandlerFunc:   service.TemplateToCopy,
		SlaMs:         100,
		RolesRequired: []string{"user"},
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
