package user

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgxpoolRepo struct {
	pool *pgxpool.Pool
}

func NewPgxpoolRepo(pool *pgxpool.Pool) UserRepo {
	return &pgxpoolRepo{
		pool: pool,
	}
}

func (r *pgxpoolRepo) FindByID(ctx context.Context, id int64) (*User, error) {
	sql := `
		SELECT id, email, password_hash, full_name, created_at, updated_at, activated_at
		FROM users
		WHERE id = $1
	`

	rows, err := r.pool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}

	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *pgxpoolRepo) FindByEmail(ctx context.Context, email string) (*User, error) {
	sql := `
		SELECT id, email, password_hash, full_name, created_at, updated_at, activated_at
		FROM users
		WHERE email = $1
	`

	rows, err := r.pool.Query(ctx, sql, email)
	if err != nil {
		return nil, err
	}

	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *pgxpoolRepo) Create(ctx context.Context, email string, password_hash string, fullName string) (*User, error) {
	sql := `
		INSERT INTO users (email, password_hash, full_name) 
		VALUES ($1, $2, $3)
		RETURNING id, email, password_hash, full_name, created_at, updated_at, activated_at
	`

	rows, err := r.pool.Query(ctx, sql, email, password_hash, fullName)
	if err != nil {
		return nil, err
	}

	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])

	if err != nil {
		return nil, err
	}

	return &u, nil
}
