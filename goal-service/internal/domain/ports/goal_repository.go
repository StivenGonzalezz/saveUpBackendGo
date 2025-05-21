package ports

import "goal-service/internal/domain/model"

type GoalRepository interface {
	CreateGoal(goal *model.Goal) error
	DeleteGoal(goalId uint) error
	GetGoalsByUserId(userId uint) ([]model.Goal, error)
	UpdateGoalName(goalId uint, name string) error
	UpdateGoalCurrentAmount(goalId uint, amount uint) error
}