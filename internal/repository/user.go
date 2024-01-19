package repository

import (
	"context"
	"database/sql"

	"user-app/internal/dto/input"
	"user-app/internal/model"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	CreateUser(context.Context, input.CreateUserRequest) (model.User, error)
	UpdateUser(context.Context, uuid.UUID, input.UpdateUserRequest) (model.User, error)
}

type UserRepository struct {
	squirrel squirrel.StatementBuilderType
	db       *sql.DB
}

func NewUserRepository(squirrel squirrel.StatementBuilderType, db *sql.DB) *UserRepository {
	return &UserRepository{
		squirrel: squirrel,
		db:       db,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, input input.CreateUserRequest) (model.User, error) {
	// Build the insert query
	insertQuery := u.squirrel.Insert("users").Columns("name", "email").Values(input.Name, input.Email).Suffix("RETURNING *")
	// Execute the query and get the result
	row := insertQuery.RunWith(u.db).QueryRowContext(ctx)

	// Scan the result into a User model
	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, userId uuid.UUID, input input.UpdateUserRequest) (model.User, error) {
	// Build the update query
	updateQuery := u.squirrel.Update("users").Set("name", input.Name).Set("email", input.Email).Where(squirrel.Eq{"id": userId}).Suffix("RETURNING *")
	// Execute the query and get the result
	row := updateQuery.RunWith(u.db).QueryRowContext(ctx)

	// Scan the result into a User model
	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return model.User{}, err
	}

	return user, nil
}
