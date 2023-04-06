package entities

import (
	"github.com/google/uuid"
)

type Link struct {
	ID   uuid.UUID `json:"id"`
	Url  string    `json:"url"`
	Hash string    `json:"hash"`
}

func NewLink(url string, hash string) (*Link, error) {
	return &Link{
		ID:   uuid.New(),
		Url:  url,
		Hash: hash,
	}, nil
}
