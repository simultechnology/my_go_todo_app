package service

import (
	"context"
	"fmt"
	"github.com/simultechnology/my_go_todo_app/entity"
	"github.com/simultechnology/my_go_todo_app/store"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	DB   store.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(ctx context.Context, u *entity.User) (*entity.User, error) {
	encrypt, err := passwordEncrypt(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = encrypt
	err = r.Repo.RegisterUser(ctx, r.DB, u)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return u, nil
}

// 暗号(Hash)化
func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 暗号(Hash)と入力された平パスワードの比較
func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
