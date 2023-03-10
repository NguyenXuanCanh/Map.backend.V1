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
	err := database.Collection("packages").FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	// return json.NewEncoder(response).Encode(packages)
	return res
}

func UpdateClockinStatus(id string) any {
	return ""
}
