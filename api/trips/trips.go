package trips

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/config"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
)

type SubmitData struct {
	Vehicles    []types.Vehicle `json:"vehicles"`
	Jobs        []types.Job     `json:"jobs"`
	VehicleType string          `json:"vehicle_type"`
}

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
	vehicle.Capacity = append(vehicle.Capacity, 15)
	vehicles = append(vehicles, vehicle)

	//get all waypoints
	var way_points = CreatePackageWayPoint()
	//create job for trips
	//now from packages > create jobs
	var jobs []types.Job
	for i := 0; i < len(way_points); i++ {
		var job types.Job
		job.Id = i + 1
		job.Location = way_points[i].Location
		job.Amount = append(job.Amount, 1)
		job.Description = way_points[i].Description
		jobs = append(jobs, job)
	}

	var submitData SubmitData
	submitData.Jobs = jobs
	submitData.Vehicles = vehicles
	// var req_url = "https://maps.vietmap.vn/api/vrp?api-version=1.1&apikey=" + config.API_KEY + "&jobs=" + _api_call.Jobs + "&vehicles=" + _api_call.Vehicles

	return CreateTripRoute(submitData)
}

func CreateTripRoute(data SubmitData) any {
	urlReq := "https://maps.vietmap.vn/api/vrp?api-version=1.1&apikey=" + config.API_KEY
	values := map[string]interface{}{
		"jobs":         data.Jobs,
		"vehicles":     data.Vehicles,
		"vehicle_type": "car",
	}
	json_data, err := json.Marshal(values)
	// fmt.Println("data data: \n", data)
	// fmt.Printf("json data: %s\n", json_data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return ""
	}

	resp, err := http.Post(urlReq, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return ""
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	// fmt.Println(res)
	return res
}
