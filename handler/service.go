package handler

import (
	"context"
	"github.com/simultechnology/my_go_todo_app/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService RegisterUserService
type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type ListUsersService interface {
	ListUsers(ctx context.Context) (entity.Users, error)
}
