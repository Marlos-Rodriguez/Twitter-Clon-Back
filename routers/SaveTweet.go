package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//SaveTwitter Save the tweet in the DB
func SaveTwitter(w http.ResponseWriter, r *http.Request) {
	//Basic model for decode body
	var message models.Tweet

	//Decode the body for the Tweet
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	//Create the complete tweet model
	register := models.SaveTweet{
		UserID:  IDUsuario,
		Message: message.Message,
		Date:    time.Now(),
	}

	//Insert in the DB
	_, status, err := db.InsertTweet(register)

	if err != nil {
		http.Error(w, "Error al insertar datos"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se grabo el Tweet", 400)
		return
	}

	//Return a HTTP Successfully code
	w.WriteHeader(http.StatusCreated)
}
