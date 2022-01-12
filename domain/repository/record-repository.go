package repository

import (
	"pure-restfull-api/domain/entity"
	"pure-restfull-api/domain/query"
)

type RecordRepository interface {
	GetRecordsByFilter(filter *query.GetRecordsByTimeAndCountRangeQuery) ([]*entity.Record, error)
}
