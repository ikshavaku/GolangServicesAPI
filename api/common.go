package api

type APIError struct {
	Status   int
	Response ErrorResponse
}

type ErrorResponse struct {
	Details string `json:"message"`
	Error   string `json:"error"`
}
