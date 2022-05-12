package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"react-golang-todo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getAllTasks(w http.ResponseWriter, r *http.Request) []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M

	for cur.Next(context.Background()) {
		var result bson.M

		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}

		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())

	return results
}

func insertOneTask(task models.ToDoList) {
	insertRes, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single record", insertRes)
}

func taskComplete(id string) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
}

func taskUndo(id string) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
}

func deleteOneTask(id string) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted document:", result.DeletedCount)
}

func deleteAllTasks() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}
