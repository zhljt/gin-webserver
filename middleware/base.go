package middleware

import "github.com/gin-gonic/gin"

func SetRequestID() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("ReqId", "1234")
		c.Next()
	}

}
