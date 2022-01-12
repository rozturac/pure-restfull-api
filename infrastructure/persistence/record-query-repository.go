package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pure-restfull-api/domain/query"
	"pure-restfull-api/domain/repository"
)

type recordQueryRepository struct {
	collection *mongo.Collection
}

func NewRecordRepository(collection *mongo.Collection) repository.RecordQueryRepository {
	return &recordQueryRepository{
		collection: collection,
	}
}

func (r *recordQueryRepository) GetRecordsByFilter(ctx context.Context, filter *query.GetRecordsByTimeAndCountRangeQuery) (*query.GetRecordsByTimeAndCountRangeResult, error) {
	pipeline := mongo.Pipeline{
		{
			{"$match", bson.D{
				{"createdAt", bson.D{
					{"$gte", filter.StartDate},
					{"$lte", filter.EndDate},
				}},
			}},
		},
		{
			{"$project", bson.D{
				{"key", 1},
				{"createdAt", 1},
				{"totalCount", bson.D{{"$sum", "$counts"}}},
			}},
		},
		{
			{"$match", bson.D{
				{"totalCount", bson.D{
					{"$lte", filter.MaxCount},
					{"$gte", filter.MinCount},
				}},
			}},
		},
	}

	cur, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	result := &query.GetRecordsByTimeAndCountRangeResult{
		Code:    0,
		Message: "Success",
	}

	for cur.Next(ctx) {
		var item *query.GetRecordsByTimeAndCountRangeItem
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		result.Records = append(result.Records, item)
	}

	return result, err
}
