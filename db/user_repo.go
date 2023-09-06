package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type UserRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn) *UserRepo {
	return &UserRepo{conn: conn}
}
func (r UserRepo) GetUsers(ctx context.Context) ([]User, error) {
	var users []User
	rows, err := r.conn.Query(ctx, "select id, fullname,email from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user = User{}
		err = rows.Scan(&user.Id, &user.FullName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}
