package usecase

import (
	"errors"
	"log"
	"os"
	"time"
	"travelnest/repository"
	"travelnest/utils"

	"github.com/golang-jwt/jwt/v4"
)

type UserUseCase struct {
    UserRepo repository.UserRepository
    JWTKey   string
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
    utils.LoadEnv()
		return UserUseCase{
				UserRepo: repo,
				JWTKey:   (os.Getenv("JWT_SECRET_KEY")),
		}
}

func (u *UserUseCase) Login(username, password string) (string, error) {
    user, err := u.UserRepo.GetByUsername(username)	
    if err != nil || !utils.CheckPasswordHash(password, user.Password) {
        return "", errors.New("invalid username or password")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":  user.ID,
        "role": user.Role,
        "exp": time.Now().Add(30 * time.Minute).Unix(),
        
    })
    log.Println("KEY: ", u.JWTKey)
    jwtKeyBytes := []byte(u.JWTKey)

    return token.SignedString(jwtKeyBytes)
}
