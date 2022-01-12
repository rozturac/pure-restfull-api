package api

import (
	"fmt"
	"github.com/rozturac/go-mediator"
	"log"
	"net/http"
	"os"
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
	mongoDB := common.NewMongoHelper(config.MongoDB.URI, config.MongoDB.Database, config.MongoDB.Timeout)
	configRepository := persistence.NewConfigRepository(inMemoryDB)
	recordRepository := persistence.NewRecordRepository(mongoDB.GetCollection("records"))

	mediator := ResolveMediator(logger, configRepository, recordRepository)
	configController := controllers.NewConfigController(mediator, config.Globalization)
	recordController := controllers.NewRecordController(mediator, config.Globalization)

	mux := http.NewServeMux()
	healthCheckHandler := func() http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
		}
	}

	mux.Handle("/api/v1/health-check", healthCheckHandler())
	mux.Handle("/api/v1/records", middleware.ErrorMiddleware(recordController.ServeHTTP()))
	mux.Handle("/api/v1/configs", middleware.ErrorMiddleware(configController.ServeHTTP()))

	srv := &http.Server{
		ReadTimeout:  time.Duration(config.Host.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Host.WriteTimeout) * time.Second,
		Addr:         fmt.Sprintf(":%s", GetPort(config)),
		Handler:      mux,
		ErrorLog:     logger,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func ResolveMediator(logger *log.Logger,
	configRepository repository.ConfigRepository, recordRepository repository.RecordQueryRepository) mediator.Mediator {
	m := mediator.Create()
	m.WithBehavior(behavior.NewPerformanceBehavior(logger).Execute)

	fmt.Println(reflect.TypeOf(&create_config.CreateConfigCommand{}).Name())

	err := m.RegisterCommand(
		&create_config.CreateConfigCommand{},
		create_config.NewCreateConfigCommandHandler(configRepository).Handle,
	)
	if err != nil {
		logger.Fatal(err)
		panic(err)
	}

	err = m.RegisterCommand(
		&get_config.GetConfigQuery{},
		get_config.NewGetConfigQueryHandler(configRepository).Handle,
	)
	if err != nil {
		logger.Fatal(err)
		panic(err)
	}

	err = m.RegisterCommand(
		&get_records.GetRecordsByTimeAndCountRangeQuery{},
		get_records.NewGetRecordQueryHandler(recordRepository).Handle,
	)
	if err != nil {
		logger.Fatal(err)
		panic(err)
	}

	return m
}

func GetPort(config *configs.Config) string {
	if envPort := os.Getenv("PORT"); len(envPort) != 0 {
		return envPort
	}

	return config.Host.Port
}
