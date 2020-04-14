package response

type ErrorResponse struct {
	Message string
}

func CreateErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Message: message}
}
