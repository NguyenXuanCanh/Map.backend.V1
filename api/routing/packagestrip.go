package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/NguyenXuanCanh/go-starter/config"
	"github.com/NguyenXuanCanh/go-starter/types"
)

type Data struct {
	Addresses []string `json:"addresses"`
}

type Feature struct {
	Type     string
	Geometry struct {
		Type        string
		Coordinates types.Location
	}
	Properties struct {
		Layer        string
		Name         string
		Housenumber  string
		Street       string
		Distance     float64
		Accuracy     string
		Region       string
		Region_gid   string
		County       string
		County_gid   string
		Locality     string
		Locality_gid string
		Label        string
		Address      string
		Addendum     struct{}
		Block        int
		Floor        int
	}
	Bbox types.Location
	Id   string
}

type Address struct {
	Features []Feature
	Type     string
	bbox     types.Location
	License  string
}

type MapType struct {
	Code    string
	Message string
	Data    Address
}

type GoongRes struct {
	Status    string `json:"status"`
	Plus_code any
	Results   []struct {
		Address_components []any  `json:"address_components"`
		Formatted_address  string `json:"formatted_address"`
		Geometry           struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
		Place_id  string `json:"place_id"`
		Reference string `json:"reference"`
		Plus_code struct {
			Compound_code string `json:"compound_code"`
			Global_code   string `json:"global_code"`
		} `json:"plus_code"`
		Types []any `json:"types"`
	} `json:"results"`
}

func CreateLocation(str string) types.Location {
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
	var data MapType
	json.Unmarshal(body, &data)
	// fmt.Println(data)
	arrLength := len(data.Data.Features)
	if arrLength > 0 {
		return data.Data.Features[0].Geometry.Coordinates
	} else {
		return nil
	}
}

func CreateLocationGoong(str string) types.Location {
	stringReq := strings.Replace(str, " ", "%20", -1)
	url := "https://rsapi.goong.io/geocode?address=" + stringReq + "&api_key=" + "rvWoa97j8PhzM5VUA0cr1IGNNNm5X81HoIN8GET6"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var data GoongRes
	json.Unmarshal(body, &data)
	// fmt.Println(data)
	arrLength := len(data.Results)
	if arrLength > 0 {
		var location types.Location
		location = append(location, data.Results[0].Geometry.Location.Lng)
		location = append(location, data.Results[0].Geometry.Location.Lat)
		return location
	} else {
		return nil
	}
}

// func CreatePackageWayPoint() []types.Package {
// 	data := packages.GetPackageWaiting()
// 	for i := 0; i < len(data); i++ {
// 		location := CreateLocationGoong(data[i].Description)
// 		data[i].Location = location
// 	}
// 	return data
// }
