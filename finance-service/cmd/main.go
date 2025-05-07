package main

import (
	"finance-service/internal/adapter/http"
	"finance-service/internal/adapter/repository"
	"finance-service/internal/service"
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
	// Crear repositorio e inyectar dependencias
	repo := repository.NewPostgresRepo()
	finance := &service.FinanceService{Repo: repo}

	// Configurar y arrancar el servidor
	router := gin.Default()
	http.SetupRoutes(router, finance)
	router.Run(":8081")

}
