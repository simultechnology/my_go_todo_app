package service

import (
	"context"
	"fmt"
	"github.com/simultechnology/my_go_todo_app/entity"
	"github.com/simultechnology/my_go_todo_app/store"
)

type ListUser struct {
	DB   store.Queryer
	Repo UserLister
}

func (l *ListUser) ListUsers(ctx context.Context) (entity.Users, error) {
	us, err := l.Repo.ListUsers(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return us, nil
}
