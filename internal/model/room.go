package model

import (
	"travelnest/internal/enums"

	"github.com/google/uuid"
)

type Room struct {
	ID uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Number int    `json:"number"`
	Category enums.RoomCategory    `json:"category"`
	Price float64   `json:"price"`
	IsAvailable bool `json:"is_available"`
}


func NewRoom(name string, number int, category enums.RoomCategory, price float64, is_available bool) Room {
	return Room{
		ID: uuid.New(),
		Name: name,
		Number: number,
		Category: category,
		Price: price,
		IsAvailable: is_available,
	}
}