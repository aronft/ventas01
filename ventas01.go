package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/blazte/10-PracticeProject/Ventas01/routes"
	"github.com/urfave/negroni"
)

func main() {
	//Inicia las rutas
	router := routes.InitRoutes()

	//Inicia los middlewares

	n := negroni.Classic()
	n.UseHandler(router)

	//Inicia el servidor

	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}

	log.Println("Iniciado el servidor en http://localhost:8080")
	fmt.Println(server.ListenAndServe())
	log.Println("finalizó la ejecucuin del programa")
}
