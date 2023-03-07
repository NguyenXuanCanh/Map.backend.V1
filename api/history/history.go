package history

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll() []types.History {
	var database = connection.UseDatabase()
	cur, err := database.Collection("history").Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var histories []types.History
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var prod types.History
		err := cur.Decode(&prod)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...

		// To get the bson bytes value use cursor.Current
		var raw types.History
		bsonBytes, _ := bson.Marshal(cur.Current)
		bson.Unmarshal(bsonBytes, &raw)
		histories = append(histories, raw)
	}
	if err := cur.Err(); err != nil {
		// return "error"
	}
	// return json.NewEncoder(response).Encode(histories)
	return histories
}

func GetHistoryById(id string) types.History {
	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	var database = connection.UseDatabase()
	var history types.History
	filter := bson.D{{"id", id_int}}
	err_db := database.Collection("history").FindOne(context.TODO(), filter).Decode(&history)
	if err_db != nil {
		fmt.Println(err_db)
	}
	return history
}

func Addhistory(response http.ResponseWriter, request *http.Request) any {
	response.Header().Set("content-type", "application/json")
	var database = connection.UseDatabase()

	new_history := types.History{}

	result, err := database.Collection("history").InsertOne(context.Background(), new_history)
	if err != nil {
		log.Fatal(err)
	}

	// return json.NewEncoder(response).Encode(vehicles)
	return result
}
