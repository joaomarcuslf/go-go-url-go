package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"

	configs "github.com/joaomarcuslf/go-go-url-shortener/configs"
)

type StorageService struct {
	RedisClient *redis.Client
}

var (
	storeService = &StorageService{}
)

const CacheDuration = 8760 * time.Hour

func InitializeStore(redisConfig *configs.Redis) (*StorageService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()

	if err != nil {
		return nil, fmt.Errorf("Error init Redis: %v", err)
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)

	storeService.RedisClient = redisClient

	return storeService, nil
}

func SaveUrlMapping(shortUrl string, originalUrl string) error {
	_, err := storeService.RedisClient.Get(shortUrl).Result()

	if err == nil {
		return fmt.Errorf("Key already exists | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl)
	}

	err = storeService.RedisClient.Set(shortUrl, originalUrl, CacheDuration).Err()

	if err != nil {
		return fmt.Errorf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl)
	}

	return nil
}

func RetrieveInitialUrl(shortUrl string) (string, error) {
	result, err := storeService.RedisClient.Get(shortUrl).Result()

	if err != nil {
		return "", fmt.Errorf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl)
	}

	return result, nil
}
