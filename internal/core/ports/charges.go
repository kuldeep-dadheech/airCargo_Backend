package ports

import "aircargo/internal/core/domain/repositories/rdbms"

type RdbmsChargesRepository interface {
	CreateOne(
		bq_id int,
		charges_name string,
		units string,
		tax string,
		buy_rate string,
		sell_rate string,
		total_buy_amount string,
		total_sell_amount string,
		charges_type string,
	) error

	SelectAll(bq_id int) ([]rdbms.Charges, error)
	DeleteAll(bq_id int) error
}
