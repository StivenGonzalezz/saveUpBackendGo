package repository

import (
	"expense-service/internal/domain/model"
	"expense-service/internal/domain/ports"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo() ports.ExpenseRepository {
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

	db.AutoMigrate(&model.Expense{})

	return &PostgresRepo{db: db}
}

func (p *PostgresRepo) CreateExpense(expense *model.Expense) error {
	return p.db.Create(expense).Error
}

func (p *PostgresRepo) DeleteExpense(id uint) error {
	var expense model.Expense
	if err := p.db.First(&expense, id).Error; err != nil {
		return err
	}
	return p.db.Delete(&expense).Error
}

func (p *PostgresRepo) GetExpensesByFinanceId(financeId uint) ([]model.Expense, error) {
	var expenses []model.Expense
	if err := p.db.Where("finance_id = ?", financeId).Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}


func (p *PostgresRepo) CreateIncome(income *model.Expense) error {
	return p.db.Create(income).Error
}

func (p *PostgresRepo) DeleteIncome(id uint) error {
	var income model.Expense
	if err := p.db.First(&income, id).Error; err != nil {
		return err
	}
	return p.db.Delete(&income).Error
}