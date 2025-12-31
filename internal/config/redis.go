package config

// Placeholder for future Redis configuration
// Uncomment when Redis is needed

/*
import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	ctx := context.Background()
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect to Redis")
	}
}

func GetRedis() *redis.Client {
	return RedisClient
}
*/
