package controller

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/company"
	"github.com/tsyrul-alexander/xz-data-api/model/request/query"
	"github.com/tsyrul-alexander/xz-data-api/model/response"
	"github.com/tsyrul-alexander/xz-data-api/storage"
	"net/http"
)

type Company struct {
	DataStorage storage.DataStorage
}

func (c *Company)GetRoutes() (string, []RouteRule)  {
	return "company", []RouteRule{
		{
			MethodName: "list",
			MethodType: http.MethodPost,
			Header:     c.GetCompaniesHandler,
		},
	}
}

func CreateCompanyController(dataStorage storage.DataStorage) *Company {
	return &Company{DataStorage:dataStorage}
}

func (c *Company)GetCompaniesHandler(w http.ResponseWriter, r *http.Request) interface{} {
	var op = query.Options{}
	if err := decodeJsonBody(r, &op); err != nil {
		return response.CreateQueryError(err.Error())//todo
	}
	var companies, err = c.getCompanies(&op)
	if err != nil {
		return response.CreateQueryError(err.Error())
	}
	return companies
}

func (c *Company)getCompanies(options *query.Options) ([]*company.BaseCompany, error) {
	return c.DataStorage.GetCompanies(options.GetDataOptions())
}
