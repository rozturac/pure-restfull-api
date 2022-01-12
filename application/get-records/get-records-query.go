package get_records

type GetRecordsByTimeAndCountRangeQuery struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	MinCount  int    `json:"min_count"`
	MaxCount  int    `json:"max_count"`
}