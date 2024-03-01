package service

import (
	"goTestProj/config"
	"goTestProj/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(config.Env.Get("jwt.secret"))

func GenerateToken(username string , password string) (string, error) {
	nowTime := time.Now()
	// expireTCount := config.Env.GetInt("jwt.expiretime")
	expireTime := nowTime.Add( 3 * time.Hour)

	claims := models.Claims {
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt : expireTime.Unix(),
			Issuer: "gin-blog",
		},
	}

//  帶入演算法及claims, 產出token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := tokenClaims.SignedString()
}	