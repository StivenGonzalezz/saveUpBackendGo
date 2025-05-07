package service

import (
	"auth-service/internal/domain/model"
	"auth-service/internal/domain/ports"
	"auth-service/pkg/hash"
	"auth-service/pkg/jwt"
	"errors"
	"fmt"
)

type AuthService struct {
	Repo ports.UserRepository
}

func (s *AuthService) Register(user *model.User) error {
	hashPassword, _ := hash.HashPassword(user.Password)
	user.Password = hashPassword  
	return s.Repo.Save(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	comparedSucces := hash.ComparePassword(user.Password, password)
	fmt.Println(comparedSucces)
	if comparedSucces != nil {
		return "", errors.New("invalid password")
	}

	token, err := jwt.GenerateToken(int(user.ID))

	return token, err
}

func (s *AuthService) Update(email, password string) error {
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return err
	}
	hashPassword, _ := hash.HashPassword(password)
	user.Password = hashPassword  
	return s.Repo.Update(user)
}

