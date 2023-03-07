package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/NguyenXuanCanh/go-starter/api/product"
	"github.com/NguyenXuanCanh/go-starter/api/route_map"
	"github.com/NguyenXuanCanh/go-starter/api/routing"
	"github.com/NguyenXuanCanh/go-starter/api/trips"
	"github.com/NguyenXuanCanh/go-starter/api/vehicles"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

func getAllProduct(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	response := Response{
		Status: "OK",
		Data:   product.GetAll(writer, request),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func getVehicle(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	var id = ps.ByName("id")
	var data any
	if id == "" {
		data = vehicles.GetAll()
	} else {
		data = vehicles.GetVehicleById(id)
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
		Data: vehicles.AddVehicle(writer, request),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func getTrip(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: trips.CreateTrip(writer, request),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func getRouting(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: routing.Main(),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func getRoute(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: route_map.Main(writer, request),
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
	// router.GET("/getAllPackage", getAllPackage)
	// router.GET("/getRouting", getRouting)
	router.GET("/vehicle", getVehicle)
	router.GET("/vehicle/:id", getVehicle)
	router.GET("/vehicle_add", AddVehicle)
	router.GET("/trip", getTrip)

	log.Fatal(http.ListenAndServe(":8080", router))
	// err := http.ListenAndServe("localhost:8080", router)

	// if err != nil {
	// 	return
	// }
}
