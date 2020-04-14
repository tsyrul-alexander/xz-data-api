package join

import "github.com/tsyrul-alexander/go-query-builder/core/join"

type Joins []*Join

func (jl *Joins) GetItems() []*Join {
	return *jl
}

func (jl *Joins) CreateJoinList() *join.List {
	var l = join.List{}
	for _, j := range jl.GetItems() {
		l = append(l, *j.CreateQueryJoin())
	}
	return &l
}
