package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/simultechnology/my_go_todo_app/entity"
	my_json "github.com/simultechnology/my_go_todo_app/my-json"
	"net/http"
)

type RegisterUser struct {
	Service   RegisterUserService
	Validator *validator.Validate
}

func (ru *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		my_json.RespondJSON(ctx, w, &my_json.ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := ru.Validator.Struct(b); err != nil {
		my_json.RespondJSON(ctx, w, &my_json.ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	u := &entity.User{
		Name: b.Name, Password: b.Password, Role: b.Role,
	}
	u, err := ru.Service.RegisterUser(ctx, u)
	if err != nil {
		my_json.RespondJSON(ctx, w, &my_json.ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID entity.UserID `json:"id"`
	}{ID: u.ID}
	my_json.RespondJSON(ctx, w, rsp, http.StatusOK)
}
