package model

type SuccessResponse struct {
	Result      string `json:"result"`
	RequestId   string `json:"request_id,omitempty"`
	CardanoCode string `json:"cardano_code,omitempty"`
}

type ErrorResponse struct {
	Error     string `json:"error"`
	RequestId string `json:"request_id,omitempty"`
}
