package filter

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
)

type ParameterQueryCondition struct {
	StringValue *string `json:"stringValue"`
	InvValue *int `json:"intValue"`
	GuidValue *uuid.UUID `json:"guidValue"`
}

func (pc *ParameterQueryCondition)CreateQueryFilter() condition.QueryCondition {
	return &condition.ParameterQueryCondition{
		Value: pc.GetQueryParameter(),
	}
}
func (pc *ParameterQueryCondition)GetQueryParameter() parameter.QueryParameter  {
	if pc.StringValue != nil {
		return parameter.CreateStringParameter(*pc.StringValue)
	}
	if pc.InvValue != nil {
		return parameter.CreateIntParameter(*pc.InvValue)
	}
	if pc.GuidValue != nil {
		return parameter.CreateGuidParameter(*pc.GuidValue)
	}
	return nil
}
