package middlewares

import "github.com/gin-gonic/gin"

// Authenticate waa dummy middleware hadda, kaliya wuxuu pass gareeyaa request
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Haddii aad rabto auth logic, halkan geli
		c.Next()
	}
}
