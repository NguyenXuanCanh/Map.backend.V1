package main

import (
	"encoding/json"
	"fmt"
	"log"

	// "github.com/NguyenXuanCanh/go-starter/api/product"
	// "github.com/NguyenXuanCanh/go-starter/api/route_map"
	// "github.com/NguyenXuanCanh/go-starter/api/routing"
	"github.com/NguyenXuanCanh/go-starter/api/clockin"
	"github.com/NguyenXuanCanh/go-starter/api/history"
	"github.com/NguyenXuanCanh/go-starter/api/profile"
	"github.com/NguyenXuanCanh/go-starter/api/trips"
	"github.com/NguyenXuanCanh/go-starter/api/vehicles"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}
type Product struct {
	Id       int     `json:"_id"`
	Name     string  `json:"name"`
	Weight   int     `json:"weight"`
	Size     int     `json:"size"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

// func getAllProduct(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
// 	response := Response{
// 		Status: "OK",
// 		Data:   product.GetAll(writer, request),
// 	}
// 	err := json.NewEncoder(writer).Encode(response)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func GetVehicle(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	var id = ps.ByName("id")
	var data any
	if id == "" {
		data = vehicles.GetAll()
	} else {
		data = vehicles.GetVehicleByAccountId(id)
	}
	response := Response{
		Status: "OK",
		Data:   data,
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func AddVehicle(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: vehicles.AddVehicle(request),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

// func UpdatePackage(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
// 	var id = ps.ByName("id")
// 	var status = ps.ByName("status")
// 	response := Response{
// 		Status: "OK",
// 		// Data:   compute_routes.GetComputeRoutes(),
// 		Data: packages.UpdatePackageStatus(id, status),
// 		// Data: packages.Main(),
// 	}
// 	err := json.NewEncoder(writer).Encode(response)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func GenTrip(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	var id = ps.ByName("id")
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: trips.CreateTrip(writer, request, id),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetTrip(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	var id = ps.ByName("id")
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: trips.GetTrips(id),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetClockinStatus(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	var id = ps.ByName("id")
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: clockin.GetClockinStatus(id),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func UpdateClockinStatus(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	var id = ps.ByName("id")
	var status = ps.ByName("status")

	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: clockin.UpdateClockinStatus(id, status),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetHistoryByAccountId(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	var id = ps.ByName("id")
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: history.GetHistoryByAccountId(id),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func AddHistory(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: history.AddHistory(request),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func UpdateImageProfile(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: profile.UpdateImageProfile(request),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// router := mux.NewRouter()
	router := httprouter.New()

	fmt.Println("STARTED")

	// router.GET("/getAllProduct", getAllProduct)
	// router.GET("/package_update/:id/:status", UpdatePackage)
	// router.GET("/getRouting", getRouting)
	router.GET("/vehicle", GetVehicle)
	router.GET("/vehicle/:id", GetVehicle)
	router.POST("/vehicle_add", AddVehicle)
	router.POST("/history_add", AddHistory)
	router.POST("/user_image_update", UpdateImageProfile)
	router.GET("/history/:id", GetHistoryByAccountId)
	router.GET("/trip/:id", GetTrip)
	router.GET("/generate_trip/:id", GenTrip)
	router.GET("/clockin_status/:id", GetClockinStatus)
	router.GET("/update_clockin_status/:id/:status", UpdateClockinStatus)

	log.Fatal(http.ListenAndServe(":8080", router))
	// err := http.ListenAndServe("localhost:8080", router)

	// if err != nil {
	// 	return
	// }
}
