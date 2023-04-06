package database

import (
	"fmt"

	"github.com/genesis-crypto/shortner-go/internal/entities"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entities.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	fmt.Println("M  ", user)
	return &user, nil
}

func (u *User) FindMany(page, limit int) ([]entities.User, error) {
	var users []entities.User
	var err error

	if page != 0 && limit != 0 {
		err = u.DB.Limit(limit).Offset((page - 1) * limit).Find(&users).Error
	} else {
		err = u.DB.Find(&users).Error
	}

	return users, err
}
