package helper

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Db_uri        string
	Db_name       string
	Db_collection string
)

func LoadEnv() {
	var ok bool
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Db_uri, ok = os.LookupEnv("DB_URI")
	if !ok {
		log.Fatal("DB_URI not found")
	}

	Db_name, ok = os.LookupEnv("DB_NAME")
	if !ok {
		log.Fatal("DB_NAME not found")
	}

	Db_collection, ok = os.LookupEnv("DB_COLLECTION_NAME")
	if !ok {
		log.Fatal("DB_COLLECTION not found")
	}
}

func CreateDBInstance() *mongo.Collection {
	fmt.Println(Db_uri, Db_name, Db_collection)
	clientOptions := options.Client().ApplyURI(Db_uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(Db_name).Collection(Db_collection)
	return collection
}
