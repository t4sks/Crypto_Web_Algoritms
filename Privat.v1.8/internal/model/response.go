package model

type SuccessResponse struct {
	Result      string `json:"result"`
	RequestId   string `json:"request_id,omitempty"`
	CardanoCode string `json:"cardano_code,omitempty"`
	DownloadURL string `json:"download_url,omitempty"`
}

type ErrorResponse struct {
	Error     string `json:"error"`
	RequestId string `json:"request_id,omitempty"`
}
