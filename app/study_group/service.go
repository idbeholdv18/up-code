package studygroup

import "context"

type StudyGroupService interface {
	FindByID(ctx context.Context, id int64) (*StudyGroup, error)
	FindByName(ctx context.Context, name string) ([]StudyGroup, error)
	FindByYear(ctx context.Context, year int) ([]StudyGroup, error)
	FindByNameAndYear(ctx context.Context, name string, year int) (*StudyGroup, error)
	Create(ctx context.Context, name string, year int) (*StudyGroup, error)
}

type studyGroupBaseService struct {
	repo StudyGroupRepo
}

func NewStudyGroupBaseService(repo StudyGroupRepo) StudyGroupService {
	return &studyGroupBaseService{
		repo: repo,
	}
}

func (s studyGroupBaseService) FindByID(ctx context.Context, id int64) (*StudyGroup, error) {
	return s.repo.FindByID(ctx, id)
}

func (s studyGroupBaseService) FindByName(ctx context.Context, name string) ([]StudyGroup, error) {
	return s.repo.FindByName(ctx, name)
}

func (s studyGroupBaseService) FindByYear(ctx context.Context, year int) ([]StudyGroup, error) {
	return s.repo.FindByYear(ctx, year)
}

func (s studyGroupBaseService) FindByNameAndYear(ctx context.Context, name string, year int) (*StudyGroup, error) {
	return s.repo.FindByNameAndYear(ctx, name, year)
}

func (s studyGroupBaseService) Create(ctx context.Context, name string, year int) (*StudyGroup, error) {
	return s.repo.Create(ctx, name, year)
}
