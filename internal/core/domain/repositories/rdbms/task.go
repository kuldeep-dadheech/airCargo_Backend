package rdbms

import (
	"time"
)

type Task struct {
	Id          int       `db:"id"`
	Bq_id       int       `db:"bq_id"`
	Count       int       `db:"count"`
	Created_at  time.Time `db:"created_at"`
	Modified_at time.Time `db:"modified_at"`
}
