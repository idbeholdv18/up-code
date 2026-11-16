package studygroup

import "context"

type StudyGroupRepo interface {
	FindByID(ctx context.Context, id int64) (*StudyGroup, error)
	FindByName(ctx context.Context, name string) ([]StudyGroup, error)
	FindByYear(ctx context.Context, year int) ([]StudyGroup, error)
	FindByNameAndYear(ctx context.Context, name string, year int) (*StudyGroup, error)
	Create(ctx context.Context, name string, year int) (*StudyGroup, error)
}
