package routes

import (
	"github.com/gorilla/mux"
)

//InitRoutes inicia las rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetClienteRouter(router)

	return router
}
