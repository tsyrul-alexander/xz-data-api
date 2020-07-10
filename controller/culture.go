package controller

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/culture"
	"github.com/tsyrul-alexander/xz-data-api/model/request/query"
	"github.com/tsyrul-alexander/xz-data-api/model/response"
	"github.com/tsyrul-alexander/xz-data-api/storage"
	"github.com/tsyrul-alexander/xz-identity-api/model"
	"net/http"
)

type Culture struct {
	DataStorage storage.DataStorage
}

func (c *Culture)GetRoutes() (string, []*RouteRule)  {
	return "culture", []*RouteRule{
		{
			MethodName: "list",
			MethodType: http.MethodPost,
			Handler:     c.GetCultureListHandler,
		},
		{
			MethodName: "add",
			MethodType: http.MethodPost,
			Handler:     c.AddCultureHandler,
			Roles: []model.UserRole{
				model.UserRoleAdmin,
			},
		},
	}
}

func CreateCultureController(dataStorage storage.DataStorage) *Culture {
	return &Culture{DataStorage:dataStorage}
}

func (c *Culture) GetCultureListHandler(w http.ResponseWriter, r *http.Request) interface{} {
	var op = query.Options{}
	if err := decodeJsonBody(r, &op); err != nil {
		return response.CreateErrorResponse(err.Error()) //todo
	}
	var companies, err = c.getCultures(&op)
	if err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	return companies
}

func (c *Culture) AddCultureHandler(w http.ResponseWriter, r *http.Request) interface{} {
	return nil //todo
}

func (c *Culture) getCultures(o *query.Options) ([]*culture.Culture, error) {
	return c.DataStorage.GetCultures(o.GetDataOptions())
}
