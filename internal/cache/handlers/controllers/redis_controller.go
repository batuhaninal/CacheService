package controllers

import (
	"CacheService/internal/cache/domain/models"
	"CacheService/internal/cache/domain/ports"
	"encoding/json"
	"net/http"
)

type redisController struct {
	redisService ports.ICacheService
}

func NewRedisController(service ports.ICacheService) ports.ICacheController {
	return &redisController{redisService: service}
}

func (rc redisController) Save(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var model models.CacheModel
		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			panic(err.Error())
		}

		rc.redisService.Save(r.Context(), model)

		w.WriteHeader(201)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{"message": "Eklendi"})
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (rc redisController) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		key := r.URL.Query().Get("name")

		if key == "" {
			panic("Null exception")
		}

		rc.redisService.Remove(r.Context(), key)

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Silme islemi basarili!"})
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (rc redisController) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		q := r.URL.Query()

		if q.Get("name") == "" {
			panic("Null exception")
		}

		cacheModel := rc.redisService.Get(r.Context(), q.Get("name"))

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cacheModel)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (rc redisController) Test(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var model Test
		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			panic(err)
		}

		s := []byte(model.Value)

		rc.redisService.Save(r.Context(), models.CacheModel{Key: model.Key, Data: s})

		cacheModel := rc.redisService.Get(r.Context(), model.Key)

		responseModel := Test{Key: cacheModel.Key}

		responseModel.Value = string(cacheModel.Data)

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(responseModel)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

type Test struct {
	Key   string
	Value string
}
