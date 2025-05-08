package http

func SetupRoutes(router *gin.Engine, expenseService *service.ExpenseService) {
	router.POST("/create", func(c *gin.Context) {
		var req model.expense
		if err := c.ShouldBindJSON(&req); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid request"})
			return
		}

		err := expenseService.CreateExpense(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error, could not create expense", "error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message":"Expense created successfully", "expense": &req})
	})

	router.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := expenseService.DeleteExpense(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error, could not delete expense", "error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message":"Expense deleted successfully"})
	})

	router.GET("/getExpenses/:financeId", func(c *gin.Context) {
		financeId := c.Param("financeId")
		expenses, err := expenseService.GetExpensesByFinanceId(uint(financeId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error, could not get expenses", "error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message":"Expenses retrieved successfully", "expenses": expenses})
	})
}