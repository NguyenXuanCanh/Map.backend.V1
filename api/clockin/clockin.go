package clockin

import (
	"context"
	"log"
	"time"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetClockinStatus(id string) any {
	filter := bson.D{{"account_id", id}}
	var database = connection.UseDatabase()
	var clockin_status struct {
		Status     string `json:"status"`
		Account_id string `json:"account_id"`
	}
	err := database.Collection("clockin_status").FindOne(context.TODO(), filter).Decode(&clockin_status)
	if err != nil {
		return false
	}
	// return json.NewEncoder(response).Encode(packages)
	return clockin_status
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

	var noti types.Notification
	noti.Account_id = id
	noti.Type = "clockin"
	noti.Time = primitive.NewDateTimeFromTime(time.Now())
	_, errNoti := database.Collection("notification").InsertOne(context.Background(), noti)

	if errNoti != nil {
		log.Fatal(errNoti)
	}

	// return json.NewEncoder(response).Encode(vehicles)
	return result
}
