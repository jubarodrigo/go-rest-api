package database

import (
	"context"
	"log"
	"sync"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client
//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error
//Used to execute client creation procedure only once.
var mongoOnce sync.Once
//I have used below constants just to hold required database config's.


const (
	PERSON = "person"
)

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {

	envs := LoadEnv()

	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(envs["CONNECTIONSTRING"])
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func LoadEnv() map[string]string { 

	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error reading .env file")
	  }

	return envs
}



