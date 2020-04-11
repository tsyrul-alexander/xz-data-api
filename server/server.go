package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tsyrul-alexander/xz-data-api/controller"
	"github.com/tsyrul-alexander/xz-data-api/storage"
	"net/http"
	"strconv"
)

type Server struct {
	Config *Config
	DataStorage storage.DataStorage
}

func Create(config *Config, dataStorage storage.DataStorage) *Server {
	return &Server{DataStorage: dataStorage, Config:config}
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
			router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
				controller.SetResponse(writer, r.Header(writer, request))
			}).Methods(r.MethodType)
		}
}

func (s *Server)GetControllers() []controller.Controller {
	return []controller.Controller{
		controller.CreateCompanyController(s.DataStorage),
	}
}
