package filter

import "github.com/tsyrul-alexander/go-query-builder/core/condition"

type ComparisonType int

const (
	Equal = ComparisonType(0)
	NotEqual = ComparisonType(1)
	In = ComparisonType(2)
)

func (ct ComparisonType)GetQueryComparisonType() condition.ComparisonType {
	return condition.ComparisonType(int(ct))
}
