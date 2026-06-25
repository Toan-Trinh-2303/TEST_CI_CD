package repository

import (
	"context"

	"ci-cd-go-learn/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) List(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, email, created_at
		FROM users
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *UserRepository) Create(ctx context.Context, req model.CreateUserRequest) (model.User, error) {
	var u model.User
	err := r.db.QueryRow(ctx, `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id, name, email, created_at
	`, req.Name, req.Email).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	return u, err
}