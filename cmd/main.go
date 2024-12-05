package main

import (
	"travelnest/internal/controller"
	"travelnest/internal/database"
	"travelnest/internal/repository"
	usecase "travelnest/internal/use-case"

	"github.com/gin-gonic/gin"
)

func main() {
		server := gin.Default()

		dbConnection, err := database.ConnectDatabase()
		
		if(err != nil) {
			panic(err)
		}
		
		RoomRepository := repository.NewRoomRepository(dbConnection)
		RoomUseCase := usecase.NewRoomUseCase(RoomRepository)
		RoomController := controller.NewRoomController(RoomUseCase)
		
		server.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		
		server.GET("/rooms", RoomController.GetRooms)

		
		server.Run(":8080")
}