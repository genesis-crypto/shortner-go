package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/genesis-crypto/shortner-go/internal/dto"
	"github.com/genesis-crypto/shortner-go/internal/entities"
	"github.com/genesis-crypto/shortner-go/internal/infra/database"
	"github.com/genesis-crypto/shortner-go/pkg/shortner"
	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	LinkDB database.LinkInterface
}

func NewLinkHandler(db database.LinkInterface) *LinkHandler {
	return &LinkHandler{
		LinkDB: db,
	}
}

func (u *LinkHandler) GetOneLink(c *gin.Context) {
	hash := c.Param("hash")

	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "missing param",
		})
	}

	links, err := u.LinkDB.FindByHash(hash)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, links.Url)
}

func (u *LinkHandler) GetManyLink(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	limit := c.DefaultQuery("limit", "1")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	links, err := u.LinkDB.FindMany(pageInt, limitInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": links})
}

func (u *LinkHandler) CreateLink(c *gin.Context) {
	var link dto.CreateLinkInput

	err := json.NewDecoder(c.Request.Body).Decode(&link)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	uuid, ok := c.Get("uuid")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Message": "missing uuid",
		})
		return
	}
	sortedHash := shortner.GenerateShortLink(link.Url, uuid.(string))
	l, err := entities.NewLink(link.Url, sortedHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	err = u.LinkDB.Create(l)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Message": "ok"})
}

func (u *LinkHandler) UpdateLink(c *gin.Context) {
	c.String(http.StatusOK, "must implement - UpdateLink")
}

func (u *LinkHandler) DeleteLink(c *gin.Context) {
	c.String(http.StatusOK, "must implement - DeleteLink")
}
