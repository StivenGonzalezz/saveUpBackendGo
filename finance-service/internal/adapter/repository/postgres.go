package repository

import (
	"finance-service/internal/domain/model"
	"finance-service/internal/domain/ports"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo() ports.FinanceRepository {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}
	
	db.AutoMigrate(&model.Finance{})

	return &PostgresRepo{db: db}
}

func (p *PostgresRepo) CreateFinance(finance *model.Finance) error {
	return p.db.Create(finance).Error
}

func (p *PostgresRepo) GetFinanceByUserId(userId uint) (*model.Finance, error) {
	var finance model.Finance
	if err := p.db.Where("user_id = ?", userId).First(&finance).Error; err != nil {
		return nil, err
	}
	return &finance, nil
}
