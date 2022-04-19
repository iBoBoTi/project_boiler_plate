package redis

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func redisConnect(url string, password string) *redis.Client {

	logrus.WithField("connection", url).Info("Connecting to Redis DB")
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	err := client.Ping().Err()
	if err != nil {
		logrus.Fatal(err)
	}
	return client

}
