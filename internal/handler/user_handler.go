package handler

import (
	"net/http"
	usecase "travelnest/internal/use-case"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) UserHandler {
	return UserHandler{
		UserUseCase: usecase,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    token, err := h.UserUseCase.Login(input.Username, input.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
