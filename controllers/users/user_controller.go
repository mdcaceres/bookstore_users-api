package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mdcaceres/bookstore_users-api/domain/users"
	"github.com/mdcaceres/bookstore_users-api/service"
	"github.com/mdcaceres/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestErr("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := service.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

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
