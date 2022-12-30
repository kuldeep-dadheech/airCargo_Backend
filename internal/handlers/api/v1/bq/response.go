package bq

import (
	"sagebackend/internal/core/domain/services"
	"time"
)

type bq struct {
		Id int `json:"id"`
		Is_quote bool `json:"is_quote"`
		Enquiry_Id int `json:"enquiry_id"`
		Valid_till string `json:"valid_till"`
		Vessel_liner string `json:"vessel_liner"`
		Vessel_total_transit string `json:"vessel_total_transit"`
		Vessel_free_days string `json:"vessel_free_days"`
		Freight_charges_vendor string `json:"freight_charges_vendor"`
		Freight_other_charges_vendor string `json:"freight_other_charges_vendor"`
		Total_buy_amount int `json:"total_buy_amount"`
		Total_sell_amount int `json:"total_sell_amount"`
		Profit_percentage string `json:"profit_percentage"`
		Created_at time.Time `json:"created_at"`
		Modified_at time.Time `json:"modified_at"`
		Charge []services.Charge `json:"charge"`
}