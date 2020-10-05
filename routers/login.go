package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/jwt"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//Login realiza el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es obligatorio", 400)
		return
	}

	documento, existe := db.IntentoLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrio un error al generar el token"+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
