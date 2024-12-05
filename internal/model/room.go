package model

import (
	"time"
	"travelnest/internal/enums"

	"github.com/google/uuid"
)

type Room struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Number      int                 `json:"number"`
	Category    enums.RoomCategory  `json:"category"`
	Price       float64             `json:"price"`
	IsAvailable bool                `json:"is_available"`
	CreatedAt   time.Time           `json:"created_at" format:"2006-01-02T15:04:05Z07:00"`
	UpdatedAt   time.Time           `json:"updated_at" format:"2006-01-02T15:04:05Z07:00"`
	DeletedAt   *time.Time          `json:"deleted_at,omitempty" format:"2006-01-02T15:04:05Z07:00"`
}

func NewRoom(name string, number int, category enums.RoomCategory, price float64, isAvailable bool) Room {
	now := time.Now()
	return Room{
		ID:          uuid.New(),
		Name:        name,
		Number:      number,
		Category:    category,
		Price:       price,
		IsAvailable: isAvailable,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   nil,
	}
}
