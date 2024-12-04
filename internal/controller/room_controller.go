package controller

import (
	usecase "travelnest/internal/use-case"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomUseCase usecase.RoomUseCase
}

func NewRoomController(usecase usecase.RoomUseCase) RoomController {
	return RoomController{
		roomUseCase: usecase,
	}
}

func (r *RoomController) GetRooms(c *gin.Context) {
	rooms, err := r.roomUseCase.GetRooms()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, rooms)
}