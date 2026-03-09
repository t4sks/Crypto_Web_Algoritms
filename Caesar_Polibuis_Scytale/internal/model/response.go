package model

type SuccessResponse struct {
	Result    string `json:"result"`
	RequestId string `json:"request_id,omitempty"`
}

type ErrorResponse struct {
	Error     string `json:"error"`
	RequestId string `json:"request_id,omitempty"`
}
