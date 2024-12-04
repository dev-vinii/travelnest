package usecase

import (
	"travelnest/internal/model"
	"travelnest/internal/repository"
)

type RoomUseCase struct {
	reposotory repository.RoomRepository
}

func NewRoomUseCase(repo repository.RoomRepository) RoomUseCase {
	return RoomUseCase{
		reposotory: repo,
	}
}

func (r *RoomUseCase) GetRooms() ([]model.Room, error) {
	return r.reposotory.GetRooms()
}