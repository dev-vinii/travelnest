package handler

import (
	usecase "travelnest/use-case"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	roomUseCase usecase.RoomUseCase
}

func NewRoomHandler(usecase usecase.RoomUseCase) RoomHandler {
	return RoomHandler{
		roomUseCase: usecase,
	}
}

func (r *RoomHandler) GetRooms(c *gin.Context) {
	rooms, err := r.roomUseCase.GetRooms()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, rooms)
}