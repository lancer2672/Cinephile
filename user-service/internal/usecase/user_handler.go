package usecase

import (
	"github.com/vudinhan2525/TaskFlow/internal/domain/entities"
	"github.com/vudinhan2525/TaskFlow/internal/domain/repo"
)

type UserUseCaseImpl struct {
	userRepo repo.IUserRepository
}

func NewUserUseCase(repo repo.IUserRepository) UserUseCase {
	return &UserUseCaseImpl{
		userRepo: repo,
	}
}

func (u *UserUseCaseImpl) GetUserByID(id int64) (*entities.User, error) {
	return u.userRepo.FindByID(id)
}
func (u *UserUseCaseImpl) CreateUser(email, hashedPassword, fullName, role string) (*entities.User, error) {
	return u.userRepo.CreateUser(email, hashedPassword, fullName, role)
}
