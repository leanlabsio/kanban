package models

type Response struct {
	Data interface{} `json:"data"`
	Event string     `json:"event"`
}
