package config

import (
	"context"
	"hendralijaya/user-management-project/helper"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() *mongo.Client {
	err := godotenv.Load()
	helper.PanicIfError(err)
	uri := os.Getenv("MONGODB_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	helper.PanicIfError(err)
	return client
}

func CloseMongoDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client.Disconnect(ctx)
}