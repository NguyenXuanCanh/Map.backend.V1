package packages

import (
	"context"
	"log"

	"github.com/NguyenXuanCanh/go-starter/api/connection"
	"github.com/NguyenXuanCanh/go-starter/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll() []types.Package {
	var database = connection.UseDatabase()
	cur, err := database.Collection("packages").Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var packages []types.Package
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var prod types.Package
		err := cur.Decode(&prod)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...

		// To get the bson bytes value use cursor.Current
		var raw types.Package
		bsonBytes, _ := bson.Marshal(cur.Current)
		bson.Unmarshal(bsonBytes, &raw)
		packages = append(packages, raw)
	}
	if err := cur.Err(); err != nil {
		// return "error"
	}
	// return json.NewEncoder(response).Encode(packages)
	return packages
}

func GetPackageWaiting() []types.Package {
	filter := bson.D{{"status", "waiting"}}
	var database = connection.UseDatabase()
	cur, err := database.Collection("packages").Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var packages []types.Package
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var prod types.Package
		err := cur.Decode(&prod)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...

		// To get the bson bytes value use cursor.Current
		var raw types.Package
		bsonBytes, _ := bson.Marshal(cur.Current)
		bson.Unmarshal(bsonBytes, &raw)
		packages = append(packages, raw)
	}
	if err := cur.Err(); err != nil {
		// return "error"
	}
	// return json.NewEncoder(response).Encode(packages)
	return packages
}

func UpdatePackageStatus(id int, status string) any {
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", bson.D{{"status", status}}}}

	var database = connection.UseDatabase()
	result, err := database.Collection("packages").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
