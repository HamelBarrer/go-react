package routers

import (
	"net/http"

	"github.com/HamelBarrer/go-react/bd"
	"github.com/HamelBarrer/go-react/models"
)

// AltaRelacion realiza el registro de la relacion entre usuarios
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al inserta relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
