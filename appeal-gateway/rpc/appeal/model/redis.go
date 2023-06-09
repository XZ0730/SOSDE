package model

import (
	"appeal/internal/config"

	"github.com/go-redis/redis/v8"
)

func InitRedis(c *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: "147258", // no password set
		DB:       1,        // use default DB
	})
}
func InitRedis2(c *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: "147258", // no password set
		DB:       2,        // use default DB
	})
}
func InitRedis3(c *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: "147258", // no password set
		DB:       3,        // use default DB
	})
}
func InitRedis6(c *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: "147258", // no password set
		DB:       6,        // use default DB
	})
}
func InitRedis7(c *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: "147258", // no password set
		DB:       7,        // use default DB
	})
}
