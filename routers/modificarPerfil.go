package routers

import (
	"encoding/json"
	"net/http"

	"github.com/HamelBarrer/go-react/bd"
	"github.com/HamelBarrer/go-react/models"
)

// ModificarPerfil modifica el perfil de usuario
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro de usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
