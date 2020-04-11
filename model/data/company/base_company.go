package company

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/address"
	"github.com/tsyrul-alexander/xz-data-api/model/data/base"
	"github.com/tsyrul-alexander/xz-data-api/model/data/image"
)

type BaseCompany struct {
	base.Lookup
	Address address.BaseAddress
	Icon image.Image
}
