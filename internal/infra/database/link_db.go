package database

import (
	"github.com/genesis-crypto/shortner-go/internal/entities"
	"gorm.io/gorm"
)

type Link struct {
	DB *gorm.DB
}

func NewLink(db *gorm.DB) *Link {
	return &Link{DB: db}
}

func (l *Link) Create(link *entities.Link) error {
	return l.DB.Create(link).Error
}

func (l *Link) FindByHash(hash string) (*entities.Link, error) {
	var link entities.Link
	err := l.DB.First(&link, "hash = ?", hash).Error
	return &link, err
}

func (l *Link) FindMany(page, limit int) ([]entities.Link, error) {
	var links []entities.Link
	var err error

	if page != 0 && limit != 0 {
		err = l.DB.Limit(limit).Offset((page - 1) * limit).Find(&links).Error

	} else {
		err = l.DB.Find(&links).Error
	}
	return links, err
}
