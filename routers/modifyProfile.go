package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//ModifyProfile modifica el perfil de usuario
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	//Create a base User Model
	var t models.Usuario

	//Decode the body
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), 400)
		return
	}

	var status bool

	//Try to modify the profile in DB
	status, err := db.ModifyProfile(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el Perfil "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Datos Incorrectos", 400)
		return
	}

	//Return a 200 code if everything is fine
	w.WriteHeader(http.StatusCreated)
}
