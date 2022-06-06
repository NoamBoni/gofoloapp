package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/NoamBoni/gofoloapp/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(ctx *gin.Context) {
	helpers.LoadEnv()
	cookie, err := ctx.Cookie("token")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  "please login to continue",
		})
		return
	}
	token, err := jwt.ParseWithClaims(
		cookie,
		&Claims{},
		getSecret,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  "please login to continue",
		})
		return
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  "please login to continue",
		})
		return
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  "authentication expired",
		})
		return
	}
	ctx.Set("role", claims.Role)
	ctx.Set("user-id", claims.User_id)
	ctx.Set("name", claims.Name)
	ctx.Next()
}

func getSecret(t *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}
