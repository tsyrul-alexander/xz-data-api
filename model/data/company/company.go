package company

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/address"
	"github.com/tsyrul-alexander/xz-data-api/model/data/base"
	"github.com/tsyrul-alexander/xz-data-api/model/data/image"
	"github.com/tsyrul-alexander/xz-data-api/model/data/user"
)

type Company struct {
	*base.Lookup
	Address *address.Address `json:"address"`
	Owner   *user.BaseUser   `json:"owner"`
	Icon    *image.Image     `json:"icon"`
	Images  []*image.Image
}
