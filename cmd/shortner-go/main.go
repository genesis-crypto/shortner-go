package main

import (
	entities "github.com/genesis-crypto/shortner-go/internal/entities"
	handler "github.com/genesis-crypto/shortner-go/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/shortner-go"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.User{})

	r := gin.Default()

	userHandler := handler.UserHandler{}
	linkHandler := handler.LinkHandler{}

	r.GET("/users", userHandler.GetManyUser)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:uuid", userHandler.GetOneUser)
	r.PATCH("/users/:uuid", userHandler.UpdateUser)
	r.DELETE("/users/:uuid", userHandler.DeleteUser)
	r.GET("/links", linkHandler.GetManyLink)
	r.POST("/links", linkHandler.CreateLink)
	r.GET("/links/:uuid", linkHandler.GetOneLink)
	r.PATCH("/links/:uuid", linkHandler.UpdateLink)
	r.DELETE("/links/:uuid", linkHandler.DeleteLink)

	r.Run(":3000")
}
