package database

import (
	"context"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"
)

var ctx = context.Background()

func ConnectRedis() *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DOMAIN") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	db.Set(ctx, "ping", "pong", 0)
	ok, err := db.Get(ctx, "ping").Result()
	if err != nil || ok != "pong" {
		log.Fatal("redis connection error", err)
	}

	color.Green("redis connected")
	return db
}
