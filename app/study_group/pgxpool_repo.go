package studygroup

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgxpoolStudyGroupRepo struct {
	pool *pgxpool.Pool
}

func NewPgxpoolStudyGroupRepo(pool *pgxpool.Pool) StudyGroupRepo {
	return &pgxpoolStudyGroupRepo{
		pool: pool,
	}
}

func (r *pgxpoolStudyGroupRepo) FindByID(ctx context.Context, id int64) (*StudyGroup, error) {
	sql := `
		SELECT id, name, year, created_at, updated_at
		FROM study_groups
		WHERE id = $1
	`

	rows, err := r.pool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}

	found, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[StudyGroup])
	if err != nil {
		return nil, err
	}

	return &found, nil
}

func (r *pgxpoolStudyGroupRepo) FindByName(ctx context.Context, name string) ([]StudyGroup, error) {
	sql := `
		SELECT id, name, year, created_at, updated_at
		FROM study_groups
		WHERE name = $1
	`

	rows, err := r.pool.Query(ctx, sql, name)
	if err != nil {
		return nil, err
	}

	found, err := pgx.CollectRows(rows, pgx.RowToStructByName[StudyGroup])
	if err != nil {
		return nil, err
	}

	return found, err
}

func (r *pgxpoolStudyGroupRepo) FindByYear(ctx context.Context, year int) ([]StudyGroup, error) {
	sql := `
		SELECT id, name, year, created_at, updated_at
		FROM study_groups
		WHERE year = $1
	`

	rows, err := r.pool.Query(ctx, sql, year)
	if err != nil {
		return nil, err
	}

	found, err := pgx.CollectRows(rows, pgx.RowToStructByName[StudyGroup])
	if err != nil {
		return nil, err
	}

	return found, nil
}

func (r *pgxpoolStudyGroupRepo) FindByNameAndYear(ctx context.Context, name string, year int) (*StudyGroup, error) {
	sql := `
		SELECT id, name, year, created_at, updated_at
		FROM study_groups
		WHERE name = $1 AND year = $2
	`

	rows, err := r.pool.Query(ctx, sql, name, year)
	if err != nil {
		return nil, err
	}

	found, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[StudyGroup])
	if err != nil {
		return nil, err
	}

	return &found, nil
}

func (r *pgxpoolStudyGroupRepo) Create(ctx context.Context, name string, year int) (*StudyGroup, error) {
	sql := `
		INSERT INTO study_groups (name, year)
		VALUES ($1, $2)
		WHERE name = $1 AND year = $2
		RETURNING id, name, year, created_at, updated_at
	`

	rows, err := r.pool.Query(ctx, sql, name, year)
	if err != nil {
		return nil, err
	}

	found, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[StudyGroup])
	if err != nil {
		return nil, err
	}

	return &found, nil
}
