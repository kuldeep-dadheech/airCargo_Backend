package chargesrepo

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

func (r *Repository) CreateOne(
	bq_id int,
	charges_name string,
	units string,
	tax string,
	buy_rate string,
	sell_rate string,
	total_buy_amount string,
	total_sell_amount string,
	charges_type string,
) error {
	_, err := r.goquDB.Insert(TABLE).Prepared(true).Rows(
		goqu.Record{
			BQ_ID:             bq_id,
			CHARGES_NAME:      charges_name,
			UNITS:             units,
			TAX:               tax,
			BUY_RATE:          buy_rate,
			SELL_RATE:         sell_rate,
			TOTAL_BUY_AMOUNT:  total_buy_amount,
			TOTAL_SELL_AMOUNT: total_sell_amount,
			CHARGES_TYPE:      charges_type,
		},
	).Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SelectAll(bq_id int) ([]rdbms.Charges, error) {
	var charges []rdbms.Charges
	err := r.goquDB.From(TABLE).Select().Where(goqu.C(BQ_ID).Eq(bq_id)).ScanStructs(&charges)
	if err != nil {
		return nil, err
	}
	return charges, nil
}

func (r *Repository) DeleteAll(bq_id int) error {
	_, err := r.goquDB.Delete(TABLE).Where(goqu.C(BQ_ID).Eq(bq_id)).Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}
