package model

type Expense struct {
	ID          uint `gorm:"primarykey"`
	FinanceId   uint
	Description string
	Amount      float64
	Icon        string
	IsExpense   bool
	CreatedAt   string
}
