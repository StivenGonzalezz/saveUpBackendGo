package model

type Goal struct {
	ID     uint `gorm:"primarykey"`
	UserID uint
	Name   string
	TargetAmount uint
	CurrentAmount uint
	Date   string
}