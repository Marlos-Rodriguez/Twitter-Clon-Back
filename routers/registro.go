package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//Registro funcion para crear nuevo usuario en DB
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email es obligatorio", 400)
		return
	}
	if len(t.Email) < 6 {
		http.Error(w, "Contraseña debe tener 6 o mas caracteres", 400)
		return
	}

	_, encontrado, _ := db.CheckExistingUser(t.Email)

	if encontrado {
		http.Error(w, "Usuario ya existente con ese email", 400)
		return
	}

	_, status, err := db.InsertRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un Error al realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se inserto el nuevo usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
