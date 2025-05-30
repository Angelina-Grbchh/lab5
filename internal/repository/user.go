package repository

import (
	"context"
	"database/sql"
	"user-crud-api/internal/model"
)

type UserRepo interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	GetByID(ctx context.Context, id int64) (model.User, error)
	List(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, id int64, user model.User) (model.User, error)
	Delete(ctx context.Context, id int64) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user model.User) (model.User, error) {
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id",
		user.FullName, user.Email, user.Age).Scan(&user.ID)
	return user, err
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (model.User, error) {
	var user model.User
	err := r.db.QueryRowContext(ctx,
		"SELECT id, name, email, age FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.FullName, &user.Email, &user.Age)
	return user, err
}

func (r *userRepo) List(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.FullName, &u.Email, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepo) Update(ctx context.Context, id int64, user model.User) (model.User, error) {
	_, err := r.db.ExecContext(ctx,
		"UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4",
		user.FullName, user.Email, user.Age, id)
	user.ID = id
	return user, err
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
