package column

import (
	"github.com/tsyrul-alexander/go-query-builder/core/column"
)

type TableColumn struct {
	TableName string `json:"tableName"`
	ColumnName string `json:"columnName"`
}

func (c *TableColumn) CreateQueryTable() column.QueryColumn {
	return &column.TableColumn{
		TableName:  c.TableName,
		ColumnName: c.ColumnName,
	}
}
