package repository

import (
	"goal-service/internal/domain/model"
	"goal-service/internal/domain/ports"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo() ports.GoalRepository {
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
	
	db.AutoMigrate(&model.Goal{})

	return &PostgresRepo{db: db}
}

func (p *PostgresRepo) CreateGoal(goal *model.Goal) error {
	return p.db.Create(goal).Error
}

func (p *PostgresRepo) GetGoalsByUserId(userId uint) ([]model.Goal, error) {
	var goals []model.Goal
	if err := p.db.Where("user_id = ?", userId).Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}

func (p *PostgresRepo) GetGoalByUserId(userId uint) (*model.Goal, error) {
	var goal model.Goal
	if err := p.db.Where("user_id = ?", userId).First(&goal).Error; err != nil {
		return nil, err
	}
	return &goal, nil
}

func (p *PostgresRepo) UpdateGoalName(goalId uint, name string) error {
	return p.db.Model(&model.Goal{}).Where("id = ?", goalId).Update("name", name).Error
}

func (p *PostgresRepo) DeleteGoal(goalId uint) error {
	return p.db.Delete(&model.Goal{}, goalId).Error
}

func (p *PostgresRepo) UpdateGoalCurrentAmount(goalId uint, newAmount uint) error {
	return p.db.Model(&model.Goal{}).Where("id = ?", goalId).Update("current_amount", newAmount).Error
}


