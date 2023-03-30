package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/NguyenXuanCanh/go-starter/api/packages"
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

func CreatePackageWayPoint() []types.Package {
	data := packages.GetPackageWaiting()
	for i := 0; i < len(data); i++ {
		location := CreateLocation(data[i].Description)
		data[i].Location = location
	}
	return data
}
