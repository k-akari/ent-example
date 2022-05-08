package model

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}
