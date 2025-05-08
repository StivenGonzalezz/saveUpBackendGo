package service

import (
	"expense-service/internal/domain/ports"
	"expense-service/internal/domain/model"
)

type ExpenseService struct {
	Repo ports.ExpenseRepository
}

func (s *ExpenseService) CreateExpense(expense *model.expense) error {
	return s.Repo.CreateExpense(expense)
}

func (s *ExpenseService) DeleteExpense(id uint) error {
	return s.Repo.DeleteExpense(id)
}

func (s *ExpenseService) GetExpensesByFinanceId(financeId uint) ([]model.expense, error) {
	return s.Repo.GetExpensesByFinanceId(financeId)
}