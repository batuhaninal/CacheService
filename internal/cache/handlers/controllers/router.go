package controllers

import (
	"CacheService/internal/cache/domain/ports"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(controller ports.ICacheController) {

	router := mux.NewRouter()
	baseGroup := router.PathPrefix("/api/v1").Subrouter()

	redisGroup := baseGroup.PathPrefix("/redis").Subrouter()
	redisGroup.Path("/test").Methods("POST").HandlerFunc(controller.Test)
	redisGroup.Path("/").Methods("POST").HandlerFunc(controller.Save)
	redisGroup.Path("/").Methods("GET").Queries("key", "{key}").HandlerFunc(controller.Get)
	redisGroup.Path("/").Methods("DELETE").Queries("key", "{key}").HandlerFunc(controller.Delete)

	http.Handle("/", router)
}
