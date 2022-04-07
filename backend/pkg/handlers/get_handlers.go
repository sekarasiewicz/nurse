package handlers

import "net/http"

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

func GetRoutes() []Route {
	return []Route{
		{"/", apiBuilder(HomeHandler)},
	}
}
