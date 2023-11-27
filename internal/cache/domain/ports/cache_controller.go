package ports

import "net/http"

type ICacheController interface {
	Save(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Test(w http.ResponseWriter, r *http.Request)
}
