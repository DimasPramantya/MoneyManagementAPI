package middleware

import (
	"fmt"
	"strings"
	"time"

	. "github.com/dimas-pramantya/money-management/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type Claims struct {
	jwt.StandardClaims
}

var secret = viper.GetString("SECRET_KEY")

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := GetJwtTokenFromHeader(c)
		if err != nil {
			fmt.Println("Error getting token from header:", err)
			err := UnauthorizedError("Unauthorized", nil)
			c.Error(err)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println("Unexpected signing method:", token.Header["alg"])
				err := UnauthorizedError("Unauthorized", nil)
				c.Error(err)
				c.Abort()
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("Error parsing token:", err)
			err := UnauthorizedError("Unauthorized", nil)
			c.Error(err)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			fmt.Println("Invalid token claims")
			err := UnauthorizedError("Unauthorized", nil)
			c.Error(err)
			c.Abort()
			return
		}

		// Set auth context data
		c.Set("user_id", claims["user_id"])
		c.Set("username", claims["username"])

		c.Next()
	}
}

func GetJwtTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", UnauthorizedError("authorization header missing", nil)
	}

	// Expecting format: Bearer <token>
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", UnauthorizedError("authorization header format must be Bearer {token}", nil)
	}

	return parts[1], nil
}

//TOKEN EXPIRED IN 1 DAY
func GenerateJwtToken(username string, userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), 
		"iat":      time.Now().Unix(),                    
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}