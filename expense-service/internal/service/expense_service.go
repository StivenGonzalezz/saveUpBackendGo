package service

import (
	"expense-service/internal/domain/model"
	"expense-service/internal/domain/ports"
	"time"
)

type ExpenseService struct {
	Repo ports.ExpenseRepository
}

func (s *ExpenseService) CreateExpense(expense *model.Expense) error {
	expense.IsExpense= true
	createdAt := time.Now()
	expense.CreatedAt = createdAt.Format("2006-01-02") 
	return s.Repo.CreateExpense(expense)
}

func (s *ExpenseService) DeleteExpense(id uint) error {
	return s.Repo.DeleteExpense(id)
}

func (s *ExpenseService) GetExpensesByFinanceId(financeId uint) ([]model.Expense, error) {
	return s.Repo.GetExpensesByFinanceId(financeId)
}

func (s *ExpenseService) CreateIncome(income *model.Expense) error {
	income.IsExpense= false
	createdAt := time.Now()
	income.CreatedAt = createdAt.Format("2006-01-02") 
	return s.Repo.CreateIncome(income)
}	

func (s *ExpenseService) DeleteIncome(id uint) error {
	return s.Repo.DeleteIncome(id)
}
