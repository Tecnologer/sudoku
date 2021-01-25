package router

import (
	"github.com/gorilla/mux"
	"github.com/tecnologer/sudoku/clients/api/middleware"
)

//Router provides a new router instance
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("api/game", middleware.NewGame).Methods("POST", "OPTIONS")

	return router
}
