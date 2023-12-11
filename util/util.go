package util

import (
	"time"

	"github.com/AlfrinP/point_calculator/config"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(id uint, role string) (string, error) {
	tokenByte := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)
	config, _ := config.LoadConfig(".")
	claims["user_id"] = id
	claims["role"] = role
	claims["exp"] = now.Add(config.JwtExpiresIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	return tokenByte.SignedString([]byte(config.JwtSecret))

}
