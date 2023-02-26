package trips

import (
	"context"
	"log"
	"net/http"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/config"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTrips(response http.ResponseWriter, request *http.Request) []types.Trip {
	response.Header().Set("content-type", "application/json")
	var database = connection.UseDatabase()
	cur, err := database.Collection("trips").Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var trips []types.Trip
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var prod types.Trip
		err := cur.Decode(&prod)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...

		// To get the bson bytes value use cursor.Current
		var raw types.Trip
		bsonBytes, _ := bson.Marshal(cur.Current)
		bson.Unmarshal(bsonBytes, &raw)
		trips = append(trips, raw)
	}
	if err := cur.Err(); err != nil {
		// return "error"
	}
	// return json.NewEncoder(response).Encode(trips)
	return trips
}

func CreateTrip() any {
	var vehicles []types.Vehicle
	var vehicle types.Vehicle
	vehicle.Id = 1
	vehicle.Start = config.GetDefaultStoreLocation()
	vehicle.End = config.GetDefaultStoreLocation()
	vehicles[0] = vehicle

	return vehicle
}
