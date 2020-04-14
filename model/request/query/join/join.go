package join

import "github.com/tsyrul-alexander/go-query-builder/core/join"

type Join struct {
	JoinTableName string `json:"joinTableName"`
	JoinTableColumnName string `json:"joinTableColumnName"`
	MainTableName string `json:"mainTableName"`
	MainTableColumnName string `json:"mainTableColumnName"`
}

func (j *Join) CreateQueryJoin() *join.TableJoin {
	return join.CreateLeftJoin(j.JoinTableName, j.JoinTableColumnName, j.MainTableName, j.MainTableColumnName)
}
