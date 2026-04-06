package model

type Request struct {
	Algorithm string `json:"algorithm"`
	Data      string `json:"data"`
	Language  string `json:"language,omitempty"`
	Operation string `json:"operation"`
	Key       int    `json:"key,omitempty"`
	Code      string `json:"code,omitempty"`
	KeyString string `json:"keyString,omitempty"`
}
