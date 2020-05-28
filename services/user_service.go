package services

import (
	"fmt"

	"github.com/kadekchrisna/openbook/domains/users"
	"github.com/kadekchrisna/openbook/logger"
	"github.com/kadekchrisna/openbook/utils/errors"
)

type (
	usersService struct{}

	// UsersServicesInterface polymorph
	UsersServicesInterface interface {
		CreateUser(users.User) (*users.User, *errors.ResErr)
		GetUser(int64) (*users.User, *errors.ResErr)
		UpdateUser(bool, users.User) (*users.User, *errors.ResErr)
		DeleteUser(int64) *errors.ResErr
		SearchUser(string) (users.Users, *errors.ResErr)
		LogginUser(users.LoginRequest) (*users.User, *errors.ResErr)
	}
)

var (
	// UsersService for make mock easier and grouping the user services on usersServices
	UsersService UsersServicesInterface = &usersService{}
)

// CreateUser services for bussiness logic
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.ResErr) {
	if err := user.Validate(); err != nil {
		fmt.Printf("Error on CreateUser \n%s\n", err)
		return nil, err
	}
	if err := user.Save(); err != nil {
		fmt.Printf("Error on CreateUser \n%s\n", err)
		return nil, err
	}
	return &user, nil
}

// GetUser ads
func (s *usersService) GetUser(user int64) (*users.User, *errors.ResErr) {
	result := &users.User{Id: user}
	if err := result.Get(); err != nil {
		logger.Info(fmt.Sprintf("Error on GetUser \n%s\n", err))
		fmt.Printf("Error on GetUser \n%s\n", err)
		return nil, err
	}
	return result, nil
}

// UpdateUser asd
func (s *usersService) UpdateUser(partial bool, user users.User) (*users.User, *errors.ResErr) {
	current, errGetUser := s.GetUser(user.Id)
	if errGetUser != nil {
		fmt.Printf("Error on GetUser \n%s\n", errGetUser.Message)
		return nil, errGetUser
	}
	if partial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		if err := user.Validate(); err != nil {
			fmt.Printf("Body must contain email, firstname, and lastname. \n%s\n", err.Message)
			return nil, err
		}
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		fmt.Printf("Error on Updating user. \n%s\n", errGetUser.Message)
		return nil, err
	}
	return current, nil
}

// DeleteUser deleting user
func (s *usersService) DeleteUser(userID int64) *errors.ResErr {
	var user users.User
	user.Id = userID

	_, err := s.GetUser(user.Id)
	if err != nil {
		fmt.Printf("Error on deleting user %d. %s", userID, err.Message)
		return err
	}

	if err := user.Delete(); err != nil {
		fmt.Printf("Error on deleting user %d. %s", userID, err.Message)
		return err
	}
	return nil
}

// Search find user by status
func (s *usersService) SearchUser(search string) (users.Users, *errors.ResErr) {
	dao := &users.User{}
	result, errFind := dao.FindUserByStatus(search)
	if errFind != nil {
		fmt.Printf("Error when search %s", search)
		return nil, errFind
	}
	return result, nil
}

func (s *usersService) LogginUser(user users.LoginRequest) (*users.User, *errors.ResErr) {
	dao := users.User{}
	dao.Email = user.Email
	dao.Password = user.Password
	if err := dao.FindUserByEmailAndPassword(); err != nil {
		return nil, err
	}
	return &dao, nil
}
