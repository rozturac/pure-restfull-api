package query

import "time"

type GetRecordsByTimeAndCountRangeQuery struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	MinCount  int       `json:"min_count"`
	MaxCount  int       `json:"max_count"`
}

type GetRecordsByTimeAndCountRangeResult struct {
	Code    int                                  `json:"code"`
	Message string                               `json:"msg"`
	Records []*GetRecordsByTimeAndCountRangeItem `json:"records"`
}

type GetRecordsByTimeAndCountRangeItem struct {
	Key        string    `json:"key" bson:"key"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	TotalCount int       `json:"totalCount" bson:"totalCount"`
}
