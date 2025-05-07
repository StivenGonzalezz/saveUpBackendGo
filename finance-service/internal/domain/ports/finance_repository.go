package ports

import "finance-service/internal/domain/model"

type FinanceRepository interface {
	CreateFinance(finance *model.Finance) error
	GetFinanceByUserId(userId uint) (*model.Finance, error)
}