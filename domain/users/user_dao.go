package users

import (
	"fmt"
	"github.com/mdcaceres/bookstore_users-api/datasources/mysql/users_db"
	"github.com/mdcaceres/bookstore_users-api/utils/dates"
	"github.com/mdcaceres/bookstore_users-api/utils/errors"
	"strings"
)

const (
	UNIQUE_EMAIL         = "user.email"
	NO_ROWS_IN_RESULTSET = "no rows in result set"
	INSERT_USER          = "INSERT INTO user (first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	GET_USER             = "SELECT id, first_name, last_name, email, date_created FROM user WHERE id = ?;"
)

var (
	usersDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(GET_USER)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), NO_ROWS_IN_RESULTSET) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error while trying get user with id %d: %s ", user.Id, err.Error()))
	}

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
