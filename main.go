package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/NguyenXuanCanh/go-starter/api/packages"
	"github.com/NguyenXuanCanh/go-starter/api/product"
	"github.com/NguyenXuanCanh/go-starter/api/trips"
	"github.com/gorilla/mux"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

func getAllProduct(writer http.ResponseWriter, request *http.Request) {
	response := Response{
		Status: "OK",
		Data:   product.GetAll(writer, request),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func getAllPackage(writer http.ResponseWriter, request *http.Request) {
	// api.TestGetAPI()
	response := Response{
		Status: "OK",
		Data:   packages.GetAll(writer, request),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func getTest(writer http.ResponseWriter, request *http.Request) {
	// api.TestGetAPI()
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: trips.CreateTrip(),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := mux.NewRouter()
	fmt.Println("STARTED")

	router.HandleFunc("/getAllProduct", getAllProduct).Methods("GET")
	router.HandleFunc("/getAllPackage", getAllPackage).Methods("GET")
	router.HandleFunc("/getTest", getTest).Methods("GET")
	err := http.ListenAndServe("localhost:8080", router)

	if err != nil {
		return
	}
}
