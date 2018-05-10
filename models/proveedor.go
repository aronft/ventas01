package models

import (
	"log"

	"github.com/blazte/Ventas01/configuration"
)

//Proveedor es la estructura de proveedores de la BD
type Proveedor struct {
	ID              int
	Nombre          string
	Nif             string
	DireccionCalle  string
	DireccionNumero string
}

//CrearProveedor Crea un registro de un proveedor
func (c Proveedor) CrearProveedor() {
	q := `INSERT INTO 
			proveedores(nombres, ruc, direccion_numero, direccion_calle, telefono, )
			VALUES ($1, $2, $3)`
	db := configuration.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}
