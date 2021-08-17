package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongoDB() *mongo.Database {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(
			fmt.Sprintf(
				"mongodb://%s:%s@%s:%s",
				os.Getenv("DB_USERNAME"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_DOMAIN"),
				os.Getenv("DB_PORT"),
			),
		),
	)

	if err != nil {
		log.Fatal("mongo connection error:", err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("mongo connection error:", err)
	}

	color.Green("mongoDB connected")

	return client.Database(os.Getenv("DB_NAME"))
}
