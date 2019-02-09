package handlers

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"github.com/jinzhu/gorm"
)

type User struct {
	Email      string
	Nickname   string
	Password   string
	Authorized bool
	DateJoined time.Time
}

func Login(c echo.Context) error {

	var user User

	user.Email = c.QueryParam("email")
	user.Password = c.QueryParam("password")

	if user.Email != "giho" || user.Password != "zxcv1234!" {
		log.Info(user)
		return echo.ErrUnauthorized
	}


	token, err := CreateJwtToken(user.Email)
	if err != nil {
		log.Error("Error occur during to create JWT")
		return echo.ErrInternalServerError
	}

	jwtCookie := &http.Cookie{}
	jwtCookie.HttpOnly = true
	jwtCookie.Name = "JWTCookie"
	jwtCookie.Value = token
	jwtCookie.Expires = time.Now().Add(48 * time.Hour)

	c.SetCookie(jwtCookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "You were logged in!",
		"token":   token,
	})
}

//func SignUp(c echo.Context) error {
//	var user User
//
//	if err := c.Bind(user); err != nil {
//		log.Error("SigUp error")
//	}
//
//	//user := findUserByEmail(user.Email,nil)
//}

func findUserByEmail(email string, db gorm.DB) User {
	var user User

	db.Where("email = ?", email).First(&user)

	return user
}