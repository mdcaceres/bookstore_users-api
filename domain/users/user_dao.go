package users

import (
	"fmt"
	"github.com/mdcaceres/bookstore_users-api/utils/errors"
)

var (
	usersDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
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
	current := usersDb[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestErr(fmt.Sprintf("email %s already registered", user.Id))
		}
		return errors.NewBadRequestErr(fmt.Sprintf("user %d añready exists", user.Id))
	}
	usersDb[user.Id] = user
	return nil
}
