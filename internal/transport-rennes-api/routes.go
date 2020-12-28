package api

import "github.com/gorilla/mux"

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// endpoints for apartment entity
	router.HandleFunc("/upcomingbus", getUpcomingBus).Methods("GET")

	return router
}
