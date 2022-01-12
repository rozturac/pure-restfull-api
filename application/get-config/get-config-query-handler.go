package get_config

import (
	"context"
	"github.com/rozturac/go-mediator"
	"pure-restfull-api/application/common"
	"pure-restfull-api/domain/repository"
)

type GetConfigQueryHandler struct {
	repository repository.ConfigRepository
}

func NewGetConfigQueryHandler(repository repository.ConfigRepository) *GetConfigQueryHandler {
	return &GetConfigQueryHandler{
		repository: repository,
	}
}

func (g *GetConfigQueryHandler) Handle(ctx context.Context, command mediator.Command) (interface{}, error) {
	query, ok := command.(*GetConfigQuery)
	if !ok {
		return nil, common.UnExpectedCommand(query)
	}

	if len(query.Key) == 0 {
		common.NullOrEmptyReferenceError("key")
	}

	config, err := g.repository.GetByKey(query.Key)
	return config, err
}
