package models

type Response struct {
	Error
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
