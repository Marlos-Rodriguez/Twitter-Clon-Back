package db

import (
	"context"
	"log"
	"os"

	//Autoload the env
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la DB */
var MongoCN = ConectarDB()

/*ConectarDB funcion que se conecta a la base de datos */
func ConectarDB() *mongo.Client {

	//Get MongoDB URL
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		log.Fatal("Variable de entorno de la DB no encontrada")
	}

	//Options for MongoDB with MongoDB Url
	var clientOptions = options.Client().ApplyURI(dbURL)

	//Conenct with the database
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	//Verify the connection with a ping
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con la DB")
	return client
}

/*CheckConnection es el ping a la DB */
func CheckConnection() bool {
	//Verify the connection with a ping
	if err := MongoCN.Ping(context.TODO(), nil); err != nil {
		return false
	}

	return true
}
