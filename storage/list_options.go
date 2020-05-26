package storage

import (
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
	"github.com/tsyrul-alexander/go-query-builder/core/join"
	"github.com/tsyrul-alexander/xz-data-api/model/data/culture"
)

type ListOptions struct {
	RowOffset int
	RowCount int
	Condition condition.QueryCondition
	Join *join.List
	CultureId culture.CultureId
}