package storage

import "github.com/tsyrul-alexander/go-query-builder/core/condition"

type ListOptions struct {
	RowOffset int
	RowCount int
	Condition condition.QueryCondition
}