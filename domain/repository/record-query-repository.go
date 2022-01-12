package repository

import (
	"context"
	"pure-restfull-api/domain/query"
)

type RecordQueryRepository interface {
	GetRecordsByFilter(ctx context.Context, filter *query.GetRecordsByTimeAndCountRangeQuery) (*query.GetRecordsByTimeAndCountRangeResult, error)
}
