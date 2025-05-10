package main

import (
	"expense-service/internal/adapter/http"
	"expense-service/internal/adapter/repository"
	"expense-service/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables del sistema si existen.")
	}
}

func main() {
	//Crear repositorio e inyectar dependendencias
	repo := repository.NewPostgresRepo()
	expense := &service.ExpenseService{Repo: repo}

	//Configurar y arrancar el servidor
	router := gin.Default()
	http.SetupRoutes(router, expense)
	router.Run(":8082")
}
 