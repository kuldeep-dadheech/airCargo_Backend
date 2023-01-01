//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package api

import (
	"aircargo/internal/configs"
	"aircargo/internal/core/ports"

	"aircargo/internal/core/services/bqsrv"
	"aircargo/internal/core/services/countriessrv"
	"aircargo/internal/core/services/currenciessrv"
	"aircargo/internal/core/services/enquiriessrv"
	"aircargo/internal/core/services/taskssrv"
	"aircargo/internal/core/services/userssrv"
	v1 "aircargo/internal/handlers/api/v1"
	"aircargo/internal/handlers/api/v1/bq"
	"aircargo/internal/handlers/api/v1/countries"
	"aircargo/internal/handlers/api/v1/currencies"
	"aircargo/internal/handlers/api/v1/enquiries"
	"aircargo/internal/handlers/api/v1/tasks"
	"aircargo/internal/handlers/api/v1/users"
	"aircargo/internal/repositories/pgdb/bqrepo"
	"aircargo/internal/repositories/pgdb/cargosrepo"
	"aircargo/internal/repositories/pgdb/chargesrepo"
	"aircargo/internal/repositories/pgdb/countriesrepo"
	"aircargo/internal/repositories/pgdb/tasksrepo"
	"aircargo/internal/repositories/pgdb/currenciesrepo"
	"aircargo/internal/repositories/pgdb/enquiriesrepo"
	"aircargo/internal/repositories/pgdb/usersrepo"

	"github.com/google/wire"
)

func InitializeApp(
	appName configs.AppName,
	baseRouterPrefix configs.BaseRouterPrefix,
	pgDbConfig configs.PgDbConfig,
	logConfig configs.LogConfig,
) (*app, error) {
	wire.Build(
		NewLogger,
		NewPgDbInstance,
		NewGoquInstance,

		//Repositories
		countriesrepo.New,
		currenciesrepo.New,
		usersrepo.New,
		enquiriesrepo.New,
		tasksrepo.New,

		cargosrepo.New,

		bqrepo.New,

		chargesrepo.New,

		//Repo Bindings
		wire.Bind(new(ports.RdbmsCountriesRepository), new(*countriesrepo.Repository)),
		wire.Bind(new(ports.RdbmsCurrenciesRepository), new(*currenciesrepo.Repository)),
		wire.Bind(new(ports.RdbmsUsersRepository), new(*usersrepo.Repository)),
		wire.Bind(new(ports.RdbmsEnquiriesRepository), new(*enquiriesrepo.Repository)),

		wire.Bind(new(ports.RdbmsCargosRepository), new(*cargosrepo.Repository)),
		wire.Bind(new(ports.RdbmsTasksRepository), new(*tasksrepo.Repository)),

		wire.Bind(new(ports.RdbmsBqRepository), new(*bqrepo.Repository)),

		wire.Bind(new(ports.RdbmsChargesRepository), new(*chargesrepo.Repository)),

		//Services
		countriessrv.New,
		currenciessrv.New,
		userssrv.New,
		enquiriessrv.New,
		taskssrv.New,
		bqsrv.New,

		//Service Bindings
		wire.Bind(new(ports.CountriesService), new(*countriessrv.Service)),
		wire.Bind(new(ports.CurrenciesService), new(*currenciessrv.Service)),
		wire.Bind(new(ports.UsersService), new(*userssrv.Service)),
		wire.Bind(new(ports.EnquiriesService), new(*enquiriessrv.Service)),
		wire.Bind(new(ports.TasksService), new(*taskssrv.Service)),
		wire.Bind(new(ports.BqService), new(*bqsrv.Service)),

		//RouteHandlers
		countries.NewHandler,
		currencies.NewHandler,
		users.NewHandler,
		enquiries.NewHandler,
		tasks.NewHandler,

		bq.NewHandler,
		// accounts.NewHandler,

		//RouteHandlerBindings
		wire.Bind(new(countries.RoutesHandler), new(*countries.Handler)),
		wire.Bind(new(currencies.RoutesHandler), new(*currencies.Handler)),
		wire.Bind(new(users.RoutesHandler), new(*users.Handler)),
		wire.Bind(new(enquiries.RoutesHandler), new(*enquiries.Handler)),
		wire.Bind(new(tasks.RoutesHandler), new(*tasks.Handler)),

		wire.Bind(new(bq.RoutesHandler), new(*bq.Handler)),
		// wire.Bind(new(accounts.RoutesHandler), new(*accounts.Handler)),

		//Group Routes
		countries.New,
		currencies.New,
		users.New,
		enquiries.New,
		tasks.New,

		bq.New,
		// accounts.New,
		v1.New,

		//Group Route Bindings
		wire.Bind(new(countries.GroupRoutes), new(*countries.Routes)),
		wire.Bind(new(currencies.GroupRoutes), new(*currencies.Routes)),
		wire.Bind(new(users.GroupRoutes), new(*users.Routes)),
		// wire.Bind(new(accounts.GroupRoutes), new(*accounts.Routes)),
		wire.Bind(new(enquiries.GroupRoutes), new(*enquiries.Routes)),
		wire.Bind(new(tasks.GroupRoutes), new(*tasks.Routes)),

		wire.Bind(new(bq.GroupRoutes), new(*bq.Routes)),
		wire.Bind(new(v1.GroupRoutes), new(*v1.Routes)),

		NewHttpEngine,
		NewApp,
	)
	return &app{}, nil
}
