package model

type Request struct {
	Algoritm  string `json:"algoritm"`
	Data      string `json:"data"`
	Language  string `json:"language"`
	Operation string `json:"operation"`
	Key       int    `json:"key"`
}
