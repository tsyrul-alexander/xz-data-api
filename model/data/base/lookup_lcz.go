package base

import "github.com/tsyrul-alexander/xz-data-api/model/data/culture"

type LookupLcz struct {
	Object
	Name *culture.ValueList
}
