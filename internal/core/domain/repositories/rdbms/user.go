package rdbms

import (
	"time"
)

type Users struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	Role        string    `db:"role"`
	Password    string    `db:"password"`
	Created_at  time.Time `db:"created_at"`
	Modified_at time.Time `db:"modified_at"`
}
