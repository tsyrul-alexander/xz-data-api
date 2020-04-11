package user

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/address"
	"github.com/tsyrul-alexander/xz-data-api/model/data/base"
)

type User struct {
	base.Lookup
	Address address.Address
	Roles []Role
}
