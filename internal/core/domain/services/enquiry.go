package services

import "time"

type Enquiry struct {
	Id int `json:"id"`
	Customer string `json:"customer"`
	Terms_of_shipment string `json:"terms_of_shipment"`
	Origin_customs bool `json:"origin_customs"`
	Door_pickup bool `json:"door_pickup"`
	Origin_port string `json:"origin_port"`
	Pickup_address string `json:"pickup_address"`
	Cargo_ready_date string `json:"cargo_ready_date"`
	Destination_customs bool `json:"destination_customs"`
	Door_delivery bool `json:"door_delivery"`
	Destination_port string `json:"destination_port"`
	Delivery_address string `json:"delivery_address"`
	Hs_code string `json:"hs_code"`
	Remarks string `json:"remarks"`
	Cargo []Cargo `json:"cargo"`
	Created_at time.Time `json:"created_at"`
	Modified_at time.Time `json:"modified_at"`
}

type Cargo struct {
	Id int `json:"id"`
	Enquiry_id int `json:"enquiry_id"`
	Count int `json:"count"`
	Weight int `json:"weight"`
	Length int `json:"length"`
	Width int `json:"width"`
	Height int `json:"height"`
	Unit string `json:"unit"`
	Volume string `json:"volume"`
	Total_volume string `json:"total_volume"`
	Total_weight string `json:"total_weight"`
	Created_at time.Time `json:"created_at"`
	Modified_at time.Time `json:"modified_at"`
}

type EnquiryResponse struct {
	Enquiry Enquiry
	Cargo []Cargo
}
	