package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la DB */
var MongoCN = ConectarDB()

//Delete in deploy
const dbURL string = "mongodb+srv://root:root@cluster0-jjbii.mongodb.net/merntask"

var clientOptions = options.Client().ApplyURI(dbURL)

/*ConectarDB funcion que se conecta a la base de datos */
func ConectarDB() *mongo.Client {
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
