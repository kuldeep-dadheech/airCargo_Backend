package ports

import (
	"aircargo/internal/core/domain/repositories/rdbms"
	"aircargo/internal/core/domain/services"
)


type RdbmsTasksRepository interface {
	CreateOne(
		bq_id int,
		count int,
	) (int,error)

	SelectOne(id int) (rdbms.Task, error)

	UpdateOne(
		id int,
		bq_id int,
		count int,
	) error
}

type TasksService interface {
	CreateTask(
		bq_id int,
		count int,
	) error

	SelectTask(id int) (services.Task, error)

	UpdateTask(
		id int,
		bq_id int,
		count int,
	) error
}