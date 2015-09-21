package models

type Response struct {
	Data interface{}       `json:"data"`
	Meta map[string]string `json:"meta"`
}
