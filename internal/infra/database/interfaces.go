package database

import "github.com/m3k3r1/go-api-template/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
