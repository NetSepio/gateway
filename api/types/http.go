package types

type ApiResponse struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Error      string      `json:"error,omitempty"`
	Message    string      `json:"message,omitempty"`
	Payload    interface{} `json:"payload,omitempty"`
}
