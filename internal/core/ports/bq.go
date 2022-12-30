package ports

import (
	"aircargo/internal/core/domain/repositories/rdbms"
	"aircargo/internal/core/domain/services"
)

type RdbmsBqRepository interface {
	CreateOne(
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
	) (int, error)
	SelectAll() ([]rdbms.Bq, error)
	SelectOne(id int) (rdbms.Bq, error)
	DeleteOne(id int) error
	UpdateOne(
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
	) error
}

type BqService interface {
	CreateBq(
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
		charge []services.Charge,
	) error

	GetAllBq() ([]services.Bq, error)
	GetBqByID(id int) (services.BqResponse, error)
	DeleteBq(id int) error
	UpdateBq(
		id int,
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
		charge []services.Charge,
	) error
}
