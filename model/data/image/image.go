package image

import "github.com/tsyrul-alexander/xz-data-api/model/data/base"

type Image struct {
	*base.Object
	Url string `json:"url"`
}
