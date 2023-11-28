package models

type CacheModel struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
}
