package pq

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/tsyrul-alexander/go-query-builder/pq/builder"
	"github.com/tsyrul-alexander/go-query-builder/query"
	"github.com/tsyrul-alexander/xz-data-api/model/data/address"
	"github.com/tsyrul-alexander/xz-data-api/model/data/base"
	"github.com/tsyrul-alexander/xz-data-api/model/data/company"
	"github.com/tsyrul-alexander/xz-data-api/model/data/image"
	"github.com/tsyrul-alexander/xz-data-api/storage"
)

type DataStorage struct {
	Config Config
}

const providerName = "postgres"

func Create(config Config) *DataStorage {
	return &DataStorage{Config:config}
}

func (ds *DataStorage)GetCompanies(op *storage.ListOptions) ([]*company.BaseCompany, error) {
	var s = ds.getListBaseSelect("Company", op)
	s.AddTableColumn("Company", "Id")
	s.AddTableColumn("Company", "Name")
	var companies []*company.BaseCompany
	var err = ds.getRows(s, func(rows []query.Row) {
		for _, r := range rows {
			companies = append(companies, &company.BaseCompany{
				Lookup:  base.Lookup{
					Object: base.Object{
						Id: r.GetUuidValue("CompanyId"),
					},
					Name:   r.GetStringValue("CompanyName"),
				},
				Address: address.BaseAddress{},
				Icon:    image.Image{},
			})
		}
	})
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (ds *DataStorage) getListBaseSelect(tableName string, options *storage.ListOptions) *query.Select {
	var s = builder.CreateSelect(tableName)
	s.RowCount = options.RowCount
	s.RowOffset = options.RowOffset
	if options.Condition != nil {
		s.AddCondition(options.Condition)
	}
	return s
}

func (ds *DataStorage)getRows(s *query.Select, f func([]query.Row)) error {
	var connection, connectErr = ds.getDbConnect()
	if connectErr != nil {
		return connectErr
	}
	var rows, executeErr = s.Execute(connection)
	if executeErr != nil {
		return executeErr
	}
	f(*rows)
	return connection.Close()
}

func (ds *DataStorage)getDbConnect() (*sql.DB, error) {
	return sql.Open(providerName, ds.Config.ConnectionString)
}
