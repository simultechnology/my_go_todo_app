package service

import (
	"context"

	"github.com/simultechnology/my_go_todo_app/entity"
	"github.com/simultechnology/my_go_todo_app/store"
)

type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}
type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
