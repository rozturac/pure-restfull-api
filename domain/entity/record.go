package entity

import "time"

type Record struct {
	key        string    `json:"key" bson:"key"`
	value      string    `json:"value" bson:"value"`
	createdAt  time.Time `json:"createdAt" bson:"createdAt"`
	counts     []int     `json:"counts" bson:"counts"`
	totalCount int       `json:"totalCount" bson:"totalCount"`
}

func CreateRecord(key, value string, counts []int) *Record {
	return &Record{
		key:        key,
		value:      value,
		createdAt:  time.Now().UTC(),
		counts:     counts,
		totalCount: sum(counts),
	}
}

func sum(counts []int) int {
	var sum int
	for _, v := range counts {
		sum += v
	}
	return sum
}
