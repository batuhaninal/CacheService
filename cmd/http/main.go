package main

import (
	"CacheService/config"
	"CacheService/internal/cache/application/services"
	"CacheService/internal/cache/handlers/controllers"
	"CacheService/internal/cache/infrastructure/repository"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
)

func main() {

	redisDB, _ := strconv.Atoi(config.Get("REDIS_DB"))

	redisRepo := repository.NewRedisRepository(redis.NewClient(&redis.Options{
		Addr:     config.Get("REDIS_ENV") + config.Get("REDIS_PORT"),
		Password: config.Get("REDIS_PASSWORD"), // no password set
		DB:       redisDB,                      // use default DB
	}))

	redisService := services.NewRedisService(redisRepo)

	a := controllers.NewRedisController(redisService)

	http.HandleFunc("/test", a.Test)

	http.ListenAndServe(":7001", nil)
}
