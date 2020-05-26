package storage

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/address"
	"github.com/tsyrul-alexander/xz-data-api/model/data/company"
	"github.com/tsyrul-alexander/xz-data-api/model/data/company/category"
	"github.com/tsyrul-alexander/xz-data-api/model/data/culture"
)

type DataStorage interface {
	GetCompanies(op *ListOptions) ([]*company.BaseCompany, error)
	GetCountries(op *ListOptions) ([]*address.Country, error)
	GetCities(op *ListOptions) ([]*address.City, error)
	GetCategories(options *ListOptions) ([]*category.Category, error)
	GetCultures(options *ListOptions) ([]*culture.Culture, error)
	AddCompany(c *company.Company) error
	AddCountry(country *address.Country) error
	AddCity(city *address.City) error
	AddCategory(category *category.AddCategory) error
}

