package repository

import (
	"context"
	"database/sql"

	domain "github.com/vudinhan2525/TaskFlow/internal/domain/entities"
	"github.com/vudinhan2525/TaskFlow/internal/domain/repo"
	db "github.com/vudinhan2525/TaskFlow/internal/repository/db/sqlc"
)

type userRepository struct {
	queries *db.Queries
	db      *sql.DB
}

func NewUserRepository(dbstore *sql.DB) repo.IUserRepository {
	return &userRepository{
		queries: db.New(dbstore),
		db:      dbstore,
	}
}

func (r *userRepository) FindByID(id int64) (*domain.User, error) {
	user, err := r.queries.GetUserByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:       user.UserID,
		FullName: user.FullName,
		Email:    user.Email,
	}, nil
}
func (r *userRepository) CreateUser(email, hashedPassword, fullName, role string) (*domain.User, error) {

	user, err := r.queries.CreateUser(context.Background(), db.CreateUserParams{
		FullName:       fullName,
		Email:          email,
		HashedPassword: hashedPassword,
		Role:           db.UserRole(role),
	})
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:       user.UserID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     string(user.Role),
	}, nil
}
