package user

import "context"

type UserRepo interface {
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, email string, password_hash string, fullName string) (*User, error)
}
