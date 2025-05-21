package service

import (
	"goal-service/internal/domain/model"
	"goal-service/internal/domain/ports"
)

type GoalService struct {
	Repo ports.GoalRepository
}

func (s *GoalService) CreateGoal(goal *model.Goal) error {
	goal.CurrentAmount = 0
	return s.Repo.CreateGoal(goal)
}

func (s *GoalService) DeleteGoal(goalId uint) error {
	return s.Repo.DeleteGoal(goalId)
}

func (s *GoalService) GetGoalsByUserId(userId uint) ([]model.Goal, error) {
	return s.Repo.GetGoalsByUserId(userId)
}

func (s *GoalService) UpdateGoalName(goalId uint, name string) error {
	return s.Repo.UpdateGoalName(goalId, name)
}

func (s *GoalService) UpdateGoalCurrentAmount(goalId uint, amount uint) error {
	return s.Repo.UpdateGoalCurrentAmount(goalId, amount)
}
