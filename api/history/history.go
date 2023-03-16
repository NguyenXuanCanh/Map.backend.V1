package history

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/api/packages"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		// var prod types.History
		// err := cur.Decode(&prod)
		// if err != nil {
		// 	log.Fatal(err)
		// }
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

func GetHistoryByAccountId(id string) []types.HistoryRes {
	// id_account, err := strconv.Atoi(id)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	filter := bson.D{{"account_id", id}}
	var database = connection.UseDatabase()
	cur, err := database.Collection("history").Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var histories []types.HistoryRes
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		// var prod types.HistoryRes
		// err := cur.Decode(&prod)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// do something with result...

		// To get the bson bytes value use cursor.Current
		var raw types.HistoryRes
		bsonBytes, _ := bson.Marshal(cur.Current)
		bson.Unmarshal(bsonBytes, &raw)

		var package_get types.Package
		filter_package := bson.D{{"id", raw.Package_id}}
		err_package := database.Collection("packages").FindOne(context.TODO(), filter_package).Decode(&package_get)
		if err_package != nil {
			log.Fatal(err_package)
		}
		raw.Status = package_get.Status
		raw.Description = package_get.Description
		raw.Volume = package_get.Volume
		raw.Weight = package_get.Weight
		raw.Total = package_get.Total

		histories = append(histories, raw)
	}
	if err := cur.Err(); err != nil {
		// return "error"
	}
	// return json.NewEncoder(response).Encode(histories)
	return histories
}

func AddHistory(request *http.Request) any {
	decoder := json.NewDecoder(request.Body)
	var history_add types.History
	errDecode := decoder.Decode(&history_add)
	if errDecode != nil {
		panic(errDecode)
	}

	var database = connection.UseDatabase()

	packages.UpdatePackageStatus(history_add.Package_id, "success")
	var noti types.Notification
	noti.Account_id = history_add.Account_id
	noti.Type = "deliveried"
	noti.Time = primitive.NewDateTimeFromTime(time.Now())
	_, errNoti := database.Collection("notification").InsertOne(context.Background(), noti)

	if errNoti != nil {
		log.Fatal(errNoti)
	}

	result, err := database.Collection("history").InsertOne(context.Background(), history_add)
	if err != nil {
		log.Fatal(err)
	}

	// return json.NewEncoder(response).Encode(vehicles)
	return result
}
