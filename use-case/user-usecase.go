package usecase

import (
	"errors"
	"time"
	"travelnest/repository"
	"travelnest/utils"

	"github.com/golang-jwt/jwt/v4"
)

type UserUseCase struct {
    UserRepo repository.UserRepository
    JWTKey   []byte
}

func NewUserUseCase(repo repository.UserRepository, jwtKey []byte) UserUseCase {
		return UserUseCase{
				UserRepo: repo,
				JWTKey:   jwtKey,
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

    return token.SignedString(u.JWTKey)
}
