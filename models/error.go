package models

type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ReceivedDataError represents get data from storage or other api
type ReceivedDataErr struct {
	Message    string
	StatusCode int
}

// Error realised error type ReceivedDataErr interface
func (r ReceivedDataErr) Error() string {
	return r.Message
}
