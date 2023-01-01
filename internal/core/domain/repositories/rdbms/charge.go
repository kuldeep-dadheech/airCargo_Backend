package rdbms

import (
	"time"
)

type Charges struct {
	Id                int       `db:"id"`
	Bq_id             int       `db:"bq_id"`
	Charges_name      string    `db:"charges_name"`
	Units             string    `db:"units"`
	Tax               string    `db:"tax"`
	Buy_rate          string    `db:"buy_rate"`
	Sell_rate         string    `db:"sell_rate"`
	Total_buy_amount  string    `db:"total_buy_amount"`
	Total_sell_amount string    `db:"total_sell_amount"`
	Charges_type      string    `db:"charges_type"`
	Created_at        time.Time `db:"created_at"`
	Modified_at       time.Time `db:"modified_at"`
}
