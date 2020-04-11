package filter

import (
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
	"github.com/tsyrul-alexander/xz-data-api/model/request/query/column"
)

type ColumnQueryCondition struct {
	QueryColumn *column.TableColumn `json:"table"`
}

func (cc *ColumnQueryCondition)CreateQueryFilter() condition.QueryCondition {
	return &condition.ColumnQueryCondition{
		QueryColumn: cc.QueryColumn.CreateQueryTable(),
	}
}