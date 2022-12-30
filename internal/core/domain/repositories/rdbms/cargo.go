package rdbms

import (
	"time"
)

type Cargo struct {
	Id int `db:"id"`
	Enquiry_id int `db:"enquiry_id"`
	Count int `db:"count"`
	Length int `db:"length"`
	Width int `db:"width"`
	Height int `db:"height"`
	Volume string `db:"volume"`
	Weight int `db:"weight"`
	Unit string `db:"unit"`
	Total_volume string `db:"total_volume"`
	Total_weight string `db:"total_weight"`
	Created_at time.Time `db:"created_at"`
	Modified_at time.Time `db:"modified_at"`
}