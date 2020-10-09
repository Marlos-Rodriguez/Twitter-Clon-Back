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
	//Write the header for a JSON response
	w.Header().Add("content-type", "application/json")

	//Create base User model
	var t models.Usuario

	//Decode the body of request
	err := json.NewDecoder(r.Body).Decode(&t)

	//If they are a error in the decde
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(), 400)
		return
	}

	//If the email of the body is Empty
	if len(t.Email) == 0 {
		http.Error(w, "El email es obligatorio", 400)
		return
	}

	//Try the login in the DB
	documento, exist := db.IntentoLogin(t.Email, t.Password)

	//If user not exist
	if !exist {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	//Generate a JWT
	jwtKey, err := jwt.GenerateJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrio un error al generar el token"+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	//Response with the JWT in JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
