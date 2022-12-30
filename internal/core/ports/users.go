package ports

import (
	"aircargo/internal/core/domain/repositories/rdbms"
	"aircargo/internal/core/domain/services"
)

type RdbmsUsersRepository interface {
	SelectOne(email string) (rdbms.Users, bool, error)
	InsertOne(
		name string,
		email string,
		password string,
		role string,
	) error

	GetUserById(id int) (rdbms.Users, bool, error)
}

type UsersService interface {
	Fetch(
		Email string,
	) (services.User, bool, error)
	Create(
		name string,
		email string,
		password string,
		role string,
	) error
	FetchUserById(id int) (services.User, bool, error)
}
