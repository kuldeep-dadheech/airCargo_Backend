package cargosrepo

import (
	"fmt"
	"sagebackend/internal/core/domain/repositories/rdbms"
	"sagebackend/pkg/logging"

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

func (r *Repository) CreateOne(
	enquiry_id int,
		count int,
		length int,
		width int,
		height int,
		volume string,
		weight int,
		unit string,
		total_volume string,
		total_weight string,
) error {
	_, err := r.goquDB.Insert(TABLE).Prepared(true).Rows(
		goqu.Record{
			ENQUIRY_ID:   enquiry_id,
			COUNT:        count,
			LENGTH:       length,
			WIDTH:        width,
			HEIGHT:       height,
			VOLUME:       volume,
			WEIGHT:       weight,
			UNIT:         unit,
			TOTAL_VOLUME: total_volume,
			TOTAL_WEIGHT: total_weight,
		},
	).Executor().Exec()
	if err != nil {
		return err
	}
	fmt.Println("gereREJRFHGJF")
	return nil
}


func (r *Repository) SelectAll(enquiry_id int) ([]rdbms.Cargo, error) {
	var cargos []rdbms.Cargo
	err := r.goquDB.From(TABLE).Select().Where(goqu.C(ENQUIRY_ID).Eq(enquiry_id)).ScanStructs(&cargos)
	if err != nil {
		return nil, err
	}
	return cargos, nil
}


func (r *Repository) DeleteAll(enquiry_id int) error {
	_, err := r.goquDB.Delete(TABLE).Where(goqu.C(ENQUIRY_ID).Eq(enquiry_id)).Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}