package keystorage

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/gofiber/storage/redis/v3"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func InitializeRedis() *redis.Storage {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve Redis configuration from environment variables
	host := os.Getenv("REDIS_HOST")
	portStr := os.Getenv("REDIS_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Redis port is not a number" + err.Error() + " " + portStr)
	}
	username := os.Getenv("REDIS_USERNAME")
	password := os.Getenv("REDIS_PASSWORD")
	databaseStr := os.Getenv("REDIS_DATABASE")
	database, err := strconv.Atoi(databaseStr)
	if err != nil {
		fmt.Println("Redis database is not a number" + err.Error() + " " + databaseStr)
	}
	poolSizeStr := os.Getenv("REDIS_POOL_SIZE")
	poolSize, err := strconv.Atoi(poolSizeStr)
	if err != nil {
		fmt.Println("Redis poolsize is not a number" + err.Error() + " " + poolSizeStr)
	}

	store := redis.New(redis.Config{
		Host:      host,
		Port:      port,
		Username:  username,
		Password:  password,
		Database:  database,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  poolSize * runtime.GOMAXPROCS(0),
	})

	return store
}
