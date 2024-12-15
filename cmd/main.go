package main

import (
	"log"
	"travelnest/database"
	"travelnest/handler"
	"travelnest/repository"
	"travelnest/router"
	usecase "travelnest/use-case"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	
	dbConnection, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	roomRepository := repository.NewRoomRepository(dbConnection)
	userRepository := repository.NewUserRepository(dbConnection)

	roomUseCase := usecase.NewRoomUseCase(roomRepository)
	userUseCase := usecase.NewUserUseCase(userRepository)

	roomHandler := handler.NewRoomHandler(roomUseCase)
	userHandler := handler.NewUserHandler(userUseCase)

	router.SetupRoutes(server, userHandler, roomHandler)

	log.Println("Servidor rodando na porta 8080...")
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}