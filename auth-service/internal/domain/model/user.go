package model

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt string
}
