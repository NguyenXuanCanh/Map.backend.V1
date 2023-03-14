package clockin

import (
	"context"
	"log"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetClockinStatus(id string) bool {
	filter := bson.D{{"account_id", id}}
	var database = connection.UseDatabase()
	err := database.Collection("clockin_status").FindOne(context.TODO(), filter)
	if err != nil {
		return false
	}
	// return json.NewEncoder(response).Encode(packages)
	return true
}

func UpdateClockinStatus(id string, status string) any {

	var database = connection.UseDatabase()

	var insert_data types.ClockinStatus
	insert_data.Account_id = id
	insert_data.Status = status
	// packages.UpdatePackageStatus(trip_add.Id, "success")

	result, err := database.Collection("clockin_status").InsertOne(context.Background(), insert_data)
	if err != nil {
		log.Fatal(err)
	}

	// return json.NewEncoder(response).Encode(vehicles)
	return result
}
