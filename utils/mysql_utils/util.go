package mysqlutils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/kadekchrisna/openbook-users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

// ParseError parsing mysql error to customable error messagges
func ParseError(err error) *errors.ResErr {
	sqlError, errParse := err.(*mysql.MySQLError)
	if !errParse {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record found.")
		}
		return errors.NewInternalServerError("Error parsing database response.")
	}
	switch sqlError.Number {
	case 1062:
		return errors.NewBadRequestError("Duplicate record.")
	}
	return errors.NewInternalServerError("Error parsing database response.")
}
