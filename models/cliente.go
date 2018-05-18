package models

import (
	"fmt"
	"time"

	"github.com/blazte/10-PracticeProject/Ventas01/configuration"
)

//Cliente es la estructura de cliente
type Cliente struct {
	ID          int    `json:"id"`
	Dni         string `json:"dni"`
	Nombres     string `json:"nombres"`
	Apellidos   string `json:"apellidos"`
	CreadAt     time.Time
	UpdateAt    time.Time
	Navigations []Navigation `json:"navigation"`
}

//CrearCliente es el metodo para crear clientes
func (c *Cliente) CrearCliente() (int, error) {
	var id int
	q := `INSERT INTO clientes (dni, nombres, apellidos) values($1, $2, $3) RETURNING cliente_id;`

	db := configuration.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return 0, err
	}

	err = stmt.QueryRow(c.Dni, c.Nombres, c.Apellidos).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//ActualizarCliente actualiza un cliente
func (c Cliente) ActualizarCliente() error {
	q := `UPDATE clientes SET dni = $1, nombres = $2, apellidos = $3, updated_at = $4 WHERE cliente_id = $5;`

	db := configuration.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	c.UpdateAt = time.Now()
	_, err = stmt.Exec(c.Dni, c.Nombres, c.Apellidos, c.UpdateAt, c.ID)
	if err != nil {
		return err
	}
	return nil
}

//EliminarCliente elimina un registro de cliente
func (c Cliente) EliminarCliente() error {

	q := `DELETE FROM clientes WHERE cliente_id = $1`

	db := configuration.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(c.ID)
	if err != nil {
		return err
	}
	return nil
}

//BuscarCliente busca todos los clientes
func BuscarCliente() ([]Cliente, error) {
	c := []Cliente{}
	cliente := Cliente{}
	q := `SELECT cliente_id, dni, nombres, apellidos FROM clientes;`

	db := configuration.GetConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return c, err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		err = rows.Scan(&cliente.ID, &cliente.Dni, &cliente.Nombres, &cliente.Apellidos)
		c = append(c, cliente)
		if err != nil {
			return c, err
		}
		i++
	}
	return c, nil
}

//String convierte el cliente a string
func String(c []Cliente) string {
	var r string
	for _, cliente := range c {
		r += fmt.Sprintf("%d, %s, %s, %s\n", cliente.ID, cliente.Dni, cliente.Nombres, cliente.Apellidos)
	}
	return r
}

//BuscarClienteID obtiene un clietne en especifico
func BuscarClienteID(id int) (c Cliente, err error) {

	q := `SELECT cliente_id, dni, nombres, apellidos FROM clientes WHERE cliente_id = $1 `

	db := configuration.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}

	err = stmt.QueryRow(id).Scan(&c.ID, &c.Dni, &c.Nombres, &c.Apellidos)
	if err != nil {
		return
	}
	return c, nil
}
