package main

import (
	"goal-service/internal/adapter/http"
	"goal-service/internal/adapter/repository"
	"goal-service/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables del sistema si existen.")
	}
}

func main(){
	repo := repository.NewPostgresRepo()
	goal := &service.GoalService{Repo: repo}

	router := gin.Default()
	http.SetupRoutes(router, goal)
	router.Run(":8083")
}