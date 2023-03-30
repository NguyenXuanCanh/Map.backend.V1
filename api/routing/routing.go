package routing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/api/packages"
	"github.com/NguyenXuanCanh/go-starter/config"
	"github.com/NguyenXuanCanh/go-starter/types"
	geo "github.com/kellydunn/golang-geo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Params struct {
	DistanceMatrix    [][]int `json:"distanceMatrix"`
	Demands           []int   `json:"demands"`
	VehicleCapacities []int   `json:"vehicleCapacities"`
	VehicleNumber     int     `json:"vehicleNumber"`
	Depot             int     `json:"depot"`
}

type Solution struct {
	Dropped_package []int `json:"dropped_package"`
	Route           []int `json:"route"`
	Route_distance  []int `json:"route_list"`
	Route_load      []int `json:"route_load"`
	Total_distance  int   `json:"total_distance"`
	Total_load      int   `json:"total_load"`
}

type Response struct {
	Routes         []types.Package `json:"routes"`
	Route_distance []int           `json:"route_list"`
	Route_load     []int           `json:"route_load"`
	Total_distance int             `json:"total_distance"`
	Total_load     int             `json:"total_load"`
}

type TripRes struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Account_id string             `json:"account_id"`
	Trip_data  Response           `json:"trip_data"`
}

func create_distance_matrix(locations []types.Location) [][]int {
	// Tạo ma trận khoảng cách
	distanceMatrix := make([][]int, len(locations))
	for i := range distanceMatrix {
		distanceMatrix[i] = make([]int, len(locations))
	}

	// Tính toán khoảng cách giữa các địa chỉ và lưu vào ma trận
	for i := 0; i < len(locations); i++ {
		for j := 0; j < len(locations); j++ {
			if i == j {
				distanceMatrix[i][j] = 0
			} else {
				geocoder := geo.NewPoint(locations[i][1], locations[i][0])
				geocoder2 := geo.NewPoint(locations[j][1], locations[j][0])
				distanceMatrix[i][j] = int(math.Round(geocoder.GreatCircleDistance(geocoder2) * 1000))
			}
		}
	}

	return distanceMatrix
}

func create_params(packages []types.Package) Params {
	var res Params
	var locations []types.Location

	//init depot
	res.Demands = append(res.Demands, 0)
	locations = append(locations, config.GetDefaultStoreLocation())

	res.VehicleCapacities = append(res.VehicleCapacities, 1500)
	for _, item := range packages {
		res.Demands = append(res.Demands, item.Weight)
		locations = append(locations, item.Location)
	}
	res.DistanceMatrix = create_distance_matrix(locations)

	res.VehicleNumber = 1
	res.Depot = 0
	return res
}

func routing_post(data Params) Solution { //Response
	url := "http://localhost:8081/route"

	values := Params{
		DistanceMatrix:    data.DistanceMatrix,
		Demands:           data.Demands,
		VehicleCapacities: data.VehicleCapacities,
		VehicleNumber:     data.VehicleNumber,
		Depot:             data.Depot,
	}
	json_data, err := json.Marshal(values)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	var res Solution
	json.NewDecoder(resp.Body).Decode(&res)
	// fmt.Println(res)
	// return res
	return res
}

func CreateTripSolution() Response {
	// create_distance_matrix()
	var packages_list = packages.GetPackageWaiting() // get wating only

	for index := range packages_list {
		packages_list[index].Location = CreateLocation(packages_list[index].Description)
	}

	distance_matrix := create_params(packages_list)

	solution := routing_post(distance_matrix)

	var resp Response
	resp.Route_distance = solution.Route_distance
	resp.Route_load = solution.Route_load
	resp.Total_distance = solution.Total_distance
	resp.Total_load = solution.Total_load

	var store types.Package
	store.Id = 0
	store.Location = config.GetDefaultStoreLocation()
	resp.Routes = append(resp.Routes, store)

	for _, item := range solution.Route {
		if item != 0 {
			fmt.Println(packages_list[item]) // log used packages
			// change status to delivering
			resp.Routes = append(resp.Routes, packages_list[item])
			packages.UpdatePackageStatus(packages_list[item].Id, "delivering")
		}
	}

	return resp
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

func CreateTrip(id string) any {
	var res = CreateTripSolution()
	// fmt.Println(res)
	return Save(res, id)
}

func GetTrips(id string) Response {
	var database = connection.UseDatabase()

	var tripRes TripRes
	filter := bson.D{{"account_id", id}}
	err := database.Collection("trips").FindOne(context.TODO(), filter).Decode(&tripRes)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(tripRes)
	// return json.NewEncoder(response).Encode(trips)
	return tripRes.Trip_data
}
