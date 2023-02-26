package product

import (
	"context"
	"log"
	"net/http"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(response http.ResponseWriter, request *http.Request) []types.Product {
	response.Header().Set("content-type", "application/json")
	var database = connection.UseDatabase()
	cur, err := database.Collection("products").Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var products []types.Product
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var prod types.Product
		err := cur.Decode(&prod)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...

		// To get the bson bytes value use cursor.Current
		var raw types.Product
		bsonBytes, _ := bson.Marshal(cur.Current)
		bson.Unmarshal(bsonBytes, &raw)
		products = append(products, raw)
	}
	if err := cur.Err(); err != nil {
		// return "error"
	}
	// return json.NewEncoder(response).Encode(products)
	return products
}
