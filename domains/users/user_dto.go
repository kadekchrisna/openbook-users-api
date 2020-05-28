package users

import (
	"strings"

	"github.com/kadekchrisna/openbook-users-api/utils/errors"
)

// User template
type (
	User struct {
		Id          int64  `json:"id"`
		FirstName   string `json:"firstname"`
		LastName    string `json:"lastname"`
		Email       string `json:"email"`
		DateCreated string `json:"date_created"`
		Status      string `json:"status"`
		Password    string `json:"password"` //->> if you want to display value to json and vice versa
		// Password    string `json:"-"`  ->> if you want to not display value to json and vice versa
	}
	Users []User
)

const (
	StatusActive = "active"
)

// Validate user input by giving user object a method validate
func (user *User) Validate() *errors.ResErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	if user.FirstName == "" {
		return errors.NewBadRequestError("invalid FirstName")
	}
	if user.LastName == "" {
		return errors.NewBadRequestError("invalid LastName")
	}
	return nil
}
