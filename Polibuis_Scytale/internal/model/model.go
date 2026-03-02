package model

type Request struct {
	Algoritm  string `json:"algoritm"`
	Data      string `json:"data"`
	Language  string `json:"language"`
	Operation string `json:"operation"`
	Key       int    `json:"key"`
}

type Response struct {
	Result string `json:"result"`
	Error  error  `json:"error"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
