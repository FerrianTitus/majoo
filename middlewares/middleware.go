package middlewares

import (
	"github.com/majoo/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	Config = middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
)

type jwtCustomClaims struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"user_name"`
	jwt.StandardClaims
}

func CreateToken(User models.Users) (string, error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		1,
		"admin1",
		"admin1",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "-", err
	}

	return t, nil
}

func GetDataFromToken(c echo.Context) (models.DataFromToken, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)

	return models.DataFromToken{
		ID:       int(claims.ID),
		Name:     claims.Name,
		Username: claims.Username,
	}, nil
}
