package filter

import "github.com/tsyrul-alexander/go-query-builder/core/condition"

type QueryCondition struct {
	BinaryQueryCondition *BinaryQueryCondition       `json:"binary"`
	ColumnQueryCondition *ColumnQueryCondition       `json:"column"`
	ParameterQueryCondition *ParameterQueryCondition `json:"parameter"`
	GroupQueryCondition *GroupQueryCondition         `json:"group"`
}

func (c *QueryCondition)CreateQueryFilter() condition.QueryCondition {
	if c.GroupQueryCondition != nil {
		return c.GroupQueryCondition.CreateQueryFilter()
	}
	if c.ParameterQueryCondition != nil {
		return c.ParameterQueryCondition.CreateQueryFilter()
	}
	if c.ColumnQueryCondition != nil {
		return c.ColumnQueryCondition.CreateQueryFilter()
	}
	if c.BinaryQueryCondition != nil {
		return c.BinaryQueryCondition.CreateQueryFilter()
	}
	return nil
}
