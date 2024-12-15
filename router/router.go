package router

import (
	"os"

	"travelnest/handler"
	"travelnest/middleware"
	"travelnest/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine, userHandler handler.UserHandler, roomHandler handler.RoomHandler) {
	utils.LoadEnv()

	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	
	server.POST("/auth/login", userHandler.Login)
	server.GET("/rooms", middleware.RoleAuthorization(jwtKey, "admin"), roomHandler.GetRooms)
}
