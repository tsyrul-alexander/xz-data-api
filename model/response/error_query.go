package response

type ErrorQuery struct {
	Message string
}

func CreateQueryError(message string) *ErrorQuery {
	return &ErrorQuery{Message:message}
}
