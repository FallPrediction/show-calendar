package utils

import (
	"os"
	custom_errors "souflair/errors"
	"souflair/initialize"
	"souflair/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct{}

type UserData struct {
	UserId      uint32 `json:"user_id"`
	Name        string `json:"user_name"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	RefreshTime int64  `json:"refresh_time"`
}

type userClaims struct {
	UserData
	jwt.RegisteredClaims
}

func (j *Jwt) CreateUserToken(user *models.User) (string, time.Time, error) {
	logger := initialize.NewLogger()
	key := os.Getenv("APP_KEY")
	app := os.Getenv("APP_NAME")
	lifeTime, err := strconv.Atoi(os.Getenv("JWT_LIFETIME"))
	if err != nil {
		logger.Error("The lifetime in JWT is invalid. ", err)
		lifeTime = 7200
	}
	refreshTime, err := strconv.Atoi(os.Getenv("JWT_REFRESHTIME"))
	if err != nil {
		logger.Error("The refreshtime in JWT is invalid. ", err)
		refreshTime = 10800
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(lifeTime))
	claims := userClaims{
		UserData{
			UserId:      user.Id,
			Name:        user.Name,
			Avatar:      user.Avatar,
			Email:       user.Email,
			RefreshTime: time.Now().Add(time.Second * time.Duration(refreshTime)).Unix(),
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    app,
			Audience:  []string{app},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		logger.Error("Generate a JWT token failed. ", err)
	}
	return tokenString, expiresAt, err
}

func (j *Jwt) ParseUserToken(tokenString string) (*userClaims, error) {
	logger := initialize.NewLogger()
	key := os.Getenv("APP_KEY")
	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		logger.Error("parse token failed. ", err)
		return nil, err
	} else if claims, ok := token.Claims.(*userClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, custom_errors.ErrInvalidToken
	}
}
