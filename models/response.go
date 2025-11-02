package models

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
	Data      any    `json:"data,omitempty"`
}
