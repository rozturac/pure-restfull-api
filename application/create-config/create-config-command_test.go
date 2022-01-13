package create_config

import (
	"context"
	"pure-restfull-api/application/common"
	"pure-restfull-api/domain/entity"
	"pure-restfull-api/mock"
	"testing"
)

func TestCreateConfigCommandHandler_HandleWhenSuccess(t *testing.T) {
	ctx := context.Background()
	command := &CreateConfigCommand{
		Key:   "Hello",
		Value: "World",
	}
	expectedConfig := entity.CreateConfig(command.Key, command.Value)
	repository := mock.NewConfigRepository(mock.NewInMemoryDB())
	handler := NewCreateConfigCommandHandler(repository)

	_, err := handler.Handle(ctx, command)
	if err != nil {
		t.Errorf("Expected is success, received: %v", err)
	}

	if receivedConfig, _ := repository.GetByKey(command.Key); receivedConfig.GetValue() != expectedConfig.GetValue() {
		t.Errorf("Expected config: %v, received: %v", expectedConfig, receivedConfig)
	}
}

func TestCreateConfigCommandHandler_HandleWhenKeyIsEmpty(t *testing.T) {
	ctx := context.Background()
	command := &CreateConfigCommand{
		Key:   "",
		Value: "World",
	}
	expectedError := common.NullOrEmptyReferenceError("Key")
	repository := mock.NewConfigRepository(mock.NewInMemoryDB())
	handler := NewCreateConfigCommandHandler(repository)

	if _, err := handler.Handle(ctx, command); err == nil || expectedError.Error() != err.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}

func TestCreateConfigCommandHandler_HandleWhenValueIsEmpty(t *testing.T) {
	ctx := context.Background()
	command := &CreateConfigCommand{
		Key:   "Hello",
		Value: "",
	}
	expectedError := common.NullOrEmptyReferenceError("Value")
	repository := mock.NewConfigRepository(mock.NewInMemoryDB())
	handler := NewCreateConfigCommandHandler(repository)

	if _, err := handler.Handle(ctx, command); err == nil || expectedError.Error() != err.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}

func TestGetConfigQueryHandler_HandleWhenQueryUnCorrect(t *testing.T) {
	ctx := context.Background()
	type UnexpectedCommand struct {
	}
	command := &UnexpectedCommand{}
	expectedError := common.UnExpectedCommand("mediator.Command", command)
	repository := mock.NewConfigRepository(mock.NewInMemoryDB())
	handler := NewCreateConfigCommandHandler(repository)

	if _, err := handler.Handle(ctx, command); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, received: %v", expectedError, err)
	}
}
