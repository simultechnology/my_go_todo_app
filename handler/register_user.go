package handler

import (
	"encoding/json"
	"github.com/simultechnology/my_go_todo_app/entity"
	my_json "github.com/simultechnology/my_go_todo_app/my-json"
	"net/http"
)

type RegisterUser struct {
	Service RegisterUserService
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
	u := &entity.User{
		Name: b.name, Password: b.password, Role: b.role,
	}
	u, err := ru.Service.RegisterUser(ctx, u)
	if err != nil {
		my_json.RespondJSON(ctx, w, &my_json.ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
}
