package model

type Finance struct {
	ID     uint `gorm:"primarykey"`
	UserID uint
}
