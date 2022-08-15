package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/NoamBoni/gofoloapp/helpers"
	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type ExtraClaims struct {
	User_id uint
	Role    string
	Name    string
}

type Claims struct {
	jwt.RegisteredClaims
	User_id uint   `json:"user-id"`
	Role    string `json:"role"`
	Name    string `json:"name"`
}

func Login(ctx *gin.Context) {
	var loginUser models.User
	var usersList []models.User
	if err := ctx.ShouldBindBodyWith(&loginUser, binding.JSON); err != nil {
		ReturnError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := models.GetUsersByName(loginUser.Name, &usersList); err != nil {
		ReturnError(ctx, http.StatusBadRequest, errors.New("invalid name or password"))
		return
	}
	for _, user := range usersList {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err == nil {
			claims := ExtraClaims{
				Name:    user.Name,
				Role:    user.Role,
				User_id: user.Id,
			}
			token, _ := GenerateJWT(&claims)
			helpers.LoadEnv()
			secure := os.Getenv("SERVER_STATE") != "development"
			ctx.SetCookie("token", token, 3600*3, "/", "localhost", secure, secure)
			if user.Role == models.Role.T {
				user.Password = ""
				ctx.JSON(http.StatusOK, gin.H{
					"status": "successful authentication",
					"data":   user,
				})
				return
			} else {
				patient := models.Patient{
					User_id: user.Id,
					Role:    user.Role,
					Name:    user.Name,
				}
				_ = patient.Select()
				ctx.JSON(http.StatusOK, gin.H{
					"status": "successful authentication",
					"data":   patient,
				})
				return
			}
		}
	}
	ReturnError(ctx, http.StatusBadRequest, errors.New("invalid name or password"))
}

func GenerateJWT(ex *ExtraClaims) (string, error) {
	helpers.LoadEnv()
	JWTsecret := []byte(os.Getenv("JWT_SECRET"))
	TokenDuration := time.Hour * 3
	c := Claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenDuration)),
			Issuer:    "gofoloapp",
		},
		ex.User_id,
		ex.Role,
		ex.Name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(JWTsecret)
}
