package vehicles

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

func GetAll() []types.VehicleDB {
	var database = connection.UseDatabase()
	cur, err := database.Collection("vehicles").Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var vehicles []types.VehicleDB
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var prod types.VehicleDB
		err := cur.Decode(&prod)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...

		// To get the bson bytes value use cursor.Current
		var raw types.VehicleDB
		bsonBytes, _ := bson.Marshal(cur.Current)
		bson.Unmarshal(bsonBytes, &raw)
		vehicles = append(vehicles, raw)
	}
	if err := cur.Err(); err != nil {
		// return "error"
	}
	// return json.NewEncoder(response).Encode(vehicles)
	return vehicles
}

func GetVehicleByAccountId(id string) types.VehicleDB {
	// id_int, err := strconv.Atoi(id)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	var database = connection.UseDatabase()
	var vehicle types.VehicleDB
	filter := bson.D{{"account_id", id}}
	err_db := database.Collection("vehicles").FindOne(context.TODO(), filter).Decode(&vehicle)
	if err_db != nil {
		fmt.Println(err_db)
	}
	return vehicle
}

func AddVehicle(request *http.Request) any {
	decoder := json.NewDecoder(request.Body)
	var vehicle_add types.History
	errDecode := decoder.Decode(&vehicle_add)
	if errDecode != nil {
		panic(errDecode)
	}
	var database = connection.UseDatabase()

	result, err := database.Collection("vehicles").InsertOne(context.Background(), vehicle_add)
	if err != nil {
		log.Fatal(err)
	}

	// return json.NewEncoder(response).Encode(vehicles)
	return result
}
