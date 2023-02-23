package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/NguyenXuanCanh/go-starter/api/packages"
	"github.com/NguyenXuanCanh/go-starter/api/product"
	"github.com/NguyenXuanCanh/go-starter/api/vietmap"
	"github.com/gorilla/mux"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func getAllProduct(writer http.ResponseWriter, request *http.Request) {
	// api.TestGetAPI()
	response := Response{
		Status: "OK",
		Data:   product.GetAll(),
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
		Data:   packages.GetAll(),
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
		Data: vietmap.Main(),
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
