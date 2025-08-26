package database

import "apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error {
		FindBtEmail(email string) (*entity.User, error)
	}
}
