package product

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NguyenXuanCanh/go-starter/types"
)

func GetAll() []types.Product {
	url := "https://63d3d0f5a93a149755b37951.mockapi.io/api/product"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var data []types.Product
	json.Unmarshal(body, &data)
	return data
}
