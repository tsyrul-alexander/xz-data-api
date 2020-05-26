package controller

type ErrorCode int

type Error struct {
	ErrorCode ErrorCode
	Message string
}

const (
	RequestError = ErrorCode(201)
)