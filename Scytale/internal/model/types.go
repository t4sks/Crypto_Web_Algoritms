// Только структуры
package model

type CipherRequest struct {
	Text      string `json:"text"`
	Key       int    `json:"key"`
	Operation string `json:"operation"`
}

type EncryptResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
