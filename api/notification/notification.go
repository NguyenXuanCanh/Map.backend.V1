package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNotificationById(id string) types.Notification {
	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	var database = connection.UseDatabase()
	var notification types.Notification
	filter := bson.D{{"account_id", id_int}}
	err_db := database.Collection("notification").FindOne(context.TODO(), filter).Decode(&notification)
	if err_db != nil {
		fmt.Println(err_db)
	}
	return notification
}

func GetNotificationByAccountId(id string) []types.Notification {
	filter := bson.D{{"account_id", id}}
	var database = connection.UseDatabase()
	var notifications []types.Notification
	cur, err := database.Collection("notification").Find(context.Background(), filter)
	if err != nil {
		return notifications
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To get the bson bytes value use cursor.Current
		var raw types.Notification
		bsonBytes, _ := bson.Marshal(cur.Current)
		bson.Unmarshal(bsonBytes, &raw)

		notifications = append(notifications, raw)
	}
	if err := cur.Err(); err != nil {
		return notifications
	}
	// return json.NewEncoder(response).Encode(notifications)
	return notifications
}

func AddNotification(request *http.Request) any {
	decoder := json.NewDecoder(request.Body)
	var notification_add types.Notification
	errDecode := decoder.Decode(&notification_add)
	if errDecode != nil {
		panic(errDecode)
	}

	var database = connection.UseDatabase()

	result, err := database.Collection("notification").InsertOne(context.Background(), notification_add)
	if err != nil {
		log.Fatal(err)
	}

	// return json.NewEncoder(response).Encode(vehicles)
	return result
}
