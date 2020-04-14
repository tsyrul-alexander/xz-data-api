package query

import (
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
	queryJoin "github.com/tsyrul-alexander/go-query-builder/core/join"
	"github.com/tsyrul-alexander/xz-data-api/model/request/query/filter"
	"github.com/tsyrul-alexander/xz-data-api/model/request/query/join"
	"github.com/tsyrul-alexander/xz-data-api/storage"
)

type Options struct {
	RowOffset int
	RowCount int
	Condition *filter.QueryCondition
	Joins *join.Joins
}

func (o *Options) GetDataOptions() *storage.ListOptions {
	return &storage.ListOptions{
		RowOffset: o.RowOffset,
		RowCount:  o.RowCount,
		Condition: o.GetCondition(),
		Join:     o.GetJoin(),
	}
}

func (o *Options) GetJoin() *queryJoin.List {
	if o.Joins != nil {
		return o.Joins.CreateJoinList()
	}
	return nil
}

func (o *Options) GetCondition() condition.QueryCondition {
	if o.Condition != nil {
		return o.Condition.CreateQueryFilter()
	}
	return nil
}