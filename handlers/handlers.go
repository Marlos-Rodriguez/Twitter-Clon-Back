package handlers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/middlewares"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/routers"
)

//Handlers seteo mi puerto, el handler y empiezo el servidor
func Handlers() {
	//Crear Router
	app := fiber.New()

	//Use the cors
	app.Use(cors.New())

	//Route of Register
	app.Post("/registro", middlewares.ChequeoDB(), routers.Registro)
	//Route of Login
	app.Post("/login", middlewares.ChequeoDB(), routers.Login)
	//Route for See info profile
	app.Get("/profile", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.LookProfile)
	//Route for Modify Profile Info
	app.Put("/modifyProfile", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.ModifyProfile)
	//Route for Create a Tweet
	app.Post("/tweet", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.SaveTwitter)
	//Route for read tweet of one User
	app.Get("/readTweet", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.ReadTweet)

	//Get the Port from ENV
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "3000"
	}

	log.Println("Server running in Port: " + PORT)

	log.Fatal(app.Listen(":" + PORT))
}
