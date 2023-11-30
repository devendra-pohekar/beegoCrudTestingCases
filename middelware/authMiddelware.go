package middelware

import (
	"net/http"

	// "github.com/astaxie/beego/context"
	"github.com/beego/beego/v2/server/web/context"

	"github.com/dgrijalva/jwt-go"
)

type Users struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type JwtClaim struct {
	Email  string `json:"user_email"`
	UserID int    `json:"user_id"`
	jwt.StandardClaims
}

var jwtKey = []byte("devendra_secretkey")

func Auth(ctx *context.Context) {
	tokenString := ctx.Input.Header("Authorization")
	if tokenString == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]string{"error": "Unauthorized"}, true, false)
		return
	}

	tokenString = tokenString[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]string{"error": "Invalid token"}, true, false)
		return
	}
	ctx.Input.SetData("LoginUserData", token.Claims.(jwt.MapClaims))
}
