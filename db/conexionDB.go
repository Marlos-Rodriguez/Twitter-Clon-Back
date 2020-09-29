package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la DB */
var MongoCN = ConectarDB()

/*ConectarDB funcion que se conecta a la base de datos */
func ConectarDB() *mongo.Client {
	loadEnv()

	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		log.Fatal("Variable de entorno de la DB no encontrada")
	}

	var clientOptions = options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con la DB")
	return client
}

/*CheckConnection es el ping a la DB */
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return false
	}

	return true
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading ENV")
	}
}
