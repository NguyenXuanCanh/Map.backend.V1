package vietmap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/NguyenXuanCanh/go-starter/config"
)

type Data struct {
	Addresses []string `json:"addresses"`
}

func CreateData() Data {
	// """Creates the data."""
	var data Data
	data.Addresses = []string{
		"159 Hung Phu, phuong 8 quan 8", // depot
		// "273 An Duong Vuong, phuong 3 quan 5",
		// "1283 Huynh Tan Phat, quan 7",
		// "1 Nguyen Bieu Phuong 1 Quan 5 TP HCM",
		// "50 Lac Long Quan Phuong 3 Quan 11",
		// "17 Duong Dinh Nghe Phuong 8 Quan 11",
	}
	return data
}

type Feature struct {
	Type     string `json:"type"`
	geometry struct {
		Type        string `json:"type"`
		coordinates []int
		properties  struct {
			layer        string
			name         string
			housenumber  string
			street       string
			distance     int
			accuracy     string
			region       string
			region_gid   string
			county       string
			county_gid   string
			locality     string
			locality_gid string
			label        string
			address      string
			addendum     struct{}
			block        int
			floor        int
		}
	}
}

type Address struct {
	features []Feature
	Type     string `json:"type"`
	bbox     []int
	License  string
}

type MapType struct {
	data    any
	message string
	code    string
}

func create_distance_matrix(data Data) any {
	addresses := data.Addresses

	// var distance_matrix_duration []any
	// var distance_matrix_distance []any

	var place_detail_matrix []any
	for i := 0; i < len(addresses); i++ {
		var temp any = send_request(addresses[i])
		place_detail_matrix = append(place_detail_matrix, temp)
	}

	// store_location := fmt.Sprintf("%f", place_detail_matrix[0].features[0].geometry.coordinates) + "," + fmt.Sprintf("%f", place_detail_matrix[0].Geometry.Location.Lng)
	// way_points := ""
	// // Send q requests, returning max_rows rows per request.
	// for i := 1; i < len(place_detail_matrix); i++ {
	// 	way_points += fmt.Sprintf("%f", place_detail_matrix[i].Geometry.Location.Lat) + "," + fmt.Sprintf("%f", place_detail_matrix[i].Geometry.Location.Lng)
	// 	if i != len(place_detail_matrix)-1 {
	// 		way_points += ";"
	// 	}
	// }
	// response := send_request_distance_matrix(store_location, way_points)
	// fmt.Print(response)
	// distance_matrix_duration = append(distance_matrix_duration, response)

	return place_detail_matrix
	// return distance_matrix
}

func send_request(str string) any {
	stringReq := strings.Replace(str, " ", "%20", -1)
	url := "https://maps.vietmap.vn/api/search?api-version=1.1&apikey=" + config.API_KEY + "&text=" + stringReq
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var data any
	json.Unmarshal([]byte(body), &data)
	fmt.Println(data)
	return data
}

func send_request_distance_matrix(origin_address string, dest_address string) any {
	// stringReq := strings.Replace(str, " ", "%20", -1)
	url := "https://rsapi.goong.io/trip?origin=" + origin_address + "&waypoints=" + dest_address + "&api_key=" + config.API_KEY
	res, err := http.Get(url)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var data MapType
	json.Unmarshal([]byte(body), &data)
	fmt.Println(data.data)
	return data
}

func Main() any {
	// """Entry point of the program"""
	// # Create the data.
	data := CreateData()
	distance_matrix := create_distance_matrix(data)
	return distance_matrix
}
