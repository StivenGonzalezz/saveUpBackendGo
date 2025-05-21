package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing Authorization header"})
			c.Abort()
			return
		}

		req, err := http.NewRequest("POST", "http://auth-service:8080/validate", nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error creating request"})
			c.Abort()
			return
		}
		req.Header.Set("Authorization", authHeader)

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil || res.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
