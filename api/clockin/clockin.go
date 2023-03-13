package clockin

import (
	"context"
	"log"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"go.mongodb.org/mongo-driver/bson"
)

func GetClockinStatus(id string) any {
	filter := bson.D{{"account_id", id}}
	var database = connection.UseDatabase()
	var res any
	err := database.Collection("clockin_status").FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return false
	}
	// return json.NewEncoder(response).Encode(packages)
	return res
}

func UpdateClockinStatus(id string, status string) any {
	// filter := bson.D{{"id", id}}
	// update := bson.D{{"$set", bson.D{{"status", status}}}}
	var insert_data struct {
		Id     string
		Status string
	}
	insert_data.Id = id
	insert_data.Status = status

	var database = connection.UseDatabase()
	result, err := database.Collection("clockin_status").InsertOne(context.TODO(), insert_data)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
