package routes

import (
	"goblog/app/http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterWebRoutes(r *mux.Router) {
	pc := new(controllers.PagesController)
	ac := new(controllers.ArticlesController)
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
}
