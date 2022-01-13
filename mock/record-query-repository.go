package mock

import (
	"context"
	"pure-restfull-api/domain/query"
	"pure-restfull-api/domain/repository"
	"time"
)

type recordQueryRepository struct {
}

func NewRecordRepository() repository.RecordQueryRepository {
	return &recordQueryRepository{}
}

func (r *recordQueryRepository) GetRecordsByFilter(ctx context.Context, filter *query.GetRecordsByTimeAndCountRangeQuery) (*query.GetRecordsByTimeAndCountRangeResult, error) {
	result := &query.GetRecordsByTimeAndCountRangeResult{
		Code:    0,
		Message: "Success",
	}

	result.Records = append(result.Records, &query.GetRecordsByTimeAndCountRangeItem{
		Key:        "mockKey",
		CreatedAt:  time.Now(),
		TotalCount: 1071,
	})

	return result, nil
}
