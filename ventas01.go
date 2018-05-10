package main

import (
	"github.com/blazte/Ventas01/controllers"
	"github.com/blazte/Ventas01/models"
)

func main() {
	var c models.Cliente

	c = models.Cliente{
		Dni:       "1231",
		Nombres:   "Aron",
		Apellidos: "Flores",
	}

	controllers.CrearClienteController(c)
}
