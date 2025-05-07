package http

import (
	"finance-service/internal/domain/model"
	"finance-service/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, financeService *service.FinanceService) {
	router.POST("/create", func(c *gin.Context) {
		var req model.Finance
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err := financeService.CreateFinance(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, could not create finance", "error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Finance created successfully", "finance": &req})
	})

	router.GET("/getFinance/:userId", func(c *gin.Context) {
		userId := c.Param("userId")
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
			return
		}
		finance, err := financeService.GetFinanceByUserId(uint(userIdInt))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, could not get finance", "error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Finance retrieved successfully", "finance": finance})
	})
}
