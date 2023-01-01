package usersrepo

import (
	"aircargo/internal/core/domain/repositories/rdbms"
	"aircargo/pkg/logging"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type Repository struct {
	logger logging.Logger
	goquDB *goqu.Database
}

func New(
	logger logging.Logger,
	goquDB *goqu.Database,
) *Repository {
	return &Repository{
		logger: logger,
		goquDB: goquDB,
	}
}

func (r *Repository) SelectOne(
	email string,
) (rdbms.Users, bool, error) {
	var user rdbms.Users
	exists, err := r.goquDB.From(
		TABLE,
	).Select(
		ID,
		NAME,
		EMAIL,
		ROLE,
		PASSWORD,
		CREATED_AT,
		MODIFIED_AT,
	).Where(goqu.C(EMAIL).Eq(email)).ScanStruct(&user)
	// fmt.Println(rdbms.Users{},"dfgdfg")
	if err != nil {
		return rdbms.Users{}, false, err
	}
	return user, exists, nil
}

func (r *Repository) InsertOne(
	name string,
	email string,
	password string,
	role string,
) error {
	var userId Id
	insert := r.goquDB.Insert(TABLE).Prepared(true).Rows(
		goqu.Record{
			NAME:     name,
			EMAIL:    email,
			ROLE:     role,
			PASSWORD: password,
		},
	).Returning("id")
	if _, err := insert.Executor().ScanStruct(&userId); err != nil {
		return err
	} else {
		fmt.Println("Inserted 1 user")
	}
	return nil

}

func (r *Repository) GetUserById(id int) (rdbms.Users, bool, error) {
	var user rdbms.Users
	exist, err := r.goquDB.From(
		TABLE,
	).Prepared(true).Select(
		ID,
		NAME,
		EMAIL,
		PASSWORD,
		ROLE,
		CREATED_AT,
		MODIFIED_AT,
	).Where(goqu.C(ID).Eq(id)).ScanStruct(&user)
	if err != nil {
		return rdbms.Users{}, exist, err
	}
	if !exist {
		return rdbms.Users{}, exist, err
	}
	fmt.Println(user)
	return user, exist, nil
}

type Id struct {
	Id int `db:"id"`
}
