package service

import (
    "context"
    "user-crud-api/internal/model"
    "user-crud-api/internal/repository"
)

type UserService interface {
    Create(ctx context.Context, user model.User) (model.User, error)
    GetByID(ctx context.Context, id int64) (model.User, error)
    List(ctx context.Context) ([]model.User, error)
    Update(ctx context.Context, id int64, user model.User) (model.User, error)
    Delete(ctx context.Context, id int64) error
}

type userService struct {
    repo repository.UserRepo
}

func NewUserService(r repository.UserRepo) UserService {
    return &userService{repo: r}
}

func (s *userService) Create(ctx context.Context, user model.User) (model.User, error) {
    return s.repo.Create(ctx, user)
}

func (s *userService) GetByID(ctx context.Context, id int64) (model.User, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *userService) List(ctx context.Context) ([]model.User, error) {
    return s.repo.List(ctx)
}

func (s *userService) Update(ctx context.Context, id int64, user model.User) (model.User, error) {
    return s.repo.Update(ctx, id, user)
}

func (s *userService) Delete(ctx context.Context, id int64) error {
    return s.repo.Delete(ctx, id)
}
