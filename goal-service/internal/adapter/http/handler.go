package http

import (
	"goal-service/internal/domain/model"
	"goal-service/internal/service"
	"net/http"
	"strconv"

	"goal-service/internal/adapter/http/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, goalService *service.GoalService) {
	router.POST("/create", middleware.AuthMiddleware(), func(c *gin.Context){
		var req model.Goal
		if err := c.ShouldBindJSON(&req);err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid request"})
			return
		}
		err := goalService.CreateGoal(&req)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error, could not create goal", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message":"Goal created successfully", "goal":&req})
	})

	router.GET("/getGoals/:userId", middleware.AuthMiddleware(), func(c *gin.Context){
		userId := c.Param("userId")
		userIdInt, err := strconv.Atoi(userId)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid user ID"})
			return
		}
		goals, err := goalService.GetGoalsByUserId(uint(userIdInt))
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error, could not get goals", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message":"Goals retrieved successfully", "goals":goals})
	})
	
	router.DELETE("/delete/:id", middleware.AuthMiddleware(), func(c *gin.Context){
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid goal ID"})
			return
		}
		err = goalService.DeleteGoal(uint(idInt))
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error, could not delete goal", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message":"Goal deleted successfully"})
	})

	router.PUT("/updateName", middleware.AuthMiddleware(), func(c *gin.Context){
		var req model.Goal
		if err := c.ShouldBindJSON(&req);err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid request"})
			return
		}
		err := goalService.UpdateGoalName(uint(req.ID), req.Name)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error, could not update goal name", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message":"Goal name updated successfully", "goal":&req})
	})

	router.PUT("/updateAmount", middleware.AuthMiddleware(), func(c *gin.Context){
		var req model.Goal
		if err := c.ShouldBindJSON(&req);err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid request"})
			return
		}
		err := goalService.UpdateGoalCurrentAmount(uint(req.ID), uint(req.CurrentAmount))
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error, could not update goal amount", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message":"Goal amount updated successfully", "goal":&req})
	})
}
