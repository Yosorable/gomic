package middleware

import (
	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/model"
	"github.com/Yosorable/gomic/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, path := range global.SKIP_AUTH_PATH {
			if path == c.FullPath() {
				c.Next()
				return
			}
		}

		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			response.FailWithMessage("Unauthorized", c)
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.CONFIG.Secret), nil
		})

		if err != nil {
			response.FailWithMessage("Unauthorized", c)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*model.JWTClaims); ok && token.Valid {
			c.Set("user", claims)
			c.Next()
		} else {
			response.FailWithMessage("Unauthorized", c)
			c.Abort()
			return
		}
	}
}
