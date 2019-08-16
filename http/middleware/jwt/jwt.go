package jwt

import (
	"ginapi/http/handler"
	"ginapi/pkg/exception"
	"ginapi/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT this is middleware JWT
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header["Authorization"]
		if len(token) == 0 {
			handler.SendJSONFail(c, exception.ErrotAuthCheckTokenFail, token)
			c.Abort()
			return
		}
		if token[0] == "" {
			handler.SendJSONFail(c, exception.ErrotAuthCheckTokenFail, token)
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token[0])
		if err != nil {
			handler.SendJSONFail(c, exception.ErrotAuthCheckTokenFail, err)
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			handler.SendJSONFail(c, exception.ErrotAuthCheckTokenTimeout, err)
			c.Abort()
			return
		}
		c.Next()
	}
}
