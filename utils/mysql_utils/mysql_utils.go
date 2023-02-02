package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/mdcaceres/bookstore_users-api/utils/errors"
	"strings"
)

const (
	NoRowsInResultSet = "no rows in result set"
	UniqueEmail       = "user.email"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), NoRowsInResultSet) {
			return errors.NewNotFoundError("no user found with given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		if strings.Contains(err.Error(), UniqueEmail) {
			return errors.NewBadRequestErr("the email exits in database")
		} else {
			return errors.NewBadRequestErr("invalid data")
		}
	}
	return errors.NewInternalServerError("error processing request")
}
