package database

import (
	"github.com/genesis-crypto/shortner-go/internal/entities"
)

type UserInterface interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindMany(page, limit int) ([]entities.User, error)
}

type LinkInterface interface {
	Create(product *entities.Link) error
	FindByHash(hash string) (*entities.Link, error)
	FindMany(page, limit int) ([]entities.Link, error)
}
