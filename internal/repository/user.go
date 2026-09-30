package repository

import (
	"context"
	"rest_api/internal/model"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct{
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn)

func (r *UserRepository)Create(ctx context.Context,
	user model.User,)(model.User,error){
	err:=r.db.QueryRow(ctx,
	`INSERT INTO users(
	email,password_hash)
	VALUES($1,$2)
	RETURNING id`,
	user.Email,
	user.PasswordHash).Scan(&user.ID)
	return user,err
}

func(r *UserRepository)GetByEmail(ctx context.Context,
email string)(model.User,error){
	var user model.User
	err:=r.db.QueryRow(ctx,
	`SELECT id,email,password_hash
	FROM users
	WHERE email = $1`,
email).Scan(&user.ID,
&user.Email,
&user.PasswordHash)
return user,err
}