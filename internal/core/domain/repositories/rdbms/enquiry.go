package rdbms

import (
	"time"
)

type Enquiry struct {
	Id                  int       `db:"id"`
	Customer            string    `db:"customer"`
	Terms_of_shipment   string    `db:"terms_of_shipment"`
	Origin_customs      bool      `db:"origin_customs"`
	Door_pickup         bool      `db:"door_pickup"`
	Origin_port         string    `db:"origin_port"`
	Pickup_address      string    `db:"pickup_address"`
	Cargo_ready_date    string    `db:"cargo_ready_date"`
	Destination_customs bool      `db:"destination_customs"`
	Door_delivery       bool      `db:"door_delivery"`
	Destination_port    string    `db:"destination_port"`
	Delivery_address    string    `db:"delivery_address"`
	Hs_code             string    `db:"hs_code"`
	Remarks             string    `db:"remarks"`
	Created_at          time.Time `db:"created_at"`
	Modified_at         time.Time `db:"modified_at"`
}
