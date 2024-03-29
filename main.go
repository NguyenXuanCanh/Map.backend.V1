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
	"github.com/NguyenXuanCanh/go-starter/api/notification"
	"github.com/NguyenXuanCanh/go-starter/api/profile"
	"github.com/NguyenXuanCanh/go-starter/api/routing"
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

func UpdateVehicle(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: vehicles.UpdateVehicle(request),
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
		// Data: trips.CreateTrip(writer, request, id),
		Data: routing.CreateTrip(id),
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
		Data: routing.GetTrips(id),
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

func GetHistoryByFilterData(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	var id = request.URL.Query().Get("account_id")
	var start_date = request.URL.Query().Get("start_date")
	var end_date = request.URL.Query().Get("end_date")
	response := Response{
		Status: "OK",
		Data:   history.GetHistoryByFilterData(id, start_date, end_date),
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

func GetNotificationByAccountId(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	var id = ps.ByName("id")
	response := Response{
		Status: "OK",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: notification.GetNotificationByAccountId(id),
		// Data: packages.Main(),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func AddNotification(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	response := Response{
		Status: "OK/",
		// Data:   compute_routes.GetComputeRoutes(),
		Data: notification.AddNotification(request),
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

func GetTest(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	var id = ps.ByName("id")
	response := Response{
		Status: "OK",
		Data:   routing.CreateTrip(id),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func UpdateStepTrip(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	response := Response{
		Status: "OK",
		Data:   routing.UpdateStepTrip(request),
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func DeleteUnavailableHistory(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// api.TestGetAPI()
	response := Response{
		Status: "OK",
		Data:   history.DeleteUnavailableHistory(),
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

	router.GET("/vehicle", GetVehicle)
	router.GET("/vehicle/:id", GetVehicle)
	router.POST("/vehicle_add", AddVehicle)
	router.POST("/vehicle_update", UpdateVehicle)

	router.POST("/history_add", AddHistory)
	router.GET("/history", GetHistoryByFilterData)

	router.POST("/user_image_update", UpdateImageProfile)

	router.GET("/trip/:id", GetTrip)
	router.POST("/update_steps_trip", UpdateStepTrip)
	router.GET("/generate_trip/:id", GenTrip)
	router.GET("/GetTest", GetTest)

	router.GET("/clockin_status/:id", GetClockinStatus)
	router.GET("/clockin_status_update/:id/:status", UpdateClockinStatus)

	router.GET("/delete_unavailable_history", DeleteUnavailableHistory)

	router.GET("/notification/:id", GetNotificationByAccountId)
	router.POST("/notification_add", AddNotification)

	log.Fatal(http.ListenAndServe(":8080", router))
	// err := http.ListenAndServe("localhost:8080", router)

	// if err != nil {
	// 	return
	// }
}
