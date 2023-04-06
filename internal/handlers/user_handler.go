package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/genesis-crypto/shortner-go/internal/dto"
	"github.com/genesis-crypto/shortner-go/internal/entities"
	"github.com/genesis-crypto/shortner-go/internal/infra/database"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (h *UserHandler) GetManyUser(c *gin.Context) {
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

	users, err := h.UserDB.FindMany(pageInt, limitInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *UserHandler) GetOneUser(c *gin.Context) {
	c.String(http.StatusOK, "must implement - GetOneUser")
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user dto.CreateUserInput

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	u, err := entities.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	c.String(http.StatusOK, "must implement - UpdateUser")
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	c.String(http.StatusOK, "must implement - DeleteUser")
}

func (h *UserHandler) GetJWT(c *gin.Context) {
	jwt, _ := c.Get("jwt")
	jwtExpiresIn, _ := c.Get("jwtExpiresIn")

	var user dto.GetJWTInput

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"Message": err.Error()})
		return
	}

	if err := u.ValidatePassword(user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
		return
	}

	_, tokenString, _ := jwt.(*jwtauth.JWTAuth).Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn.(int))).Unix(),
	})

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	c.JSON(http.StatusOK, gin.H{"data": accessToken})
}
