package pq

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
	"github.com/tsyrul-alexander/go-query-builder/pq/builder"
	"github.com/tsyrul-alexander/go-query-builder/query"
	"github.com/tsyrul-alexander/xz-data-api/model/data/base"
	"github.com/tsyrul-alexander/xz-data-api/model/data/culture"
	"github.com/tsyrul-alexander/xz-data-api/storage"
)

func (ds *DataStorage) getListBaseSelect(tableName string, options *storage.ListOptions) *query.Select {
	var s = builder.CreateSelect(tableName)
	s.RowCount = options.RowCount
	s.RowOffset = options.RowOffset
	if options.Condition != nil {
		s.AddCondition(options.Condition)
	}
	if options.Join != nil {
		s.Joins = options.Join
	}
	return s
}

func (ds *DataStorage) getListLookupLczSelect(tableName string, options *storage.ListOptions) *query.Select {
	var s = ds.getListBaseSelect(tableName, options)
	s.AddTableColumn(tableName, "Id")
	s.AddLocalizeColumn(tableName, "Name", options.CultureId)
	return s
}

func (ds *DataStorage) getListLookupSelect(tableName string, options *storage.ListOptions) *query.Select {
	var s = ds.getListBaseSelect(tableName, options)
	s.AddTableColumn(tableName, "Id")
	s.AddTableColumn(tableName, "Name")
	return s
}

func (ds *DataStorage) getRows(s *query.Select, f func([]query.Row)) error {
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

func (ds *DataStorage) getLookupInsert(tableName string, lookup *base.Lookup) *query.Insert {
	var columnValues = column.ValueList{}
	columnValues["Id"] = parameter.CreateGuidParameter(lookup.Id)
	columnValues["Name"] = parameter.CreateStringParameter(lookup.Name)
	return builder.CreateInsert(tableName, &columnValues)
}

func (ds *DataStorage) getLookupLczInsert(tableName string, lookup *base.LookupLcz) *[]query.Insert {
	var columnValues = column.ValueList{}
	columnValues["Id"] = parameter.CreateGuidParameter(lookup.Id)
	columnValues["Name"] = parameter.CreateStringParameter(lookup.Name.DefValue)
	var lczInserts = ds.getLczInserts(tableName, "Name", lookup.Id, lookup.Name.Values)
	var baseInsert = builder.CreateInsert(tableName, &columnValues)
	var newLczInsertsArray = append([]query.Insert{*baseInsert}, *lczInserts...)
	return &newLczInsertsArray
}

func (ds *DataStorage) getLczInserts(tableName string, lczColumnName string, recordId uuid.UUID,
	values *[]culture.Value) *[]query.Insert {
	var inserts []query.Insert
	for _, value := range *values {
		inserts = append(inserts, *ds.getLczInsert(tableName, lczColumnName, recordId, value))
	}
	return &inserts
}

func (ds *DataStorage) getLczInsert(tableName string, lczColumnName string, recordId uuid.UUID,
		value *culture.Value) *query.Insert {
	var columnValues = column.ValueList{}
	columnValues["Id"] = parameter.CreateGuidParameter(uuid.New())
	columnValues[lczColumnName] = parameter.CreateStringParameter(value.Value)
	columnValues["CultureId"] = parameter.CreateGuidParameter(uuid.UUID(value.CultureId))
	columnValues["RecordId"] = parameter.CreateGuidParameter(recordId)
	return builder.CreateInsert(tableName + query.LczTablePrefix, &columnValues)
}

func (ds *DataStorage) getDbConnect() (*sql.DB, error) {
	return sql.Open(providerName, ds.Config.ConnectionString)
}

func (ds *DataStorage) executeInsert(i *query.Insert) error {
	var db, errDb = ds.getDbConnect()
	if errDb != nil {
		return errDb
	}
	if _, errExecute := i.Execute(db); errExecute != nil {
		_ = db.Close()
		return errDb
	}
	return db.Close()
}

func (ds *DataStorage) getLookupValue(r *query.Row, tableName string) *base.Lookup {
	return base.CreateLookup(r.GetUuidValue(tableName + "Id"), r.GetStringValue(tableName + "Name"))
}
