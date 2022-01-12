package get_records

import (
	"context"
	"github.com/rozturac/cerror"
	"github.com/rozturac/go-mediator"
	"pure-restfull-api/application/common"
	"pure-restfull-api/domain/entity"
	domain_query "pure-restfull-api/domain/query"
	"pure-restfull-api/domain/repository"
	"time"
)

type GetRecordQueryHandler struct {
	repository repository.RecordRepository
}

func NewGetRecordQueryHandler(repository repository.RecordRepository) *GetRecordQueryHandler {
	return &GetRecordQueryHandler{
		repository: repository,
	}
}

func (g *GetRecordQueryHandler) Handle(ctx context.Context, command mediator.Command) (interface{}, error) {
	query, ok := command.(*GetRecordsByTimeAndCountRangeQuery)
	if !ok {
		return nil, common.UnExpectedCommand(query)
	}

	var (
		records      []*entity.Record
		startDate    time.Time
		endDate      time.Time
		location     *time.Location
		err          error
	)

	locationName := ctx.Value(common.Location).(string)
	if location, err = time.LoadLocation(locationName); err != nil {
		return nil, cerror.New(cerror.ApplicationError, "Can not fetch the location for 'Europe/Istanbul'").With(err)
	}

	if startDate, err = time.ParseInLocation("2006-01-02", query.StartDate, location); err != nil {
		return nil, cerror.InvalidCastError(query.StartDate, startDate).With(err)
	}

	if endDate, err = time.ParseInLocation("2006-01-02", query.EndDate, location); err != nil {
		return nil, cerror.InvalidCastError(query.EndDate, endDate).With(err)
	}

	if query.MinCount < 0 {
		return nil, common.InvalidValueError("MinCount must be greater than or equal 0")
	}

	if query.MaxCount <= query.MinCount {
		return nil, common.InvalidValueError("MaxCount value must be greater than MinCount value")
	}

	records, err = g.repository.GetRecordsByFilter(&domain_query.GetRecordsByTimeAndCountRangeQuery{
		StartDate: startDate.UTC(),
		EndDate:   endDate.UTC(),
		MaxCount:  query.MaxCount,
		MinCount:  query.MinCount,
	})

	return records, err
}