package helpers

type EncryptRequest struct {
	Text string `json:"text"`
	Key string `json:"Key"`
}

type EncryptResponse struct {
	Message string `json:"message"`
	Err string `json:"error"`
}

type DecryptRequest struct {
	Message string `json:"message"`
	Key string `json:"Key"`
}

type DecryptResponse struct {
	Text string `json:"text"`
	Err string `json:"error"`
}

