package ports

import "auth-service/internal/domain/model"

type UserRepository interface {
	Save(user *model.User) error
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
}