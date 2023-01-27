package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mdcaceres/bookstore_users-api/domain/users"
	"github.com/mdcaceres/bookstore_users-api/service"
	"github.com/mdcaceres/bookstore_users-api/utils/errors"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := service.Create(user)
	if err != nil {
		fmt.Println("error in service create")
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
