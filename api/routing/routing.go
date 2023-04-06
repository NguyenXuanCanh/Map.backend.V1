package routing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/api/packages"
	"github.com/NguyenXuanCanh/go-starter/config"
	"github.com/NguyenXuanCanh/go-starter/types"
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
	Steps      int                `json:"steps"`
}

type DistanceResponse struct {
	// Geocoded_waypoints []float64 `json:"geocoded_waypoints"`
	Routes []struct {
		Legs struct {
			Distance struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"distance"`
		} `json:"legs"`
		Overview_polyline struct {
			Points string `json:"points"`
		} `json:"overview_polyline"`
		// Warnings       []float64 `json:"warnings"`
		// Waypoint_order []float64 `json:"waypoint_order"`
	} `json:"routes"`
}

type OSMRes struct {
	Code   string `json:"code"`
	Routes []struct {
		Legs        []int   `json:"legs"`
		Weight_name string  `json:"weight_name"`
		Weight      float64 `json:"weight"`
		Duration    float64 `json:"duration"`
		Distance    float64 `json:"distance"`
	} `json:"routes"`
	Waypoints []struct {
		Hint     string `json:"hint"`
		Distance string `json:"distance"`
		Name     string `json:"name"`
		Location types.Location
	} `json:"waypoints"`
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
				distanceMatrix[i][j] = create_distance(locations[i], locations[j])
				// geocoder := geo.NewPoint(locations[i][1], locations[i][0])
				// geocoder2 := geo.NewPoint(locations[j][1], locations[j][0])
				// distanceMatrix[i][j] = int(math.Round(geocoder.GreatCircleDistance(geocoder2) * 1000))
			}
		}
	}

	return distanceMatrix
}

func create_distance(start types.Location, end types.Location) int {
	startLat := fmt.Sprintf("%f", start[1])
	startLong := fmt.Sprintf("%f", start[0])
	endLat := fmt.Sprintf("%f", end[1])
	endLong := fmt.Sprintf("%f", end[0])
	url := "https://routing.openstreetmap.de/routed-bike/route/v1/driving/" + startLong + "," + startLat + ";" + endLong + "," + endLat + "?overview=false&alternatives=true&steps=false"
	// url := "https://rsapi.goong.io/Direction?origin=" + startLat + "," + startLong + "&destination=" + endLat + "," + endLong + "&vehicle=car&api_key=rvWoa97j8PhzM5VUA0cr1IGNNNm5X81HoIN8GET6"
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var data OSMRes
	json.Unmarshal(body, &data)
	fmt.Println(data)
	if data.Code == "Ok" {
		return int(data.Routes[0].Distance)
	}
	return -1
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

	for _, index := range solution.Route {
		if index != 0 {
			// bug here
			fmt.Println(packages_list[index-1]) // log used packages
			// change status to delivering
			packages.UpdatePackageStatus(packages_list[index-1].Id, "delivering")
			resp.Routes = append(resp.Routes, packages_list[index-1])
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

func GetTrips(id string) TripRes {
	var database = connection.UseDatabase()

	var tripRes TripRes
	filter := bson.D{{"account_id", id}}
	err := database.Collection("trips").FindOne(context.TODO(), filter).Decode(&tripRes)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(tripRes)
	// return json.NewEncoder(response).Encode(trips)

	return tripRes
}

type StepIdTrip struct {
	Account_id string `json:"account_id"`
	Steps      int    `json:"steps"`
}

func UpdateStepTrip(request *http.Request) any {
	decoder := json.NewDecoder(request.Body)
	var step_id StepIdTrip
	errDecode := decoder.Decode(&step_id)
	if errDecode != nil {
		panic(errDecode)
	}
	var database = connection.UseDatabase()

	filter := bson.D{{"account_id", step_id.Account_id}}
	update := bson.D{{"$set", bson.D{{"steps", step_id.Steps}}}}

	result, err := database.Collection("trips").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// return json.NewEncoder(response).Encode(vehicles)
	return result
}
