package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

type JwtClaims struct {
	Email string
	jwt.StandardClaims
}

func CreateJwtToken(email string) (string, error) {
	claims := JwtClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("secretKey"))


	if err != nil {
		return "", err
	}
	return token, nil
}

func MainJwt(c echo.Context) error {
	user := c.Get("user")

	token := user.(*jwt.Token)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}
	log.Print("User email : ", claims["Email"].(string))

	return c.String(http.StatusOK, "you are on the top secret jwt page!")
}