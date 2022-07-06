package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

// Define struct for the Redis client.
type StorageService struct {
	redisClient *redis.Client
}

// Package-level declarations for storeService and Redis context.
var (
	storeService = &StorageService{}
	ctx = context.Background()
)

// Cache duration.
const CacheDuration = 6 * time.Hour

// Initialize the store service and return a store pointer.
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "password",
		DB: 0,
	})

	pong, err := redisClient.Ping(ctx).Result();
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v\n", err));
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong);
	storeService.redisClient = redisClient;
	
	return storeService;
}


// Want to be able to save the mapping between the original URL and the
// generated short URL.
func SaveURLMapping(shortURL string, originalURL string, userID string) {
	err := storeService.redisClient.Set(ctx, shortURL, originalURL, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortURL: %s - originalURL: %s\n", err, shortURL, originalURL))
	}
}

// Should be able to retrieve the initial long URL once the short is provided.
func RetrieveInitialURL(shortURL string) string {
	result, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to retrieve original URL | Error: %v - shortURL: %s\n", err, shortURL))
	}

	return result
}


