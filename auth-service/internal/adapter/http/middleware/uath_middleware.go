package middleware

import(
	"auth-service/pkg/jwt"
)


func AuthMiddleware(c *gin.Context){
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	jwt.ValidateToken(tokenString)
}