package tasksrepo

import (
	"aircargo/internal/core/domain/repositories/rdbms"
	"aircargo/pkg/logging"

	"github.com/doug-martin/goqu/v9"
)

type Repository struct {
	logger logging.Logger
	goquDB *goqu.Database
}

func New(
	logger logging.Logger,
	goquDB *goqu.Database,
) *Repository {
	return &Repository{
		logger: logger,
		goquDB: goquDB,
	}
}

type Id struct {
	Id int `db:"id"`
}

func (r *Repository) CreateOne(
	bq_id int,
	count int,
) (int, error) {
	var id Id
	_, err := r.goquDB.Insert(TABLE).Prepared(true).Rows(
		goqu.Record{
			BQ_ID: bq_id,
			COUNT: count,
		},
	).Returning("id").Executor().ScanStruct(&id)
	if err != nil {
		return 0, err
	}
	return id.Id, nil
}

func (r *Repository) SelectOne(id int) (rdbms.Task, error) {
	var task rdbms.Task
	_, err := r.goquDB.From(TABLE).Select().Where(goqu.C(ID).Eq(id)).ScanStruct(&task)

	if err != nil {
		return rdbms.Task{}, err
	}
	return task, nil
}

func (r *Repository) UpdateOne(
	id int,
	bq_id int,
	count int,
) error {
	_, err := r.goquDB.Update(TABLE).Set(
		goqu.Record{
			BQ_ID: bq_id,
			COUNT: count,
		},
	).Where(goqu.C(ID).Eq(id)).Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}
