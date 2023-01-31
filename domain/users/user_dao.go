package users

import (
	"fmt"
	"github.com/mdcaceres/bookstore_users-api/datasources/mysql/users_db"
	"github.com/mdcaceres/bookstore_users-api/utils/dates"
	"github.com/mdcaceres/bookstore_users-api/utils/errors"
	"strings"
)

const (
	UNIQUE_EMAIL = "user.email"
	INSERT_USER  = "INSERT INTO user (first_name, last_name, email, date_created) VALUES (?,?,?,?);"
)

var (
	usersDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {

	}
	result := usersDb[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user with id %d", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(INSERT_USER)
	if err != nil {
		return errors.NewInternalServerError("error when tying to save user")
	}
	defer stmt.Close()

	user.DateCreated = dates.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		if strings.Contains(saveErr.Error(), UNIQUE_EMAIL) {
			return errors.NewBadRequestErr(
				fmt.Sprintf("email %s is already exists", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when tying to save user : %s ", saveErr.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when tying to save user : %s", err.Error()))
	}

	user.Id = userId

	return nil
}
