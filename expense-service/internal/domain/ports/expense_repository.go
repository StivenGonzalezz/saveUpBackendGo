package ports

import "expense-service/internal/domain/model"

type ExpenseRepository interface {
	CreateExpense(expense *model.expense) error
	DeleteExpense(id uint) error
	GetExpensesByFinanceId(financeId uint) ([]model.expense, error)
}
