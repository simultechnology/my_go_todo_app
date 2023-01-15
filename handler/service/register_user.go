package service

import (
	"context"
	"fmt"
	"github.com/simultechnology/my_go_todo_app/entity"
	"github.com/simultechnology/my_go_todo_app/store"
)

type RegisterUser struct {
	DB   store.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(ctx context.Context, u *entity.User) (*entity.User, error) {
	err := r.Repo.RegisterUser(ctx, r.DB, u)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return u, nil
}
