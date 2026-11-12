package user

import "context"

type Service interface {
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, email string, passwordHash string, fullName string) (*User, error)
}

type baseService struct {
	repo UserRepo
}

func NewService(repo UserRepo) Service {
	return &baseService{
		repo: repo,
	}
}

func (s *baseService) FindByID(ctx context.Context, id int64) (*User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *baseService) FindByEmail(ctx context.Context, email string) (*User, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *baseService) Create(ctx context.Context, email string, passwordHash string, fullName string) (*User, error) {
	return s.repo.Create(ctx, email, passwordHash, fullName)
}
