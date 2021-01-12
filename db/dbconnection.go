package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection is a var that contains db client
var MongoConnection = ConnectDB()
var clientOptions =  options.Client().ApplyURI("mongodb+srv://maurofuentesgo:CWcxu8ZTfhrxkof8@cluster0.zt588.mongodb.net/test")

// ConnectDB is a function that return a client if db connect is successfully
func ConnectDB() *mongo.Client  {
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

	log.Println("Connected to DB")

	return client
}

// ConnectionOn is a function that check db connection
func ConnectionOn() bool  {
	on := true

	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		on = false
	}

	return on
}