package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct{}

func (u *LinkHandler) GetOneLink(c *gin.Context) {
	c.String(http.StatusOK, "must implement - GetOneLink")
}

func (u *LinkHandler) GetManyLink(c *gin.Context) {
	c.String(http.StatusOK, "must implement - GetManyLink")
}

func (u *LinkHandler) CreateLink(c *gin.Context) {
	c.String(http.StatusOK, "must implement - CreateLink")
}

func (u *LinkHandler) UpdateLink(c *gin.Context) {
	c.String(http.StatusOK, "must implement - UpdateLink")
}

func (u *LinkHandler) DeleteLink(c *gin.Context) {
	c.String(http.StatusOK, "must implement - DeleteLink")
}
