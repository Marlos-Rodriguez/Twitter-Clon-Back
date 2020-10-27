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

	//User Routes
	app.Post("/registro", middlewares.ChequeoDB(), routers.Registro)
	app.Post("/login", middlewares.ChequeoDB(), routers.Login)
	app.Get("/profile", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.LookProfile)
	app.Put("/modifyProfile", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.ModifyProfile)

	//Relation Routes
	app.Post("/createRelation", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.SaveRelation)
	app.Delete("/deleteRelation", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.DeleteRelation)
	app.Get("/requestRelation", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.RequestRelation)
	app.Get("/listUsers", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.ListUsers)

	//Tweets Route
	app.Post("/tweet", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.SaveTwitter)
	app.Get("/readTweet", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.ReadTweet)
	app.Delete("/deleteTweet", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.DeleteTweet)

	//Images Route
	app.Post("/uploadAvatar", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.UploadAvatar)
	app.Post("/uploadBanner", middlewares.ChequeoDB(), middlewares.ValidJWT(), routers.UploadBanner)

	//Static Images
	app.Static("/uploads/", "./uploads")

	//Get the Port from ENV
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "3000"
	}

	log.Println("Server running in Port: " + PORT)

	log.Fatal(app.Listen(":" + PORT))
}
