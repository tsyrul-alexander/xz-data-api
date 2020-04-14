package controller

import (
	"github.com/tsyrul-alexander/xz-identity-api/model"
	"net/http"
)

type RouteRule struct {
	MethodName string
	MethodType string
	Handler func(w http.ResponseWriter, r *http.Request) interface{}
	Roles []model.UserRole
}
