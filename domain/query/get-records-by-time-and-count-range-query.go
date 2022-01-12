package query

import "time"

type GetRecordsByTimeAndCountRangeQuery struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	MinCount  int       `json:"min_count"`
	MaxCount  int       `json:"max_count"`
}