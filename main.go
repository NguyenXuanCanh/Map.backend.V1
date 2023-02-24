package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/api/packages"
	"github.com/NguyenXuanCanh/go-starter/api/vietmap"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"

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

var database = connection.UseDatabase()

func getAllProduct(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	cur, err := database.Collection("products").Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var products []Product
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var prod Product
		err := cur.Decode(&prod)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...

		// To get the bson bytes value use cursor.Current
		var raw Product
		bsonBytes, _ := bson.Marshal(cur.Current)
		bson.Unmarshal(bsonBytes, &raw)
		products = append(products, raw)
	}
	if err := cur.Err(); err != nil {
		// return "error"
	}
	json.NewEncoder(response).Encode(products)
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
