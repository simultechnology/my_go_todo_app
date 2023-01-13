package handler

import (
	"github.com/jmoiron/sqlx"
	"net/http"

	"github.com/simultechnology/my_go_todo_app/entity"
	my_json "github.com/simultechnology/my_go_todo_app/my-json"
	"github.com/simultechnology/my_go_todo_app/store"
)

type ListTask struct {
	DB   *sqlx.DB
	Repo *store.Repository
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Repo.ListTasks(ctx, lt.DB)
	if err != nil {
		my_json.RespondJSON(ctx, w, &my_json.ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	my_json.RespondJSON(ctx, w, rsp, http.StatusOK)
}
