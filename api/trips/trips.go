package trips

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/api/routing"
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

func CreateTrip(response http.ResponseWriter, request *http.Request) any {
	response.Header().Set("content-type", "application/json")
	var vehicles []types.Vehicle
	var vehicle types.Vehicle
	vehicle.Id = 1
	vehicle.Start = config.GetDefaultStoreLocation()
	vehicle.End = config.GetDefaultStoreLocation()
	vehicles = append(vehicles, vehicle)

	//get all waypoints
	var way_points = routing.Main()
	//create job for trips
	//now from packages > create jobs
	var jobs []types.Job
	for i := 0; i < len(way_points); i++ {
		var job types.Job
		job.Id = i
		job.Location = way_points[i]
		jobs = append(jobs, job)
	}

	var _api_call struct {
		Vehicles []types.Vehicle
		Jobs     []types.Job
	}
	_api_call.Vehicles = vehicles
	_api_call.Jobs = jobs

	// var req_url = "https://maps.vietmap.vn/api/vrp?api-version=1.1&apikey=" + config.API_KEY + "&jobs=" + _api_call.Jobs + "&vehicles=" + _api_call.Vehicles
	url := "https://maps.vietmap.vn/api/vrp?api-version=1.1&apikey={your-apikey}&{point}&point={point}&point={point}&jobs={jobs}&vehicles={vehicles}"
	fmt.Println(url)
	return _api_call
}
