package bqrepo

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

type Id struct {
	Id int `db:"id"`
}

func (r *Repository) CreateOne(
	is_quote bool,
	enquiry_id int,
	valid_till string,
	vessel_liner string,
	vessel_total_transit string,
	vessel_free_days string,
	freight_charges_vendor string,
	freight_other_charges_vendor string,
	total_buy_amount int,
	total_sell_amount int,
	profit_percentage string,
) (int, error) {
	var id Id
	_, err := r.goquDB.Insert(TABLE).Prepared(true).Rows(
		goqu.Record{
			IS_QUOTE:                     is_quote,
			ID:                           enquiry_id,
			VALID_TILL:                   valid_till,
			VESSEL_LINER:                 vessel_liner,
			VESSEL_TOTAL_TRANSIT:         vessel_total_transit,
			VESSEL_FREE_DAYS:             vessel_free_days,
			FREIGHT_CHARGES_VENDOR:       freight_charges_vendor,
			FREIGHT_OTHER_CHARGES_VENDOR: freight_other_charges_vendor,
			TOTAL_BUY_AMOUNT:             total_buy_amount,
			TOTAL_SELL_AMOUNT:            total_sell_amount,
			PROFIT_PERCENTAGE:            profit_percentage,
		},
	).Returning("id").Executor().ScanStruct(&id)
	fmt.Println(err, enquiry_id, "bqrepo")
	if err != nil {
		return 0, err
	}
	return id.Id, nil
}

func (r *Repository) SelectOne(id int) (rdbms.Bq, error) {
	var bq rdbms.Bq
	_, err := r.goquDB.From(TABLE).Select().Where(goqu.C(ID).Eq(id)).ScanStruct(&bq)

	if err != nil {
		return rdbms.Bq{}, err
	}
	return bq, nil
}

func (r *Repository) SelectAll() ([]rdbms.Bq, error) {
	var bqs []rdbms.Bq
	err := r.goquDB.From(TABLE).Select().Order(
		goqu.I(CREATED_AT).Asc(),
	).ScanStructs(&bqs)
	if err != nil {
		return nil, err
	}
	fmt.Println(bqs, "adsadasdsa")
	return bqs, nil
}

func (r *Repository) UpdateOne(
	id int,
	enquiry_id int,
	is_quote bool,
	valid_till string,
	vessel_liner string,
	vessel_total_transit string,
	vessel_free_days string,
	freight_charges_vendor string,
	freight_other_charges_vendor string,
	total_buy_amount int,
	total_sell_amount int,
	profit_percentage string,
) error {
	_, err := r.goquDB.Update(TABLE).Prepared(true).Set(
		goqu.Record{
			ID:                           enquiry_id,
			IS_QUOTE:                     is_quote,
			VALID_TILL:                   valid_till,
			VESSEL_LINER:                 vessel_liner,
			VESSEL_TOTAL_TRANSIT:         vessel_total_transit,
			VESSEL_FREE_DAYS:             vessel_free_days,
			FREIGHT_CHARGES_VENDOR:       freight_charges_vendor,
			FREIGHT_OTHER_CHARGES_VENDOR: freight_other_charges_vendor,
			TOTAL_BUY_AMOUNT:             total_buy_amount,
			TOTAL_SELL_AMOUNT:            total_sell_amount,
			PROFIT_PERCENTAGE:            profit_percentage,
		},
	).Where(goqu.C(ID).Eq(enquiry_id)).Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteOne(id int) error {
	_, err := r.goquDB.Delete(TABLE).Where(goqu.C(ID).Eq(id)).Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}
