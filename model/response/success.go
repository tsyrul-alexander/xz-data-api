package response

type SuccessResponse struct {
	Message string
}

func CreateSuccessResponse() *SuccessResponse {
	return &SuccessResponse{Message: "Ok"}
}
