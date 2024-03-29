package service

import (
	"fmt"
	"goTestProj/config"
	"goTestProj/models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(config.Env.GetString("jwt.secret"))

// var jwtSecret = base64.URLEncoding.EncodeToString(make([]byte, 10))
// var jwtSecret = []byte(setting.jwtSecret)

func GenerateToken(username string, password string) (string, error) {
	nowTime := time.Now()
	// expireTCount := config.Env.GetInt("jwt.expiretime")
	expireTime := nowTime.Add(3 * time.Hour)

	claims := models.Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	//  帶入演算法及claims, 產出token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	if err != nil {

		fmt.Println("=== orig jwtSecret : ", config.Env.GetString("jwt.secret"))
		fmt.Println("=== jwtSecret : ", jwtSecret)
		fmt.Println("=== SignedString err : ", err)
	}

	return token, err
}

func ParseToken(token string) (*jwt.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token,
		&models.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)

	if tokenClaims != nil {
		if _, ok := tokenClaims.Claims.(*models.Claims); ok && tokenClaims.Valid {
			return &tokenClaims.Claims, nil
		}
	}

	return nil, err
}

func MethodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== MethodHandler ===")
	if r.Method != http.MethodGet {
		fmt.Println("=== not Get ===")
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := authorizationHeader[len("Bearer "):]
		_, err := ParseToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func withJWTVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			tokenString := extractTokenFromHeader(c.Request)
			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
				})
			}
		}
	}
}

func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	splitToken := strings.Split(authHeader, "Bearer ")

	if len(splitToken) != 2 {
		return ""
	}

	return splitToken[1]
}

func protectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "your are accessed here",
	})
}
