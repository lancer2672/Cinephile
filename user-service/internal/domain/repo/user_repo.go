package repo

import entities "github.com/vudinhan2525/TaskFlow/internal/domain/entities"

type IUserRepository interface {
	FindByID(id int64) (*entities.User, error)
	CreateUser(email, hashedPassword, fullName, role string) (*entities.User, error)
}
