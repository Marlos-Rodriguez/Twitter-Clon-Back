package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
)

//LookProfile permite extraer los valores del perfil
func LookProfile(w http.ResponseWriter, r *http.Request) {

	//Get the id from the URL
	ID := r.URL.Query().Get("id")

	//If the ID is too short
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
	}

	//Search the user in the DB
	perfil, err := db.SearchUser(ID)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar el registro "+err.Error(), 400)
		return
	}

	//Response with the User Info i JSON
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
