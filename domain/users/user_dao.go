package users

import (
	"github.com/mdcaceres/bookstore_users-api/datasources/mysql/users_db"
	"github.com/mdcaceres/bookstore_users-api/utils/dates"
	"github.com/mdcaceres/bookstore_users-api/utils/errors"
	"github.com/mdcaceres/bookstore_users-api/utils/mysql_utils"
)

const (
	InsertUser = "INSERT INTO user (first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	GetUser    = "SELECT id, first_name, last_name, email, date_created FROM user WHERE id = ?;"
	UpdateUser = "UPDATE user SET first_name=?, last_name=?, email=? WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(GetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(InsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dates.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}

	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, stmtErr := users_db.Client.Prepare(UpdateUser)
	if stmtErr != nil {
		return errors.NewInternalServerError(stmtErr.Error())
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if execErr != nil {
		return mysql_utils.ParseError(execErr)
	}

	return nil
}
