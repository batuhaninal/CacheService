package models

type CacheModel struct {
	Key  string `json:"key"`
	Data []byte `json:"data"`
}
