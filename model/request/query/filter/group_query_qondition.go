package filter

import "github.com/tsyrul-alexander/go-query-builder/core/condition"

type GroupQueryCondition struct {
	LogicalOperation LogicalOperation  `json:"operation"`
	QueryConditions  []*QueryCondition `json:"conditions"`
}

func (gc GroupQueryCondition) CreateQueryFilter() *condition.GroupQueryCondition {
	var conditions []condition.QueryCondition
	for _, c := range gc.QueryConditions {
		conditions = append(conditions, c.CreateQueryFilter())
	}
	return &condition.GroupQueryCondition {
		LogicalOperation: gc.LogicalOperation.GetQueryLogicalOperation(),
		QueryConditions: conditions,
	}
}
