package ports

import "expense-service/internal/domain/model"

type ExpenseRepository interface {
	CreateExpense(expense *model.Expense) error
	CreateIncome(expense *model.Expense) error
	DeleteExpense(id uint) error
	DeleteIncome(id uint) error
	GetExpensesByFinanceId(financeId uint) ([]model.Expense, error)
}
