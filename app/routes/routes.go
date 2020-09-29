package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"hsaaku.com/auth-api/app/routes/handler"
)

// RequestHandlerFunction has db as an additional argument
type RequestHandlerFunction func(db *sqlx.DB, w http.ResponseWriter, r *http.Request)

// Route holds handler function parameters
type Route struct {
	method  string
	path    string
	handler RequestHandlerFunction
}

// InitializeRoutes will register all defined routes
func InitializeRoutes(db *sqlx.DB, router *mux.Router) {
	routes := []Route{
		Route{"GET", "/", handler.Hello},
		Route{"GET", "/users", handler.GetUsers},
	}

	for _, route := range routes {
		router.HandleFunc(route.path, func(w http.ResponseWriter, r *http.Request) {
			route.handler(db, w, r)
		}).Methods(route.method)
	}
}
