package model

import (
	"travelnest/internal/enums"

	"github.com/google/uuid"
)

type Room struct {
	ID uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Description string    `json:"description"`
	Status enums.Status    `json:"status"`
	Price float64   `json:"price"`
}


func NewRoom(name string, description string, status enums.Status) *Room {
	return &Room{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Status:      status,
	}
}