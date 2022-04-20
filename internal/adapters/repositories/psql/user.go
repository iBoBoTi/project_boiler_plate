package psql

import (
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) ports.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetUser() (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) CreateUser() (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) UpdateUser() (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) DeleteUser() error {
	return nil
}
