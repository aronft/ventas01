package models

import (
	"fmt"
	"time"

	"github.com/blazte/Ventas01/configuration"
)

//Cliente es la estructura de cliente
type Cliente struct {
	ID        int    `json:"id"`
	Dni       string `json:"dni"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	CreadAt   time.Time
	UpdateAt  time.Time
}

//CrearCliente es el metodo para crear clientes
func (c *Cliente) CrearCliente() {
	q := `INSERT INTO clientes (dni, nombres, apellidos) values($1, $2, $3);`
	db := configuration.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Printf("Error al preparar la consulta: %s", err)
	}

	_, err = stmt.Exec(c.Dni, c.Nombres, c.Apellidos)
	if err != nil {
		fmt.Printf("Error al ejecutar la consulta: %s", err)
	}
}
