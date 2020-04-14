package server

import (
	"errors"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tsyrul-alexander/xz-data-api/controller"
	"github.com/tsyrul-alexander/xz-data-api/core/identity"
	"github.com/tsyrul-alexander/xz-data-api/model/response"
	"github.com/tsyrul-alexander/xz-data-api/storage"
	"net/http"
	"strconv"
)

type Server struct {
	Config *Config
	DataStorage storage.DataStorage
	IdentityService *identity.Service
}

func Create(config *Config, dataStorage storage.DataStorage, identityService *identity.Service) *Server {
	return &Server{DataStorage: dataStorage, Config:config, IdentityService:identityService}
}

func (s *Server)Start() error {
	var router = s.UseRouting()
	var serverAddress = s.Config.Ip + ":" + strconv.Itoa(s.Config.Port)
	return http.ListenAndServe(serverAddress, router)
}

//UseRouting ...
func (s *Server)UseRouting() *mux.Router {
	var router = mux.NewRouter()
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.UseControllersRouting(router)
	return router
}

func (s *Server)UseControllersRouting(router *mux.Router) {
	var controllers = s.GetControllers()
	for _, c := range controllers {
		s.UseControllerRouting(router, c)
	}
}

func (s *Server)UseControllerRouting(router *mux.Router, c controller.Controller) {
		var controllerName, routes = c.GetRoutes()
		for _, r := range routes {
			var path = "/" + controllerName + "/" + r.MethodName
			router.HandleFunc(path, s.GetHandleFunc(r)).Methods(r.MethodType)
		}
}

func (s *Server) GetHandleFunc(rule *controller.RouteRule) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var err = s.ValidateUserRole(request, rule)
		if err != nil {
			controller.SetResponse(writer, response.CreateErrorResponse(err.Error()))
			return
		}
		controller.SetResponse(writer, rule.Handler(writer, request))
	}
}

func (s *Server) ValidateUserRole(request *http.Request, rule *controller.RouteRule) error {
	if rule.Roles != nil && len(rule.Roles) > 0 {
		var token = s.GetToken(request)
		if token == "" {
			return errors.New("not authorized")
		}
		var success, err = s.IdentityService.GetUserInRoles(rule.Roles, token)
		if err == nil && !success {
			return errors.New("access denied")
		}
		return err
	}
	return nil
}

func (s *Server) GetToken(request *http.Request) string {
	return request.Header.Get("Authorization")
}

func (s *Server) GetControllers() []controller.Controller {
	return []controller.Controller{
		controller.CreateCompanyController(s.DataStorage),
		controller.CreateAddressController(s.DataStorage),
		controller.CreateCategoryController(s.DataStorage),
	}
}
