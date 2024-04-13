package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var KEY = []byte("TH1S_1S_MY_SUPER_S3CR3T_K3Y_1337")

func register(c *gin.Context) {
	c.String(200, "WOW")
}

func login(c *gin.Context) {
	c.String(200, "WOW")
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	api := r.Group("api")
	{
		v1 := api.Group("v1")
		{
			v1.POST("register", register)
			v1.POST("login", login)
		}
	}
	r.GET("/", register)
	r.Run(":3000")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recover() != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "Токена нет"})
				return
			}
		}()

		authorization := c.Request.Header["Authorization"][0]

		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "Токена нет"})
			return
		}

		splittedAuthorization := strings.Split(authorization, " ")
		if splittedAuthorization[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "Bearer-токена нет"})
			return
		}

		tokenString := splittedAuthorization[1]

		auth := &models.JWT{}

		token, err := jwt.ParseWithClaims(tokenString, auth, func(token *jwt.Token) (any, error) {
			return KEY, nil
		})
		if !token.Valid || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "Токен не валиден"})
			return
		}

		c.Set("userId", auth.UserId)
		c.Next()
	}
}
