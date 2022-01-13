package get_config

import (
	"context"
	"fmt"
	"github.com/rozturac/cerror"
	"net/http"
	"pure-restfull-api/application/common"
	"pure-restfull-api/domain/entity"
	"pure-restfull-api/mock"
	"testing"
)

func TestGetConfigQueryHandler_HandleWhenSuccess(t *testing.T) {
	ctx := context.Background()
	query := &GetConfigQuery{
		Key: "Hello",
	}
	expectedConfig := entity.CreateConfig("Hello", "World")
	repository := mock.NewConfigRepository(mock.NewInMemoryDB())
	_ = repository.Insert(expectedConfig)
	handler := NewGetConfigQueryHandler(repository)

	result, err := handler.Handle(ctx, query)
	if err != nil {
		t.Errorf("Expected is success, received: %v", err)
	}
	if receivedConfig, _ := result.(*entity.Config); receivedConfig.GetValue() != expectedConfig.GetValue() {
		t.Errorf("Expected config: %v, received: %v", expectedConfig, receivedConfig)
	}
}

func TestGetConfigQueryHandler_HandleWhenNotFound(t *testing.T) {
	ctx := context.Background()
	query := &GetConfigQuery{
		Key: "Hello",
	}
	expectedError := cerror.NewWithHttpStatusCode(
		cerror.BusinessError,
		fmt.Sprintf("There isn't any value for %s", query.Key),
		http.StatusNotFound,
	)
	repository := mock.NewConfigRepository(mock.NewInMemoryDB())
	handler := NewGetConfigQueryHandler(repository)

	if _, err := handler.Handle(ctx, query); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}

func TestGetConfigQueryHandler_HandleWhenKeyIsEmpty(t *testing.T) {
	ctx := context.Background()
	query := &GetConfigQuery{
		Key: "",
	}
	expectedError := common.NullOrEmptyReferenceError("key")
	repository := mock.NewConfigRepository(mock.NewInMemoryDB())
	handler := NewGetConfigQueryHandler(repository)

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
	repository := mock.NewConfigRepository(mock.NewInMemoryDB())
	handler := NewGetConfigQueryHandler(repository)

	if _, err := handler.Handle(ctx, query); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}
