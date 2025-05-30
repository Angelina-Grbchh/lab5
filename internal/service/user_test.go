package service_test

import (
	"context"
	"errors"
	"testing"
	"user-crud-api/internal/model"
	"user-crud-api/internal/service"
)

type mockRepo struct {
	CreateFn func(ctx context.Context, user model.User) (model.User, error)
	GetFn    func(ctx context.Context, id int64) (model.User, error)
	ListFn   func(ctx context.Context) ([]model.User, error)
	UpdateFn func(ctx context.Context, id int64, user model.User) (model.User, error)
	DeleteFn func(ctx context.Context, id int64) error
}

func (m *mockRepo) Create(ctx context.Context, user model.User) (model.User, error) {
	return m.CreateFn(ctx, user)
}
func (m *mockRepo) GetByID(ctx context.Context, id int64) (model.User, error) {
	return m.GetFn(ctx, id)
}
func (m *mockRepo) List(ctx context.Context) ([]model.User, error) {
	return m.ListFn(ctx)
}
func (m *mockRepo) Update(ctx context.Context, id int64, user model.User) (model.User, error) {
	return m.UpdateFn(ctx, id, user)
}
func (m *mockRepo) Delete(ctx context.Context, id int64) error {
	return m.DeleteFn(ctx, id)
}

func TestUserServiceCRUD(t *testing.T) {
	ctx := context.Background()

	mock := &mockRepo{
		CreateFn: func(ctx context.Context, user model.User) (model.User, error) {
			user.ID = 1
			return user, nil
		},
		GetFn: func(ctx context.Context, id int64) (model.User, error) {
			if id == 1 {
				return model.User{ID: 1, FullName: "Alice", Email: "a@example.com", Age: 25}, nil
			}
			return model.User{}, errors.New("not found")
		},
		ListFn: func(ctx context.Context) ([]model.User, error) {
			return []model.User{{ID: 1, FullName: "Alice", Email: "a@example.com", Age: 25}}, nil
		},
		UpdateFn: func(ctx context.Context, id int64, user model.User) (model.User, error) {
			user.ID = id
			return user, nil
		},
		DeleteFn: func(ctx context.Context, id int64) error {
			if id == 1 {
				return nil
			}
			return errors.New("not found")
		},
	}

	svc := service.NewUserService(mock)

	t.Run("Create", func(t *testing.T) {
		user := model.User{FullName: "Alice", Email: "a@example.com", Age: 25}
		created, err := svc.Create(ctx, user)
		if err != nil || created.ID != 1 {
			t.Errorf("expected ID=1, got %v, err: %v", created.ID, err)
		}
	})

	t.Run("GetByID exists", func(t *testing.T) {
		user, err := svc.GetByID(ctx, 1)
		if err != nil || user.ID != 1 {
			t.Errorf("expected user with ID=1, got %v, err: %v", user, err)
		}
	})

	t.Run("GetByID not found", func(t *testing.T) {
		_, err := svc.GetByID(ctx, 2)
		if err == nil {
			t.Error("expected error for non-existing user")
		}
	})

	t.Run("List", func(t *testing.T) {
		users, err := svc.List(ctx)
		if err != nil || len(users) != 1 {
			t.Errorf("expected 1 user, got %v, err: %v", len(users), err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		updatedUser := model.User{FullName: "Updated", Email: "u@example.com", Age: 30}
		user, err := svc.Update(ctx, 1, updatedUser)
		if err != nil || user.FullName != "Updated" {
			t.Errorf("expected updated user, got %v, err: %v", user, err)
		}
	})

	t.Run("Delete exists", func(t *testing.T) {
		err := svc.Delete(ctx, 1)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})

	t.Run("Delete not found", func(t *testing.T) {
		err := svc.Delete(ctx, 2)
		if err == nil {
			t.Error("expected error for non-existing user")
		}
	})
}
