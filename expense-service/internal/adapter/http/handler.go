package http

import (
	"expense-service/internal/domain/model"
	"expense-service/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"expense-service/internal/adapter/http/middleware"
)

func SetupRoutes(router *gin.Engine, expenseService *service.ExpenseService) {
	router.POST("/createExpense", middleware.AuthMiddleware(), func(c *gin.Context) {

		var req model.Expense
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err := expenseService.CreateExpense(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, could not create expense", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Expense created successfully", "expense": &req})
	})

	router.DELETE("/deleteExpense/:id", middleware.AuthMiddleware(), func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
			return
		}
		err = expenseService.DeleteExpense(uint(idInt))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, could not delete expense", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
	})

	router.GET("/getExpenses/:financeId", middleware.AuthMiddleware(), func(c *gin.Context) {
		financeId := c.Param("financeId")
		financeIdInt, err := strconv.Atoi(financeId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid finance ID"})
			return
		}
		expenses, err := expenseService.GetExpensesByFinanceId(uint(financeIdInt))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, could not get expenses", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Expenses retrieved successfully", "expenses": expenses})
	})

	router.POST("/createIncome", middleware.AuthMiddleware(), func(c *gin.Context) {
		var req model.Expense
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err := expenseService.CreateIncome(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, could not create income", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Income created successfully", "income": &req})
	})

	router.DELETE("/deleteIncome/:id", middleware.AuthMiddleware(), func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
			return
		}
		err = expenseService.DeleteIncome(uint(idInt))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, could not delete income", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Income deleted successfully"})
	})
}
