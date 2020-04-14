package controller

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/company"
	"github.com/tsyrul-alexander/xz-data-api/model/request/query"
	"github.com/tsyrul-alexander/xz-data-api/model/response"
	"github.com/tsyrul-alexander/xz-data-api/storage"
	"github.com/tsyrul-alexander/xz-identity-api/model"
	"net/http"
)

type Company struct {
	DataStorage storage.DataStorage
}

func (c *Company)GetRoutes() (string, []*RouteRule)  {
	return "company", []*RouteRule{
		{
			MethodName: "list",
			MethodType: http.MethodPost,
			Handler:     c.GetCompanyListHandler,
			Roles: []model.UserRole{
				model.UserRoleClient,
				model.UserRoleAccount,
				model.UserRoleAdmin,
			},
		},
		{
			MethodName: "add",
			MethodType: http.MethodPost,
			Handler:     c.AddCompanyHandler,
		},
	}
}

func CreateCompanyController(dataStorage storage.DataStorage) *Company {
	return &Company{DataStorage:dataStorage}
}

func (c *Company) AddCompanyHandler(_ http.ResponseWriter, r *http.Request) interface{} {
	var ob = &company.Company{}
	if err := decodeJsonBody(r, &ob); err != nil {
		return response.CreateErrorResponse(err.Error()) //todo
	}
	var err = c.addCompany(ob)
	if err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	return response.CreateSuccessResponse()
}

func (c *Company) GetCompanyListHandler(_ http.ResponseWriter, r *http.Request) interface{} {
	var op = query.Options{}
	if err := decodeJsonBody(r, &op); err != nil {
		return response.CreateErrorResponse(err.Error()) //todo
	}
	var companies, err = c.getCompanies(&op)
	if err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	return companies
}

func (c *Company)getCompanies(options *query.Options) ([]*company.BaseCompany, error) {
	return c.DataStorage.GetCompanies(options.GetDataOptions())
}

func (c *Company) addCompany(ob *company.Company) error {
	return c.DataStorage.AddCompany(ob)
}

