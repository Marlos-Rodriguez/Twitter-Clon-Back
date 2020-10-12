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
	//Crear Router
	router := mux.NewRouter()

	//Route of Register
	go router.HandleFunc("/registro", middlewares.ChequeoDB(routers.Registro)).Methods("POST")
	//Route of Login
	go router.HandleFunc("/login", middlewares.ChequeoDB(routers.Login)).Methods("POST")
	//Route for See info profile
	go router.HandleFunc("/verPerfil", middlewares.ChequeoDB(middlewares.ValidJWT(routers.LookProfile))).Methods("GET")
	//Route for Modify Profile Info
	go router.HandleFunc("/modificarPerfil", middlewares.ChequeoDB(middlewares.ValidJWT(routers.ModifyProfile))).Methods("PUT")
	//Route for Create a Tweet
	go router.HandleFunc("/tweet", middlewares.ChequeoDB(middlewares.ValidJWT(routers.SaveTwitter))).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Println("Server running in Port: " + PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
