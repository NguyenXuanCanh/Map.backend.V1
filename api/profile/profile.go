package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateImageProfile(request *http.Request) any {
	decoder := json.NewDecoder(request.Body)
	var profile_add types.ProfileImage
	errDecode := decoder.Decode(&profile_add)
	if errDecode != nil {
		panic(errDecode)
	}
	var database = connection.UseDatabase()

	filter := bson.D{{"account_id", profile_add.Account_id}}
	errFind := database.Collection("user_image").FindOne(context.TODO(), filter)
	fmt.Println(profile_add)
	fmt.Println(errFind)
	if errFind != nil {
		log.Fatal(errFind)
		result, err := database.Collection("user_image").InsertOne(context.Background(), profile_add)
		if err != nil {
			log.Fatal(err)
		}
		return result
	}

	// return json.NewEncoder(response).Encode(vehicles)
	return nil
}
