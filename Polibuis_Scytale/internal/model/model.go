package model

type request struct {
	Algoritm  string `json:"algoritm"`
	Data      string `json:"data"`
	Language  string `json:"language"`
	Operation string `json:"operation"`
	Key       int    `json:"key"`
}

type response struct {
}

type errorResponse struct {
	Message string `json:"message"`
}
