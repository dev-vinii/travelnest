package repository

import (
	"travelnest/internal/model"

	"gorm.io/gorm"
)

type RoomRepository struct {
	connection *gorm.DB
}

func NewRoomRepository(connection *gorm.DB) RoomRepository {
	return RoomRepository{
		connection: connection,
	}
}

func (rr *RoomRepository) GetRooms() ([]model.Room, error) {
	var rooms []model.Room
	
	if err := rr.connection.Where("deleted_at IS NULL").Find(&rooms).Error; err != nil {
		return nil, err
	}
	
	return rooms, nil
}