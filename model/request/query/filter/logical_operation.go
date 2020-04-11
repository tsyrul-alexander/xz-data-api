package filter

import "github.com/tsyrul-alexander/go-query-builder/core/condition"

type LogicalOperation int

const (
	And = LogicalOperation(0)
	Or = LogicalOperation(1)
)

func (lo *LogicalOperation) GetQueryLogicalOperation() condition.LogicalOperation {
	if *lo == Or {
		return condition.Or
	}
	return condition.And
}