package usecase

import (
	"github.com/vudinhan2525/TaskFlow/internal/domain/entities"
)

type UserUseCase interface {
	GetUserByID(id int64) (*entities.User, error)
	CreateUser(email, hashedPassword, fullName, role string) (*entities.User, error)
}
