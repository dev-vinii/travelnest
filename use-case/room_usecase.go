package usecase

import (
	"travelnest/model"
	"travelnest/repository"
)

type RoomUseCase struct {
	repository repository.RoomRepository
}

func NewRoomUseCase(repo repository.RoomRepository) RoomUseCase {
	return RoomUseCase{
		repository: repo,
	}
}

func (r *RoomUseCase) GetRooms() ([]model.Room, error) {
	return r.repository.GetRooms()
}