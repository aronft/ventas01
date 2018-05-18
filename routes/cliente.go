package routes

import (
	"github.com/blazte/10-PracticeProject/Ventas01/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetClienteRouter ruta para el registro de cleintes
func SetClienteRouter(router *mux.Router) {
	prefix := "/api/clientes"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)

	subRouter.HandleFunc("/", controllers.BuscarClienteController).Methods("GET")

	subRouter.HandleFunc("/", controllers.CrearClienteController).Methods("POST")

	subRouter.HandleFunc("/", controllers.ActualizarClienteController).Methods("PUT")

	subRouter.HandleFunc("/", controllers.EliminarClienteController).Methods("DELETE")

	subRouter.HandleFunc("/{id}", controllers.BuscarClienteIDController).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
