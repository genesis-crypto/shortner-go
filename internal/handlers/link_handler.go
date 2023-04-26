package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/genesis-crypto/shortner-go/internal/dto"
	"github.com/genesis-crypto/shortner-go/internal/entities"
	"github.com/genesis-crypto/shortner-go/internal/infra/database"
	"github.com/genesis-crypto/shortner-go/pkg/shortner"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
)

type LinkHandler struct {
	LinkDB  database.LinkInterface
	RedisDB *redis.Client
	Queue   *amqp.Channel
}

func NewLinkHandler(db database.LinkInterface, cache *redis.Client, queue *amqp.Channel) *LinkHandler {
	return &LinkHandler{
		LinkDB:  db,
		RedisDB: cache,
		Queue:   queue,
	}
}

func (u *LinkHandler) GetOneLink(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	hash := c.Param("hash")

	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "missing param",
		})
	}

	cachedData, err := u.RedisDB.Get(ctx, hash).Result()
	if err == nil {
		c.Redirect(http.StatusFound, cachedData)
		return
	}

	links, err := u.LinkDB.FindByHash(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	err = u.RedisDB.Set(ctx, hash, links.Url, 0).Err()
	if err != nil {
		log.Println("Error caching data in Redis:", err)
	}

	c.Redirect(http.StatusFound, links.Url)
}

func (u *LinkHandler) GetManyLink(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	limit := c.DefaultQuery("limit", "1")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	cacheKey := fmt.Sprintf("links:page%d-limit%d", pageInt, limitInt)

	// Check if data is already cached
	cachedData, err := u.RedisDB.Get(ctx, cacheKey).Result()
	if err == nil {
		var links []entities.Link
		err = json.Unmarshal([]byte(cachedData), &links)
		if err != nil {
			log.Println("Error unmarshaling cached data from Redis:", err)
		}
		c.JSON(http.StatusOK, gin.H{"data": links})
		return
	}

	links, err := u.LinkDB.FindMany(pageInt, limitInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	jsonData, err := json.Marshal(links)
	if err != nil {
		log.Println("Error marshaling data for Redis cache:", err)
	} else {
		err = u.RedisDB.Set(ctx, cacheKey, jsonData, time.Minute).Err()
		if err != nil {
			log.Println("Error caching data in Redis:", err)
		}
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
