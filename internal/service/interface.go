// Package service deals with the business logic before database operations
package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/imrishhh/goqueue/internal/store"
)

type Service interface {
	JobService
	JobAttemptService
	WorkerService
}

type JobService interface {
	CreateJob(ctx context.Context, job *store.Job) (*store.Job, error)
	GetJob(ctx context.Context, jobID uuid.UUID) (*store.Job, error)
	ListJob(ctx context.Context, filter *store.JobFilter, page store.Pagination) ([]store.Job, error)
	UpdateJob(ctx context.Context, jobUpdate *store.JobUpdate) (*store.Job, error)
	DeleteJob(ctx context.Context, jobID uuid.UUID) error
}

type JobAttemptService interface {
	CreateJobAttempt(ctx context.Context, jobAttempt *store.JobAttempt) (*store.JobAttempt, error)
	GetJobAttempt(ctx context.Context, jobAttemptID uuid.UUID) (*store.JobAttempt, error)
	ListJobAttempt(ctx context.Context, filter *store.JobAttemptFilter, page store.Pagination) ([]store.JobAttempt, error)
	UpdateJobAttempt(ctx context.Context, jobAttemptUpdate *store.JobAttemptUpdate) (*store.JobAttempt, error)
	DeleteJobAttempt(ctx context.Context, jobID uuid.UUID) error
}

type WorkerService interface {
	CreateWorker(ctx context.Context, worker *store.Worker) (*store.Worker, error)
	GetWorker(ctx context.Context, workerID string) (*store.Worker, error)
	ListWorker(ctx context.Context, filter *store.WorkerFilter, page store.Pagination) ([]store.Worker, error)
	UpdateWorker(ctx context.Context, workerUpdate *store.WorkerUpdate) (*store.Worker, error)
	DeleteWorker(ctx context.Context, workerID string) error
}
