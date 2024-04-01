package route

import (
	"goblog/routes"
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func Initizalize() {
	Router = mux.NewRouter()
	routes.RegisterWebRoutes(Router)
}
func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		return ""
	}

	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
