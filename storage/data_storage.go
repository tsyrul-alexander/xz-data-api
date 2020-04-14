package storage

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/address"
	"github.com/tsyrul-alexander/xz-data-api/model/data/company"
)

type DataStorage interface {
	GetCompanies(op *ListOptions) ([]*company.BaseCompany, error)
	GetCountries(op *ListOptions) ([]*address.Country, error)
	GetCities(op *ListOptions) ([]*address.City, error)
	GetCategories(options *ListOptions) ([]*company.Category, error)
	AddCompany(c *company.Company) error
	AddCountry(country *address.Country) error
	AddCity(city *address.City) error
	AddCategory(category *company.Category) error
}

