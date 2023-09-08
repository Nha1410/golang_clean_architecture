package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/team2/real_api/config"
)

func SetupJwtKey() []byte {
	conf := config.LoadConfig()
	return []byte(conf.SECRET.JwtKey)
}

var jwtKey []byte

func init() {
	jwtKey = SetupJwtKey()
}

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}

	expirationTime, expOk := claims["exp"].(float64)
	if !expOk {
		return 0, errors.New("Token does not contain valid expiration time")
	}

	expUnix := int64(expirationTime)
	if time.Now().Unix() > expUnix {
		return 0, errors.New("Token has expired")
	}

	userID := int(claims["user_id"].(float64))
	return userID, nil
}
