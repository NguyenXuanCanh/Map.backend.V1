package packages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/NguyenXuanCanh/go-starter/config"
	"github.com/NguyenXuanCanh/go-starter/types"
)

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func GetAll() [3]types.Package {
	url := "https://63d3d0f5a93a149755b37951.mockapi.io/api/product"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var productList []types.Product
	json.Unmarshal(body, &productList)

	var packageList [3]types.Package
	packageList[0].Id = "1"
	packageList[0].Location = "ChIJD-2wOQAvdTERQMcHMUK0xPQ"
	packageList[0].Products = productList[0:3]

	packageList[1].Id = "2"
	packageList[1].Location = "ChIJg0HGgRwvdTERPHWPenqdENM"
	packageList[1].Products = productList[3:6]

	packageList[2].Id = "3"
	packageList[2].Location = "ChIJyw_gs4QvdTER9hMcse7Ed9w"
	packageList[2].Products = productList[6:10]

	return packageList
}

type Data struct {
	Addresses []string `json:"addresses"`
}

func CreateData() Data {
	// """Creates the data."""
	var data Data
	data.Addresses = []string{
		"159 Hung Phu, phuong 8 quan 8", // depot
		"273 An Duong Vuong, phuong 3 quan 5",
		"1283 Huynh Tan Phat",
		"1 Nguyen Bieu Phuong 1 Quan 5 TP HCM",
		"50 Lac Long Quan Phuong 3 Quan 11",
		"17 Duong Dinh Nghe Phuong 8 Quan 11",
	}
	return data
}

type Address struct {
	AddressComponents []struct {
		LongName  string `json:"long_name"`
		ShortName string `json:"short_name"`
	} `json:"address_components"`
	FormattedAddress string `json:"formatted_address"`
	Geometry         struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
	PlaceId   string `json:"place_id"`
	Reference string `json:"reference"`
	PlusCode  struct {
		CompoundCode string `json:"compound_code"`
		GlobalCode   string `json:"global_code"`
	} `json:"plus_code"`
}
type MapType struct {
	Results []Address `json:"results"`
	Status  string
}

func create_distance_matrix(data Data) any {
	addresses := data.Addresses
	// Distance Matrix API only accepts 100 elements per request, so get rows in multiple requests.
	// max_elements := 100
	// num_addresses := len(addresses) // 16 in this example.
	// Maximum number of rows that can be computed per request (6 in this example).
	// max_rows := int64(max_elements) // num_addresses
	// num_addresses = q * max_rows + r (q = 2 and r = 4 in this example).
	// q, r := divmod(int64(num_addresses), int64(max_rows))
	// dest_addresses := addresses
	// origin_addresses := addresses
	var distance_matrix_duration []any
	// var distance_matrix_distance []any

	var place_detail_matrix []Address
	for i := 0; i < len(addresses)-1; i++ {
		var temp Address = send_request(addresses[i])
		fmt.Println(temp)
		place_detail_matrix = append(place_detail_matrix, temp)
	}

	store_location := fmt.Sprintf("%f", place_detail_matrix[0].Geometry.Location.Lat) + "," + fmt.Sprintf("%f", place_detail_matrix[0].Geometry.Location.Lng)
	way_points := ""
	// Send q requests, returning max_rows rows per request.
	for i := 1; i < len(place_detail_matrix); i++ {
		way_points += fmt.Sprintf("%f", place_detail_matrix[i].Geometry.Location.Lat) + "," + fmt.Sprintf("%f", place_detail_matrix[i].Geometry.Location.Lng)
		if i != len(place_detail_matrix)-1 {
			way_points += ";"
		}
	}
	response := send_request_distance_matrix(store_location, way_points)
	fmt.Print(response)
	distance_matrix_duration = append(distance_matrix_duration, response)

	// // Get the remaining remaining r rows, if necessary.
	// if r > 0 {
	// 	origin_addresses := addresses[q*max_rows : q*max_rows+r]
	// 	response := send_request(origin_addresses, dest_addresses)
	// 	fmt.Println(response)
	// 	distance_matrix = append(distance_matrix, build_distance_matrix(response))
	// }
	return distance_matrix_duration
	// return distance_matrix
}

func send_request(str string) Address {
	stringReq := strings.Replace(str, " ", "%20", -1)
	url := "https://rsapi.goong.io/geocode?api_key=" + config.API_KEY + "&address=" + stringReq
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var data MapType
	json.Unmarshal([]byte(body), &data)
	return data.Results[0]
}

func send_request_distance_matrix(origin_address string, dest_address string) any {
	// stringReq := strings.Replace(str, " ", "%20", -1)
	url := "https://rsapi.goong.io/trip?origin=" + origin_address + "&waypoints=" + dest_address + "&api_key=" + config.API_KEY
	res, err := http.Get(url)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var data any
	json.Unmarshal([]byte(body), &data)
	return data
}

// func build_address_str(addresses []string) string {
// 	// # Build a pipe-separated string of addresses
// 	var address_str string
// 	for i := 0; i < len(addresses)-1; i++ {
// 		address_str += addresses[i] + "|"
// 	}
// 	address_str += addresses[len(addresses)-1]
// 	return address_str
// }

// func send_request(origin_addresses []string, dest_addresses []string) any {
// 	// """ Build and send request for the given origin and destination addresses."""
// 	requestUrl := "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial"
// 	origin_address_str := build_address_str(origin_addresses)
// 	dest_address_str := build_address_str(dest_addresses)
// 	requestUrl = requestUrl + "&origins=" + origin_address_str + "&destinations=" + dest_address_str + "&key=" + config.API_KEY
// 	// jsonResult := urllib.urlopen(request).read()
// 	// response = json.loads(jsonResult)
// 	// return response
// 	fmt.Println(requestUrl)
// 	method := "GET"

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, requestUrl, nil)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var data any
// 	json.Unmarshal(body, &data)
// 	return data
// }

// func build_distance_matrix(response any) []any {
// 	var distance_matrix []any
// 	// for i, row := range response.rows {
// 	// row_list = [row['elements'][j]['distance']['value'] for j in range(len(row['elements']))]
// 	// distance_matrix.append(row_list)
// 	// }
// 	return distance_matrix
// }

func Main() any {
	// """Entry point of the program"""
	// # Create the data.
	data := CreateData()
	distance_matrix := create_distance_matrix(data)
	return distance_matrix
}
