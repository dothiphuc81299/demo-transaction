package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"

	"demo-transaction/config"
)

var (
	client *redis.Client
)

// Connect ...
func Connect() {
	envVars := config.GetEnv()

	// Create client
	client = redis.NewClient(&redis.Options{
		Addr:     envVars.Redis.URI,  // use default Addr
		Password: envVars.Redis.Pass, // no password set
		DB:       0,                  // use default DB
	})

	// Ping Redis
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("RedisURI:", envVars.Redis.URI)
	}
	fmt.Println(pong, err)
	fmt.Println("Redis connected to", envVars.Redis.URI)
}

// Get ...
func Get(key string) string {
	value := client.Get(key).Val()
	return value
}

// Set ...
func Set(key string, value string) {
	const setTime = 30000000000

	err := client.Set(key, value, setTime).Err()
	if err != nil {
		log.Println(err)
	}
}
