package get_records

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rozturac/cerror"
	"pure-restfull-api/application/common"
	"pure-restfull-api/mock"
	"testing"
	"time"
)

func GetContext(locationName string) context.Context {
	ctx := context.WithValue(context.Background(), common.CorrelationId, uuid.NewString())
	return context.WithValue(ctx, common.Location, locationName)
}

func TestGetRecordQueryHandler_HandleWhenSuccess(t *testing.T) {
	locationName := "Europe/Istanbul"
	ctx := GetContext(locationName)
	query := &GetRecordsByTimeAndCountRangeQuery{
		StartDate: "2022-01-01",
		EndDate:   "2022-01-02",
		MinCount:  1,
		MaxCount:  2,
	}

	repository := mock.NewRecordRepository()
	handler := NewGetRecordQueryHandler(repository)

	if _, err := handler.Handle(ctx, query); err != nil {
		t.Errorf("Expected error: nill received: %v", err)
	}
}

func TestGetRecordQueryHandler_HandleWhenLocationNameIsIncorrect(t *testing.T) {
	locationName := "asd"
	ctx := GetContext(locationName)
	query := &GetRecordsByTimeAndCountRangeQuery{
		StartDate: "2022-01-01",
		EndDate:   "2022-01-02",
		MinCount:  1,
		MaxCount:  2,
	}
	expectedError := cerror.New(cerror.ApplicationError, fmt.Sprintf("Can not fetch the location for '%s'", locationName))
	repository := mock.NewRecordRepository()
	handler := NewGetRecordQueryHandler(repository)

	if _, err := handler.Handle(ctx, query); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}

func TestGetRecordQueryHandler_HandleWhenDateFormatIsIncorrect(t *testing.T) {
	ctx := GetContext("Europe/Istanbul")
	query := &GetRecordsByTimeAndCountRangeQuery{
		StartDate: "01-01-2022",
		EndDate:   "01-01-2022",
		MinCount:  1,
		MaxCount:  2,
	}
	expectedError := cerror.InvalidCastError(query.StartDate, time.Time{})
	repository := mock.NewRecordRepository()
	handler := NewGetRecordQueryHandler(repository)

	if _, err := handler.Handle(ctx, query); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}

func TestGetRecordQueryHandler_HandleWhenEndDateGreaterThanStartDate(t *testing.T) {
	locationName := "Europe/Istanbul"
	ctx := GetContext(locationName)
	query := &GetRecordsByTimeAndCountRangeQuery{
		StartDate: "2022-01-01",
		EndDate:   "2022-01-01",
		MinCount:  1,
		MaxCount:  2,
	}
	expectedError := common.InvalidValueError("EndDate must be greater than StartDate")
	repository := mock.NewRecordRepository()
	handler := NewGetRecordQueryHandler(repository)

	if _, err := handler.Handle(ctx, query); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}

func TestGetRecordQueryHandler_HandleWhenMinValueLessThanZero(t *testing.T) {
	locationName := "Europe/Istanbul"
	ctx := GetContext(locationName)
	query := &GetRecordsByTimeAndCountRangeQuery{
		StartDate: "2022-01-01",
		EndDate:   "2022-01-02",
		MinCount:  -1,
		MaxCount:  2,
	}
	expectedError := common.InvalidValueError("MinCount must be greater than or equal 0")
	repository := mock.NewRecordRepository()
	handler := NewGetRecordQueryHandler(repository)

	if _, err := handler.Handle(ctx, query); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}

func TestGetRecordQueryHandler_HandleWhenMaxValueLessThanMinValue(t *testing.T) {
	locationName := "Europe/Istanbul"
	ctx := GetContext(locationName)
	query := &GetRecordsByTimeAndCountRangeQuery{
		StartDate: "2022-01-01",
		EndDate:   "2022-01-02",
		MinCount:  2,
		MaxCount:  1,
	}
	expectedError := common.InvalidValueError("MaxCount value must be greater than MinCount value")
	repository := mock.NewRecordRepository()
	handler := NewGetRecordQueryHandler(repository)

	if _, err := handler.Handle(ctx, query); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}

func TestGetConfigQueryHandler_HandleWhenQueryUnCorrect(t *testing.T) {
	ctx := context.Background()
	type UnexpectedQuery struct {
	}
	query := &UnexpectedQuery{}
	expectedError := common.UnExpectedCommand("mediator.Command", query)
	repository := mock.NewRecordRepository()
	handler := NewGetRecordQueryHandler(repository)

	if _, err := handler.Handle(ctx, query); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}
