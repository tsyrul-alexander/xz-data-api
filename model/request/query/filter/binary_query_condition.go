package filter

import "github.com/tsyrul-alexander/go-query-builder/core/condition"

type BinaryQueryCondition struct {
	ComparisonType ComparisonType  `json:"comparisonType"`
	LeftCondition  *QueryCondition `json:"left"`
	RightCondition *QueryCondition `json:"right"`
}

func (bc *BinaryQueryCondition)CreateQueryFilter() condition.QueryCondition {
	return &condition.BinaryQueryCondition{
		ComparisonType:bc.ComparisonType.GetQueryComparisonType(),
		LeftCondition: bc.LeftCondition.CreateQueryFilter(),
		RightCondition: bc.RightCondition.CreateQueryFilter(),
	}
}