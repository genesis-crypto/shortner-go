package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func (u *UserHandler) GetOneUser(c *gin.Context) {
	c.String(http.StatusOK, "must implement - GetOneUser")
}

func (u *UserHandler) GetManyUser(c *gin.Context) {
	c.String(http.StatusOK, "must implement - GetManyUser")
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	c.String(http.StatusOK, "must implement - CreateUser")
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	c.String(http.StatusOK, "must implement - UpdateUser")
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	c.String(http.StatusOK, "must implement - DeleteUser")
}
