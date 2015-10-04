package models

type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
