package controller

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/address"
	"github.com/tsyrul-alexander/xz-data-api/model/response"
	"github.com/tsyrul-alexander/xz-data-api/storage"
	"net/http"
)

type Address struct {
	DataStorage storage.DataStorage
}

func CreateAddressController(dataStorage storage.DataStorage) *Address {
	return &Address{DataStorage:dataStorage}
}

func (a *Address)GetRoutes() (string, []*RouteRule)  {
	return "address", []*RouteRule{
		{
			MethodName: "country/add",
			MethodType: http.MethodPost,
			Handler:     a.AddCountryHandler,
		},{
			MethodName: "city/add",
			MethodType: http.MethodPost,
			Handler:     a.AddCityHandler,
		},
	}
}

func (a *Address) AddCountryHandler(_ http.ResponseWriter, r *http.Request) interface{} {
	var c = &address.Country{}
	if err := decodeJsonBody(r, c); err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	if err :=  a.addCountry(c); err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	return response.CreateSuccessResponse()
}

func (a *Address)addCountry(country *address.Country) error {
	return a.DataStorage.AddCountry(country)
}

func (a *Address) AddCityHandler(_ http.ResponseWriter, r *http.Request) interface{} {
	var c = &address.City{}
	if err := decodeJsonBody(r, c); err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	if err :=  a.addCity(c); err != nil {
		return response.CreateErrorResponse(err.Error())
	}
	return response.CreateSuccessResponse()
}

func (a *Address)addCity(city *address.City) error {
	return a.DataStorage.AddCity(city)
}
