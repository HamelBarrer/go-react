package routers

import (
	"net/http"

	"github.com/HamelBarrer/go-react/bd"
	"github.com/HamelBarrer/go-react/models"
)

// BajaRelacion realiza el borrado de la relacion entre usuarios
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar una relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
