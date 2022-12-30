// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"sagebackend/internal/configs"
	"sagebackend/internal/core/services/bqsrv"
	"sagebackend/internal/core/services/countriessrv"
	"sagebackend/internal/core/services/currenciessrv"
	"sagebackend/internal/core/services/enquiriessrv"
	"sagebackend/internal/core/services/userssrv"
	"sagebackend/internal/handlers/api/v1"
	"sagebackend/internal/handlers/api/v1/bq"
	"sagebackend/internal/handlers/api/v1/countries"
	"sagebackend/internal/handlers/api/v1/currencies"
	"sagebackend/internal/handlers/api/v1/enquiries"
	"sagebackend/internal/handlers/api/v1/users"
	"sagebackend/internal/repositories/pgdb/bqrepo"
	"sagebackend/internal/repositories/pgdb/cargosrepo"
	"sagebackend/internal/repositories/pgdb/chargesrepo"
	"sagebackend/internal/repositories/pgdb/countriesrepo"
	"sagebackend/internal/repositories/pgdb/currenciesrepo"
	"sagebackend/internal/repositories/pgdb/enquiriesrepo"
	"sagebackend/internal/repositories/pgdb/usersrepo"
)

import (
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Injectors from wire.go:

func InitializeApp(appName configs.AppName, baseRouterPrefix configs.BaseRouterPrefix, pgDbConfig configs.PgDbConfig, logConfig configs.LogConfig) (*app, error) {
	logger, err := NewLogger(appName, logConfig)
	if err != nil {
		return nil, err
	}
	db, err := NewPgDbInstance(pgDbConfig)
	if err != nil {
		return nil, err
	}
	database := NewGoquInstance(db)
	repository := countriesrepo.New(logger, database)
	service := countriessrv.New(logger, repository)
	handler := countries.NewHandler(logger, service)
	routes := countries.New(handler)
	currenciesrepoRepository := currenciesrepo.New(logger, database)
	currenciessrvService := currenciessrv.New(logger, currenciesrepoRepository)
	currenciesHandler := currencies.NewHandler(logger, currenciessrvService)
	currenciesRoutes := currencies.New(currenciesHandler)
	usersrepoRepository := usersrepo.New(logger, database)
	userssrvService := userssrv.New(logger, usersrepoRepository)
	usersHandler := users.NewHandler(logger, userssrvService)
	usersRoutes := users.New(usersHandler)
	enquiriesrepoRepository := enquiriesrepo.New(logger, database)
	cargosrepoRepository := cargosrepo.New(logger, database)
	enquiriessrvService := enquiriessrv.New(enquiriesrepoRepository, cargosrepoRepository)
	enquiriesHandler := enquiries.NewHandler(logger, enquiriessrvService)
	enquiriesRoutes := enquiries.New(enquiriesHandler)
	chargesrepoRepository := chargesrepo.New(logger, database)
	bqrepoRepository := bqrepo.New(logger, database)
	bqsrvService := bqsrv.New(chargesrepoRepository, bqrepoRepository)
	bqHandler := bq.NewHandler(logger, bqsrvService)
	bqRoutes := bq.New(bqHandler)
	v1Routes := v1.New(routes, currenciesRoutes, usersRoutes, enquiriesRoutes, bqRoutes)
	engine := NewHttpEngine(v1Routes, baseRouterPrefix)
	apiApp := NewApp(engine)
	return apiApp, nil
}