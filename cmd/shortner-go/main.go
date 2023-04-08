package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/genesis-crypto/shortner-go/configs"
	entities "github.com/genesis-crypto/shortner-go/internal/entities"
	handler "github.com/genesis-crypto/shortner-go/internal/handlers"
	"github.com/genesis-crypto/shortner-go/internal/infra/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.User{}, &entities.Link{})

	r := gin.Default()

	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)

	userDB := database.NewUser(db)
	userHandler := handler.NewUserHandler(userDB)

	r.Use(func(ctx *gin.Context) {
		ctx.Set("jwt", config.TokenAuth)
		ctx.Set("jwtExpiresIn", config.JwtExperesIn)
	})

	linkDB := database.NewLink(db)
	linkHandler := handler.NewLinkHandler(linkDB)

	userRoute := r.Group("/users")
	userRoute.POST("/generate_token", userHandler.GetJWT)
	userRoute.GET("/", userHandler.GetManyUser)
	userRoute.POST("/", userHandler.CreateUser)
	userRoute.GET("/:uuid", userHandler.GetOneUser)
	userRoute.PATCH("/:uuid", userHandler.UpdateUser)
	userRoute.DELETE("/:uuid", userHandler.DeleteUser)

	linkRoute := r.Group("/links")
	linkRoute.Use(func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BEARER_SCHEMA)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token in authorization header"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			for key, val := range claims {
				if key == "sub" {
					ctx.Set("uuid", val)
				}
				fmt.Printf("Key: %v, value: %v\n", key, val)
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		ctx.Next()
	})
	linkRoute.GET("/", linkHandler.GetManyLink)
	linkRoute.POST("/", linkHandler.CreateLink)
	linkRoute.GET("/:hash", linkHandler.GetOneLink)
	linkRoute.PATCH("/:uuid", linkHandler.UpdateLink)
	linkRoute.DELETE("/:uuid", linkHandler.DeleteLink)

	r.Run(":8080")
}
