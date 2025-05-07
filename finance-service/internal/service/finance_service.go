package service

import (
	"finance-service/internal/domain/model"
	"finance-service/internal/domain/ports"
)

type FinanceService struct {
	Repo ports.FinanceRepository
}

func (s *FinanceService) CreateFinance(finance *model.Finance) error {
	return s.Repo.CreateFinance(finance)
}

func (s *FinanceService) GetFinanceByUserId(userId uint) (*model.Finance, error) {
	return s.Repo.GetFinanceByUserId(userId)
}
