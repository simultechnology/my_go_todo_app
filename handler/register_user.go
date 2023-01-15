package handler

import (
	"encoding/json"
	"github.com/simultechnology/my_go_todo_app/handler/service"
	my_json "github.com/simultechnology/my_go_todo_app/my-json"
	"net/http"
)

type RegisterUser struct {
	Service service.RegisterUser
}

func (ru *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		name     string `json:"name" validate:"required"`
		password string `json:"password" validate:"required"`
		role     string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		my_json.RespondJSON(ctx, w, &my_json.ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	t, err := ru.Service.RegisterUser(ctx, b.Title)
	if err != nil {
		my_json.RespondJSON(ctx, w, &my_json.ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
}
