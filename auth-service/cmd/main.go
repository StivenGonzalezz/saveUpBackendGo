package main

import (
	"auth-service/internal/adapter/http"
	"auth-service/internal/adapter/repository"
	"auth-service/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables del sistema si existen.")
	}

}

func main() {
	// Crear repositorio e inyectar dependencias
	repo := repository.NewPostgresRepo()
	auth := &service.AuthService{Repo: repo}

	// Configurar y arrancar el servidor
	router := gin.Default()
	http.SetupRoutes(router, auth)
	router.Run(":8080")
}
