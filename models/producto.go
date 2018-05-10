package models

//Producto es la estructura de productos de la BD
type Producto struct {
	ID          int
	Codigo      string
	Nombre      string
	Precio      float64
	ProveedorID int
}
