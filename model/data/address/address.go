package address

import "github.com/tsyrul-alexander/xz-data-api/model/data/base"

type Address struct {
	*base.Object
	Country     *Country `json:"country"`
	City        *City    `json:"city"`
	Street      string  `json:"street"`
	Building    string  `json:"building"`
	CoordinateX float32 `json:"coordinateX"`
	CoordinateY float32 `json:"coordinateY"`
}