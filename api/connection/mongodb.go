package connection

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UseDatabase() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb+srv://root:xuancanh@cluster0.n5nhzrs.mongodb.net/?retryWrites=true&w=majority")
	client, _ := mongo.Connect(ctx, clientOptions)
	database := client.Database("vrp_map")
	return database
}
