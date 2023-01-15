package store

import (
	"context"
	"golang.org/x/crypto/bcrypt"

	"github.com/simultechnology/my_go_todo_app/entity"
)

func (r *Repository) ListUsers(
	ctx context.Context, db Queryer) (entity.Users, error) {
	users := entity.Users{}
	sql := `SELECT
			id, name, password,
			role, created, modified
		FROM user;`
	if err := db.SelectContext(ctx, &users, sql); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) RegisterUser(
	ctx context.Context, db Execer, u *entity.User,
) error {
	u.Created = r.Clocker.Now()
	u.Modified = r.Clocker.Now()
	encrypt, err := passwordEncrypt(u.Password)
	if err != nil {
		return err
	}
	sql := `INSERT INTO user
		(name, password, role, created, modified)
	VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, u.Name, encrypt, u.Role,
		u.Created, u.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = entity.UserID(id)
	return nil
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
