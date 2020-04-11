package storage

import "github.com/tsyrul-alexander/xz-data-api/model/data/company"

type DataStorage interface {
	GetCompanies(op *ListOptions) ([]*company.BaseCompany, error)
}
