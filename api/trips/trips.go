package trips

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/api/packages"
	"github.com/NguyenXuanCanh/go-starter/config"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
)

type SubmitData struct {
	Vehicles    []types.Vehicle `json:"vehicles"`
	Jobs        []types.Job     `json:"jobs"`
	VehicleType string          `json:"vehicle_type"`
}

func GetTrips(id string) types.TripDB {
	var database = connection.UseDatabase()

	var tripRes types.TripRes
	filter := bson.D{{"account_id", id}}
	err := database.Collection("trips").FindOne(context.TODO(), filter).Decode(&tripRes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tripRes)
	// return json.NewEncoder(response).Encode(trips)
	return tripRes.Trip_data
}

func CreateTrip(response http.ResponseWriter, request *http.Request, id string) any {
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
		job.Id = way_points[i].Id
		job.Location = way_points[i].Location
		job.Amount = append(job.Amount, 1)
		job.Description = way_points[i].Description
		jobs = append(jobs, job)
		packages.UpdatePackageStatus(job.Id, "delivering")
	}

	var submitData SubmitData
	submitData.Jobs = jobs
	submitData.Vehicles = vehicles
	// var req_url = "https://maps.vietmap.vn/api/vrp?api-version=1.1&apikey=" + config.API_KEY + "&jobs=" + _api_call.Jobs + "&vehicles=" + _api_call.Vehicles

	var res = CreateTripRoute(submitData)
	// fmt.Println(res)
	return Save(res, id)
}

func Save(res any, id string) any {
	var database = connection.UseDatabase()
	var trip_add struct {
		Account_id string
		Trip_data  any
	}
	trip_add.Account_id = id
	trip_add.Trip_data = res
	// packages.UpdatePackageStatus(trip_add.Id, "success")

	result, err := database.Collection("trips").InsertOne(context.Background(), trip_add)
	if err != nil {
		log.Fatal(err)
	}

	// return json.NewEncoder(response).Encode(vehicles)
	return result
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
	// return res
	return res
}
