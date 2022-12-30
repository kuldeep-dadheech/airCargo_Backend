package services

import "time"

type User struct {
	Id int
	Name string
	Email string
	Role string
	Password string
	Created_at time.Time
	Modified_at time.Time
}

