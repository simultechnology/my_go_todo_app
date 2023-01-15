package handler

import (
	"net/http"

	"github.com/simultechnology/my_go_todo_app/entity"
	my_json "github.com/simultechnology/my_go_todo_app/my-json"
)

type ListUser struct {
	Service ListUsersService
}

type user struct {
	ID       entity.UserID `json:"id"`
	Name     string        `json:"name"`
	Password string        `json:"password"`
	Role     string        `json:"role"`
}

func (lu *ListUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := lu.Service.ListUsers(ctx)
	if err != nil {
		my_json.RespondJSON(ctx, w, &my_json.ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []user{}
	for _, u := range users {
		rsp = append(rsp, user{
			ID:       u.ID,
			Name:     u.Name,
			Password: u.Password,
			Role:     u.Role,
		})
	}
	my_json.RespondJSON(ctx, w, rsp, http.StatusOK)
}
