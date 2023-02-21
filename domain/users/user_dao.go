package users

import (
	"fmt"
	"github.com/mdcaceres/bookstore_users-api/datasources/mysql/users_db"
	"github.com/mdcaceres/bookstore_users-api/utils/errors"
	"github.com/mdcaceres/bookstore_users-api/utils/mysql_utils"
)

const (
	InsertUser       = "INSERT INTO user (first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	GetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM user WHERE id = ?;"
	UpdateUser       = "UPDATE user SET first_name=?, last_name=?, email=? WHERE id=?;"
	DeleteUser       = "DELETE FROM user WHERE id = ?;"
	FindUserByStatus = "SELECT id, first_name, last_name, date_created, status FROM user WHERE status=?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(GetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
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

func (user *User) Delete() *errors.RestErr {
	stmt, stmtErr := users_db.Client.Prepare(DeleteUser)
	if stmtErr != nil {
		return errors.NewInternalServerError(stmtErr.Error())
	}
	defer stmt.Close()

	if _, deleteErr := stmt.Exec(user.Id); deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(FindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	rows, err := stmt.Query(status)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
