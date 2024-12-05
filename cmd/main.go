package main

import (
	"log"
	"travelnest/internal/database"
	"travelnest/internal/handler"
	"travelnest/internal/repository"
	usecase "travelnest/internal/use-case"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o servidor Gin
	server := gin.Default()

	// Conexão com o banco de dados
	dbConnection, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Configuração dos Repositórios
	roomRepository := repository.NewRoomRepository(dbConnection)
	userRepository := repository.NewUserRepository(dbConnection)

	// Configuração dos Casos de Uso
	roomUseCase := usecase.NewRoomUseCase(roomRepository)
	userUseCase := usecase.NewUserUseCase(userRepository, []byte("secret"))

	// Configuração dos Handlers
	roomHandler := handler.NewRoomHandler(roomUseCase)
	userHandler := handler.NewUserHandler(userUseCase)

	setupRoutes(server, userHandler, roomHandler)

	log.Println("Servidor rodando na porta 8080...")
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

func setupRoutes(server *gin.Engine, userHandler handler.UserHandler, roomHandler handler.RoomHandler) {
	server.POST("/login", userHandler.Login)
	server.GET("/rooms", roomHandler.GetRooms)
}
