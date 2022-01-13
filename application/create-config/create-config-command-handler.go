package create_config

import (
	"context"
	"github.com/rozturac/go-mediator"
	"pure-restfull-api/application/common"
	"pure-restfull-api/domain/entity"
	"pure-restfull-api/domain/repository"
)

type CreateConfigCommandHandler struct {
	repository repository.ConfigRepository
}

func NewCreateConfigCommandHandler(repository repository.ConfigRepository) *CreateConfigCommandHandler {
	return &CreateConfigCommandHandler{
		repository: repository,
	}
}

func (c *CreateConfigCommandHandler) Handle(ctx context.Context, command mediator.Command) (interface{}, error) {
	cmd, ok := command.(*CreateConfigCommand)
	if !ok {
		return nil, common.UnExpectedCommand("mediator.Command", cmd)
	}

	if len(cmd.Key) == 0 {
		return nil, common.NullOrEmptyReferenceError("Key")
	}

	if len(cmd.Value) == 0 {
		return nil, common.NullOrEmptyReferenceError("Value")
	}

	config := entity.CreateConfig(cmd.Key, cmd.Value)
	err := c.repository.Insert(config)
	return nil, err
}
