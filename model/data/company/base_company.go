package company

import (
	"github.com/tsyrul-alexander/xz-data-api/model/data/address"
	"github.com/tsyrul-alexander/xz-data-api/model/data/base"
	"github.com/tsyrul-alexander/xz-data-api/model/data/company/category"
	"github.com/tsyrul-alexander/xz-data-api/model/data/image"
)

type BaseCompany struct {
	*base.Lookup
	Category *category.Category   `json:"category"`
	Address  *address.BaseAddress `json:"address"`
	Icon     *image.Image         `json:"icon"`
}