package api

import (
	"github.com/rozturac/go-mediator"
	"log"
	"pure-restfull-api/application/common/behavior"
	create_config "pure-restfull-api/application/create-config"
	get_config "pure-restfull-api/application/get-config"
	get_records "pure-restfull-api/application/get-records"
	"pure-restfull-api/domain/repository"
)

func Init() {

}

func ResolveMediator(logger *log.Logger,
	configRepository repository.ConfigRepository, recordRepository repository.RecordRepository) mediator.Mediator {
	m := mediator.Create()
	m.WithBehavior(behavior.NewPerformanceBehavior(logger).Execute)

	err := m.RegisterCommand(
		&create_config.CreateConfigCommand{},
		create_config.NewCreateConfigCommandHandler(configRepository).Handle,
	)
	if err != nil {
		panic(err)
	}

	err = m.RegisterCommand(
		&get_config.GetConfigQuery{},
		get_config.NewGetConfigQueryHandler(configRepository).Handle,
	)
	if err != nil {
		panic(err)
	}

	err = m.RegisterCommand(
		&get_records.GetRecordQueryHandler{},
		get_records.NewGetRecordQueryHandler(recordRepository).Handle,
	)
	if err != nil {
		panic(err)
	}

	return m
}
