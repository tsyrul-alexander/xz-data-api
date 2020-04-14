package controller

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/company"
	"github.com/tsyrul-alexander/xz-data-api/model/request/query"
	"github.com/tsyrul-alexander/xz-data-api/model/response"
	"github.com/tsyrul-alexander/xz-data-api/storage"
	"net/http"
)

type Category struct {
	DataStorage storage.DataStorage
}

func (c *Category)GetRoutes() (string, []*RouteRule)  {
	return "category", []*RouteRule{
		{
			MethodName: "add",
			MethodType: http.MethodPost,
			Handler:     c.AddHandler,
		},{
			MethodName: "list",
			MethodType: http.MethodPost,
			Handler:     c.GetCategoryListHandler,
		},
	}
}


func CreateCategoryController(dataStorage storage.DataStorage) *Category {
	return &Category{DataStorage:dataStorage}
}

func (c *Category) AddHandler(_ http.ResponseWriter, r *http.Request) interface{} {
	var ob = &company.Category{}
	if err := decodeJsonBody(r, ob); err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	var err = c.add(ob)
	if err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	return response.CreateSuccessResponse()
}

func (c *Category) GetCategoryListHandler(w http.ResponseWriter, r *http.Request) interface{} {
	var o = &query.Options{}
	if err := decodeJsonBody(r, o); err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	var list ,err = c.getCategoryList(o)
	if err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	return list
}

func (c *Category) add(ob *company.Category) error {
	return c.DataStorage.AddCategory(ob)
}

func (c *Category) getCategoryList(o *query.Options) ([]*company.Category, error) {
	return c.DataStorage.GetCategories(o.GetDataOptions())
}
