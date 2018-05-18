package controllers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/blazte/10-PracticeProject/Ventas01/models"
	"github.com/blazte/10-PracticeProject/ventas01/commons"
)

//CrearClienteController es el controlador de clientes
func CrearClienteController(w http.ResponseWriter, r *http.Request) {
	cliente := models.Cliente{}
	m := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer cliente a registrar: %s\n", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	id, err := cliente.CrearCliente()
	if err != nil {
		m.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	cliente.ID = id

	uri := fmt.Sprintf("%s%s%s", r.URL.Scheme, r.Host, r.URL.Path)

	self := models.Navigation{
		Title:       "Self",
		Description: "self source",
		Link:        fmt.Sprintf("%s%d", uri, id),
	}
	prev := models.Navigation{
		Title:       "prev",
		Description: "Previous source",
		Link:        fmt.Sprintf("%s%d", uri, id-1),
	}
	next := models.Navigation{
		Title:       "next",
		Description: "Next source",
		Link:        fmt.Sprintf("%s%d", uri, id+1),
	}
	cliente.Navigations = []models.Navigation{self, prev, next}

	d := r.Header.Get("Accept")
	switch d {

	case "text/plain":

	case "application/json":
		j, err := json.Marshal(cliente)
		if err != nil {
			m.Message = fmt.Sprintf("Error al convertir los clientes a  json: %s\n", err)
			m.Code = http.StatusInternalServerError
			commons.DisplayMessage(w, m)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return

	case "text/xml":
		x, err := xml.MarshalIndent(cliente, "", "  ")
		if err != nil {
			m.Message = fmt.Sprintf("Error al convertir los clientes a  xml: %s\n", err)
			m.Code = http.StatusInternalServerError
			commons.DisplayMessage(w, m)
			return
		}
		w.Header().Set("Content-type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write(x)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Debe especificar un content-type válido"))
		return
	}

	m.Message = "Cliente creado con exito"
	m.Code = http.StatusCreated
	commons.DisplayMessage(w, m)
}

//ActualizarClienteController controlador para actualizar clientes
func ActualizarClienteController(w http.ResponseWriter, r *http.Request) {
	cliente := models.Cliente{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el cliente a actualizar: %s\n", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	err = cliente.ActualizarCliente()
	if err != nil {
		m.Message = fmt.Sprintf("Error al actualizar el cliente: %s\n", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "Cliente Actualizado con exito"
	m.Code = http.StatusOK
	commons.DisplayMessage(w, m)

}

//EliminarClienteController controlador par aeliminar un cliente
func EliminarClienteController(w http.ResponseWriter, r *http.Request) {
	cliente := models.Cliente{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		m.Message = fmt.Sprintf("Error al lee el cliente a eliminar: %s\n", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	err = cliente.EliminarCliente()
	if err != nil {
		m.Message = fmt.Sprintf("Error al eliminar el cliente: %s\n", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "Cliente eliminado con exito"
	m.Code = http.StatusOK
	commons.DisplayMessage(w, m)
}

//BuscarClienteController es el controlador de cliente
func BuscarClienteController(w http.ResponseWriter, r *http.Request) {
	clientes := []models.Cliente{}
	m := models.Message{}

	clientes, err := models.BuscarCliente()
	if err != nil {
		m.Message = fmt.Sprintf("Error al seleccionar los cliente: %s\n", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	uri := fmt.Sprintf("%s%s%s", r.URL.Scheme, r.Host, r.URL.Path)

	for i, cliente := range clientes {
		self := models.Navigation{
			Title:       "Self",
			Description: "Description self",
			Link:        fmt.Sprintf("%s%d", uri, cliente.ID),
		}
		prev := models.Navigation{
			Title:       "Prev",
			Description: "Description prev",
			Link:        fmt.Sprintf("%s%d", uri, cliente.ID-1),
		}
		next := models.Navigation{
			Title:       "Next",
			Description: "Description next",
			Link:        fmt.Sprintf("%s%d", uri, cliente.ID+1),
		}

		cliente.Navigations = []models.Navigation{self, prev, next}
		clientes[i].Navigations = cliente.Navigations
	}

	d := r.Header.Get("Accept")
	switch d {
	case "text/plain":
		r := models.String(clientes)
		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(r))
		return
	case "application/json":
		j, err := json.Marshal(clientes)
		if err != nil {
			m.Message = fmt.Sprintf("Error al convertir los clientes a  json: %s\n", err)
			m.Code = http.StatusInternalServerError
			commons.DisplayMessage(w, m)
			return
		}
		if len(clientes) > 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(j)
			return
		} else {
			m.Code = http.StatusNoContent
			m.Message = "No se encontraron comentarios"
			commons.DisplayMessage(w, m)
		}
	case "text/xml":
		x, err := xml.MarshalIndent(clientes, "", "  ")
		if err != nil {
			m.Message = fmt.Sprintf("Error al convertir los clientes a  xml: %s\n", err)
			m.Code = http.StatusInternalServerError
			commons.DisplayMessage(w, m)
			return
		}
		w.Header().Set("Content-type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write(x)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Debe especificar un content-type válido"))
		return
	}

}

//BuscarClienteIDController busca un cliente en especifico
func BuscarClienteIDController(w http.ResponseWriter, r *http.Request) {
	m := models.Message{}
	// vars := mux.Vars(r)
	// id1, _ := strconv.Atoi(vars["id"])
	ids := strings.TrimPrefix(r.URL.Path, "/api/clientes/")
	id, _ := strconv.Atoi(ids)

	c, err := models.BuscarClienteID(id)
	if err != nil {
		m.Message = fmt.Sprintf("Error al ejecutar la consulta, %s\n", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
	}

	uri := fmt.Sprintf("%s%s%s", r.URL.Scheme, r.Host, r.URL.Path)
	fmt.Println(r.Host)
	result := strings.SplitAfter(uri, "/")
	uri = fmt.Sprintf("%s%s%s", result[0], result[1], result[2])

	self := models.Navigation{
		Title:       "Self",
		Description: "Description Self",
		Link:        fmt.Sprintf("%s%d", uri, c.ID),
	}

	prev := models.Navigation{
		Title:       "Prev",
		Description: "Description Prev",
		Link:        fmt.Sprintf("%s%d", uri, c.ID-1),
	}

	next := models.Navigation{
		Title:       "Next",
		Description: "Description Next",
		Link:        fmt.Sprintf("%s%d", uri, c.ID+1),
	}

	c.Navigations = []models.Navigation{self, prev, next}

	d := r.Header.Get("Accept")
	switch d {
	case "application/json":
		j, err := json.Marshal(c)
		if err != nil {
			m.Message = fmt.Sprintf("error al convertir el cliente a json:%s\n", err)
			m.Code = http.StatusInternalServerError
			commons.DisplayMessage(w, m)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Debe especificar un content-type válido"))
		return
	}
}
