package services

import (
	"time"
)

type Bq struct {
	Id int `db:"id"`
	Is_quote bool `db:"is_quote"`
	Enquiry_Id int `db:"enquiry_id"`
	Valid_till string `db:"valid_till"`
	Vessel_liner string `db:"vessel_liner"`
	Vessel_total_transit string `db:"vessel_total_transit"`
	Vessel_free_days string `db:"vessel_free_days"`
	Freight_charges_vendor string `db:"freight_charges_vendor"`
	Freight_other_charges_vendor string `db:"freight_other_charges_vendor"`
	Total_buy_amount int `db:"total_buy_amount"`
	Total_sell_amount int `db:"total_sell_amount"`
	Profit_percentage string `db:"profit_percentage"`
	Charge []Charge `db:"charge"`
	Created_at time.Time `db:"created_at"`
	Modified_at time.Time `db:"modified_at"`
}

type Charge struct {
	Id int `db:"id"`
	Bq_id int `db:"bq_id"`
	Charges_name string `db:"charges_name"`
	Units string `db:"units"`
	Tax string `db:"tax"`
	Buy_rate string `db:"buy_rate"`
	Sell_rate string `db:"sell_rate"`
	Total_buy_amount string `db:"total_buy_amount"`
	Total_sell_amount string `db:"total_sell_amount"`
	Charges_type string `db:"charges_type"`
	Created_at time.Time `db:"created_at"`
	Modified_at time.Time `db:"modified_at"`
}

type BqResponse struct {
	Bq Bq
	Charge []Charge
}