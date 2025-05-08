package expense

type expense struct {
	ID        uint `gorm:"primarykey"`
	financeId uint
	Description string
	Amount float64
	Icon string
	CreatedAt string
}