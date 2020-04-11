package query

import (
	"github.com/tsyrul-alexander/xz-data-api/model/request/query/filter"
	"github.com/tsyrul-alexander/xz-data-api/storage"
)

type Options struct {
	RowOffset int
	RowCount int
	Condition filter.QueryCondition
}

func (o *Options) GetDataOptions() *storage.ListOptions {
	return &storage.ListOptions{
		RowOffset: o.RowOffset,
		RowCount:  o.RowCount,
		Condition: o.Condition.CreateQueryFilter(),
	}
}