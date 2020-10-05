package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/middlewares"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/routers"
)

//Handlers seteo mi puerto, el handler y empiezo el servidor
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Println("Server running in Port: " + PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
