package pq

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
	"github.com/tsyrul-alexander/go-query-builder/pq/builder"
	"github.com/tsyrul-alexander/go-query-builder/query"
	"github.com/tsyrul-alexander/xz-data-api/model/data/address"
	"github.com/tsyrul-alexander/xz-data-api/model/data/base"
	"github.com/tsyrul-alexander/xz-data-api/model/data/company"
	"github.com/tsyrul-alexander/xz-data-api/model/data/image"
	"github.com/tsyrul-alexander/xz-data-api/storage"
	"strings"
)

type DataStorage struct {
	Config Config
}

const providerName = "postgres"

func Create(config Config) *DataStorage {
	return &DataStorage{Config:config}
}

func (ds *DataStorage) GetCompanies(op *storage.ListOptions) ([]*company.BaseCompany, error) {
	var s = ds.getListLookupSelect("Company", op)
	s.AddTableColumn("Category", "Id")
	s.AddTableColumn("Category", "Name")
	s.AddTableColumn("Address", "Id")
	s.AddTableColumn("Address", "Street")
	s.AddTableColumn("Address", "Building")
	s.AddTableColumn("Country", "Name")
	s.AddTableColumn("City", "Name")
	s.AddTableColumn("Image", "Id")
	s.AddTableColumn("Image", "Url")
	s.AddLeftJoin("Image", "Id", "Company", "IconId")
	s.AddLeftJoin("Address", "Id", "Company", "AddressId")
	s.AddLeftJoin("Category", "Id", "Company", "CategoryId")
	s.AddLeftJoin("Country", "Id", "Address", "CountryId")
	s.AddLeftJoin("City", "Id", "Address", "CityId")
	var companies []*company.BaseCompany
	var err = ds.getRows(s, func(rows []query.Row) {
		for _, r := range rows {
			companies = append(companies, &company.BaseCompany{
				Lookup: base.CreateLookup(r.GetUuidValue("CompanyId"), r.GetStringValue("CompanyName")),
				Category:&company.Category{
					Lookup: base.CreateLookup(r.GetUuidValue("CategoryId"), r.GetStringValue("CategoryName")),
				},
				Address: &address.BaseAddress{
					Lookup: base.CreateLookup(r.GetUuidValue("AddressId"), strings.Join([]string{
						r.GetStringValue("CountryName"),
						r.GetStringValue("CityName"),
						r.GetStringValue("AddressStreet"),
						r.GetStringValue("AddressBuilding"),
					}, ",")),
				},
				Icon: &image.Image{
					Object: base.CreateObject(r.GetUuidValue("ImageId")),
					Url: r.GetStringValue("ImageUrl"),
				},
			})
		}
	})
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (ds *DataStorage) GetCategories(op *storage.ListOptions) ([]*company.Category, error) {
	var s = ds.getListLookupSelect("Category", op)
	var list []*company.Category
	var err = ds.getRows(s, func(rows []query.Row) {
		for _, r := range rows {
			list = append(list, &company.Category{
				Lookup: ds.getLookupValue(&r, "Category"),
			})
		}
	})
	return list, err
}

func (ds *DataStorage) GetCountries(op *storage.ListOptions) ([]*address.Country, error) {
	var s = ds.getListLookupSelect("Country", op)
	var list []*address.Country
	var err = ds.getRows(s, func(rows []query.Row) {
		for _, r := range rows {
			list = append(list, &address.Country{
				Lookup: ds.getLookupValue(&r, "Country"),
			})
		}
	})
	return list, err
}

func (ds *DataStorage) GetCities(op *storage.ListOptions) ([]*address.City, error) {
	var s = ds.getListLookupSelect("City", op)
	var list []*address.City
	var err = ds.getRows(s, func(rows []query.Row) {
		for _, r := range rows {
			list = append(list, &address.City{
				Lookup: ds.getLookupValue(&r, "City"),
			})
		}
	})
	return list, err
}

func (ds *DataStorage) AddCompany(c *company.Company) error {
	var db, err = ds.getDbConnect()
	if err != nil {
		return err
	}
	var queries = []query.Transaction{
		ds.getAddressInsert(c.Address),
		ds.getImageInsert(c.Icon),
		ds.getCompanyInsert(c),
	}
	if c.Images != nil {
		for _, im := range c.Images {
			queries = append(queries, ds.getImageInsert(im))
			queries = append(queries, ds.getCompanyImageInsert(c.Id, im.Id))
		}
	}
	return query.ExecuteQueries(queries, db)
}

func (ds *DataStorage) AddCountry(c *address.Country) error {
	var i = ds.getLookupInsert("Country", c.Lookup)
	return ds.executeInsert(i)
}

func (ds *DataStorage) AddCity(city *address.City) error {
	var i = ds.getLookupInsert("City", city.Lookup)
	return ds.executeInsert(i)
}

func (ds *DataStorage) AddCategory(category *company.Category) error  {
	var i = ds.getLookupInsert("Category", category.Lookup)
	return ds.executeInsert(i)
}

func (ds *DataStorage) getImageInsert(image *image.Image) *query.Insert {
	var columnValues = column.ValueList{}
	columnValues["Id"] = parameter.CreateGuidParameter(image.Id)
	columnValues["Url"] = parameter.CreateStringParameter(image.Url)
	return builder.CreateInsert("Image", &columnValues)
}

func (ds *DataStorage) getCompanyInsert(company *company.Company) *query.Insert {
	var columnValues = column.ValueList{}
	columnValues["Id"] = parameter.CreateGuidParameter(company.Id)
	columnValues["Name"] = parameter.CreateStringParameter(company.Name)
	columnValues["AddressId"] = parameter.CreateGuidParameter(company.Address.Id)
	columnValues["OwnerId"] = parameter.CreateGuidParameter(company.Owner.Id)
	columnValues["IconId"] = parameter.CreateGuidParameter(company.Icon.Id)
	return builder.CreateInsert("Company", &columnValues)
}

func (ds *DataStorage) getCompanyImageInsert(companyId uuid.UUID, imageId uuid.UUID) *query.Insert {
	var columnValues = column.ValueList{}
	columnValues["ImageId"] = parameter.CreateGuidParameter(imageId)
	columnValues["CompanyId"] = parameter.CreateGuidParameter(companyId)
	return builder.CreateInsert("CompanyImage", &columnValues)
}

func (ds *DataStorage) getAddressInsert(address *address.Address) *query.Insert {
	var columnValues = column.ValueList{}
	columnValues["Id"] = parameter.CreateGuidParameter(address.Id)
	columnValues["CountryId"] = parameter.CreateGuidParameter(address.Country.Id)
	columnValues["CityId"] = parameter.CreateGuidParameter(address.City.Id)
	columnValues["Building"] = parameter.CreateStringParameter(address.Building)
	columnValues["Street"] = parameter.CreateStringParameter(address.Street)
	return builder.CreateInsert("Address", &columnValues)
}