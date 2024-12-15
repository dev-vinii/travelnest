package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	ErrMissingAuthHeader   = "missing Authorization header"
	ErrInvalidTokenFormat  = "invalid token format"
	ErrTokenParseError     = "unauthorized: invalid token"
	ErrInvalidToken        = "unauthorized: token is invalid"
	ErrInvalidClaims       = "forbidden: invalid claims"
	ErrRoleNotFound        = "forbidden: role not found"
	ErrAccessDenied        = "access denied"
)

func RoleAuthorization(jwtKey []byte, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println(ErrMissingAuthHeader)
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrMissingAuthHeader})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			log.Printf("%s: %s", ErrInvalidTokenFormat, authHeader)
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidTokenFormat})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("Unexpected signing method: %v", token.Method)
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtKey, nil
		})

		if err != nil {
			log.Printf("%s: %v", ErrTokenParseError, err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrTokenParseError})
			c.Abort()
			return
		}

		if !token.Valid {
			log.Println(ErrInvalidToken)
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidToken})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println(ErrInvalidClaims)
			c.JSON(http.StatusForbidden, gin.H{"error": ErrInvalidClaims})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role == "" {
			log.Println(ErrRoleNotFound)
			c.JSON(http.StatusForbidden, gin.H{"error": ErrRoleNotFound})
			c.Abort()
			return
		}

		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				log.Printf("User authorized with role: %s", role)
				c.Next()
				return
			}
		}

		log.Printf("%s for role: %s", ErrAccessDenied, role)
		c.JSON(http.StatusForbidden, gin.H{"error": ErrAccessDenied})
		c.Abort()
	}
}
