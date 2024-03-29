package store

import (
	"context"
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
	sql := `INSERT INTO user
		(name, password, role, created, modified)
	VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, u.Name, u.Password, u.Role,
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
