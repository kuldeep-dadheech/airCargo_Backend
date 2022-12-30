package userssrv

import (
	"fmt"
	"sagebackend/internal/core/domain/repositories/rdbms"
	"sagebackend/internal/core/domain/services"
	"sagebackend/internal/core/ports"
	"sagebackend/pkg/logging"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	logger          logging.Logger
	usersRepository ports.RdbmsUsersRepository

	// use the helper fn for solving the error!
}

func New(
	logger logging.Logger,
	usersRepository ports.RdbmsUsersRepository,
) *Service {
	return &Service{
		logger:          logger,
		usersRepository: usersRepository,
	}
}

func (s *Service) FetchUserById(id int) (services.User, bool, error) {
	person, exists, err := s.usersRepository.GetUserById(id)

	if err != nil {
		return services.User{}, false, fmt.Errorf("error_occored_while_fetching_user!")
	}
	if exists == false {
		return services.User{}, false, fmt.Errorf("user_doesn't_exists_of_given_id!")
	}
	user := s.mapRepoDomainToService(person)
	return user, true, nil

}

func (s *Service) Fetch(Email string) (services.User, bool, error) {
	person, exists, err := s.fetchOneUser(Email)
	if exists == false {
		return services.User{}, false, fmt.Errorf("user_doesnt_exists")
	}
	if err != nil {
		return services.User{}, false, fmt.Errorf("something_went_wrong!")
	}
	return person, true, nil
}

func (s *Service) fetchOneUser(email string) (services.User, bool, error) {
	repoUser, exists, err := s.usersRepository.SelectOne(email)
	if err != nil {
		return services.User{}, exists, err
	}
	if !exists {
		return services.User{}, exists, nil
	}
	return s.mapRepoDomainToService(repoUser), exists, nil
}

func (s *Service) validate(
	name string,
	email string,
	password string,
) error {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	if len(name) < 5 {
		return fmt.Errorf("Name should have more than five characters")
	}
	if email == "" {
		return fmt.Errorf("email_can't_be_empty!")
	}
	if password == "" {
		return fmt.Errorf("password_can'be_empty!")
	}
	return nil

}

func (s *Service) Create(
	name string,
	email string,
	password string,
	role string,
) error {
	err := s.validate(name, email, password)
	if err != nil {
		return err
	}
	pwSlice, er := bcrypt.GenerateFromPassword([]byte(password), 16)
	if er != nil {
		return fmt.Errorf("failed_to_encrypt_the_password!")
	}
	password = string(pwSlice)

	e := s.usersRepository.InsertOne(
		name,
		email,
		password,
		role,
	)
	if e != nil {
		return e
	}
	return nil

}

func (s *Service) mapRepoDomainToService(
	u rdbms.Users,
) services.User {
	return services.User{
		Id:          u.Id,
		Name:        u.Name,
		Email:       u.Email,
		Role:        u.Role,
		Password:    u.Password,
		Created_at:  u.Created_at,
		Modified_at: u.Modified_at,
	}
}
