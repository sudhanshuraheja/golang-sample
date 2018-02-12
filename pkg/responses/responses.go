package responses

import (
	"encoding/json"
	"net/http"
)

// WriteJSON - to write json to the output response
func WriteJSON(w http.ResponseWriter, status int, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(i)

	// var b bytes.Buffer
	// json.NewEncoder(&b).Encode(i)
	// fmt.Printf(b.String())
}

// Response - structure to define the response format
type Response struct {
	Data    interface{}     `json:"data,omitempty"`
	Success string          `json:"success,omitempty"`
	Errors  []ErrorResponse `json:"errors,omitempty"`
}

// ErrorResponse - structure to define the error response
type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// NewErrorResponse - respond with an error
func NewErrorResponse(code string, message string) *Response {
	return &Response{
		Success: "false",
		Errors: []ErrorResponse{
			ErrorResponse{
				Code:    code,
				Message: message,
			},
		},
	}
}
