package service

// Implements service layer for jobattempt
import (
	"context"

	"github.com/google/uuid"
	"github.com/imrishhh/goqueue/internal/store"
)

type jobAttemptService struct {
	st store.JobAttemptStore
}

func NewJobAttemptService(st store.JobAttemptStore) JobAttemptService {
	return &jobAttemptService{st}
}

func (s *jobAttemptService) CreateJobAttempt(ctx context.Context, jobAttempt *store.JobAttempt) (*store.JobAttempt, error) {
	return nil, nil
}

func (s *jobAttemptService) GetJobAttempt(ctx context.Context, jobAttemptID uuid.UUID) (*store.JobAttempt, error) {
	return nil, nil
}

func (s *jobAttemptService) ListJobAttempt(ctx context.Context, filter *store.JobAttemptFilter, page store.Pagination) ([]store.JobAttempt, error) {
	return nil, nil
}

func (s *jobAttemptService) UpdateJobAttempt(ctx context.Context, filter *store.JobAttemptUpdate) (*store.JobAttempt, error) {
	return nil, nil
}

func (s *jobAttemptService) DeleteJobAttempt(ctx context.Context, jobAttemptID uuid.UUID) error {
	return s.st.DeleteJobAttempt(ctx, jobAttemptID)
}
