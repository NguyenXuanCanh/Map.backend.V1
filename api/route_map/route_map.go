package route_map

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NguyenXuanCanh/go-starter/config"
	"github.com/NguyenXuanCanh/go-starter/types"
)

func Main(response http.ResponseWriter, request *http.Request) string {
	response.Header().Set("content-type", "application/json")
	urlReq := "https://maps.vietmap.vn/api/route?api-version=1.1&apikey=" + config.API_ROUTE_KEY + "&point=[10.75908,106.6603]&point=[10.75727,106.65829]&vehicle=car"
	res, err := http.Get(urlReq)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var data types.Route
	json.Unmarshal(body, &data)
	// fmt.Println(data)
	return data.Paths[0].Points
}
