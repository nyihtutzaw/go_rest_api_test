package routes

import (
	"github.com/gorilla/mux"
)

//NewRouter is main routing entry point
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	apiRoutes := router.PathPrefix("/api").Subrouter()

	userHandler(apiRoutes)
	authorHandler(apiRoutes)
	bookHandler(apiRoutes)
	return router
}
