package controller

import "net/http"

type RouteRule struct {
	MethodName string
	MethodType string
	Header func(w http.ResponseWriter, r *http.Request) interface{}
}