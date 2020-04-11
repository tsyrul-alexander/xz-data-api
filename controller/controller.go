package controller

type Controller interface {
	GetRoutes() (string, []RouteRule)
}
