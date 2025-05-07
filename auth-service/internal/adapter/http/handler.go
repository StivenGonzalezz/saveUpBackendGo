package http

import (
	"auth-service/internal/domain/model"
	"auth-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authService *service.AuthService) {
	router.POST("/register", func(c *gin.Context){
		var req model.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err := authService.Register(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, could not register user", "error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": &req})
	})

	router.POST("/login", func(c *gin.Context){
		var req model.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}
		token, err := authService.Login(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, could not login user", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
	})

	router.PUT("/update", func(c *gin.Context){
		var req model.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err := authService.Update(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, could not update user", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": &req})
	})
}