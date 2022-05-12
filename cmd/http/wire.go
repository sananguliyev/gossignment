//go:build wireinject
// +build wireinject

package main

import (
	"github.com/SananGuliyev/gossignment/application/handler"
	"github.com/SananGuliyev/gossignment/config"
	"github.com/SananGuliyev/gossignment/domain/interactor"
	in_memory "github.com/SananGuliyev/gossignment/infrastructure/in-memory"
	"github.com/SananGuliyev/gossignment/infrastructure/mongodb"
	internal_validator "github.com/SananGuliyev/gossignment/infrastructure/validator"
	"github.com/SananGuliyev/gossignment/presenter"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func InitializeServer() (presenter.Server, error) {
	wire.Build(
		config.NewMongoDbConfig,
		handler.NewInMemoryHandler,
		handler.NewRecordHandler,
		in_memory.NewMemRepository,
		in_memory.NewStorage,
		internal_validator.NewMemValidator,
		internal_validator.NewRecordValidator,
		mongodb.NewRecordRepository,
		interactor.NewMemInteractor,
		interactor.NewRecordInteractor,
		presenter.NewNetHttpServer,
		validator.New,
		mongodb.NewStorage,
	)
	return &presenter.NetHttpServer{}, nil
}
