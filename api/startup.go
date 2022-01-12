package api

import (
	"fmt"
	"github.com/rozturac/go-mediator"
	"log"
	"net/http"
	"pure-restfull-api/api/configs"
	"pure-restfull-api/api/controllers"
	"pure-restfull-api/api/middleware"
	"pure-restfull-api/application/common/behavior"
	create_config "pure-restfull-api/application/create-config"
	get_config "pure-restfull-api/application/get-config"
	get_records "pure-restfull-api/application/get-records"
	"pure-restfull-api/domain/repository"
	"pure-restfull-api/infrastructure/common"
	"pure-restfull-api/infrastructure/persistence"
	"reflect"
	"time"
)

func Init(config *configs.Config) {
	logger := log.Default()
	inMemoryDB := common.NewInMemoryDB()
	configRepository := persistence.NewConfigRepository(inMemoryDB)

	mediator := ResolveMediator(logger, configRepository, nil)
	configController := controllers.NewConfigController(mediator, config.Globalization)

	mux := http.NewServeMux()
	healthCheckHandler := func() http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
		}
	}

	mux.Handle("/api/v1/health-check", healthCheckHandler())
	mux.Handle("/api/v1/configs", middleware.ErrorMiddleware(configController.ServeHTTP()))

	srv := &http.Server{
		ReadTimeout:  time.Duration(config.Host.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Host.WriteTimeout) * time.Second,
		Addr:         fmt.Sprintf(":%d", config.Host.Port),
		Handler:      mux,
		ErrorLog:     logger,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func ResolveMediator(logger *log.Logger,
	configRepository repository.ConfigRepository, recordRepository repository.RecordRepository) mediator.Mediator {
	m := mediator.Create()
	m.WithBehavior(behavior.NewPerformanceBehavior(logger).Execute)

	fmt.Println(reflect.TypeOf(&create_config.CreateConfigCommand{}).Name())

	err := m.RegisterCommand(
		&create_config.CreateConfigCommand{},
		create_config.NewCreateConfigCommandHandler(configRepository).Handle,
	)
	if err != nil {
		panic(err)
	}

	cmd := &get_config.GetConfigQuery{}
	err = m.RegisterCommand(
		cmd,
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
