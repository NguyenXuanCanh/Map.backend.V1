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

func GetHistoryByFilterData(id string, start_date string, end_date string) []types.HistoryRes {
	// id_account, err := strconv.Atoi(id)
	// start_date, err_start := time.Parse("02/01/2006", start_date)
	// end_date, err_end := time.Parse("02/01/2006", end_date)
	// if err_start != nil {
	// 	fmt.Println(err_start)
	// }
	// if err_end != nil {
	// 	fmt.Println(err_end)
	// }
	t, err := time.Parse("2006-01-02", start_date)
	if err != nil {
		fmt.Println(err)
	}
	e_t, e_err := time.Parse("2006-01-02", end_date)
	if e_err != nil {
		fmt.Println(err)
	}
	s_date := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	e_date := time.Date(e_t.Year(), e_t.Month(), e_t.Day()+1, 0, 0, 0, 0, time.UTC)

	filter := bson.M{
		"account_id": id,
		"date": bson.M{
			"$lte": e_date,
			"$gte": s_date,
		},
	}

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

func DeleteUnavailableHistory() any {
	var database = connection.UseDatabase()
	filter := bson.M{"package_id": bson.D{{"$in", 0}}}
	if _, err := database.Collection("history").DeleteMany(context.TODO(), filter); err != nil {
		// handle error
	}
	return ""
}
