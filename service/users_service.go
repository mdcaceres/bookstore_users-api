package service

import (
	"github.com/mdcaceres/bookstore_users-api/domain/users"
	"github.com/mdcaceres/bookstore_users-api/utils/errors"
)

func Create(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
