// Package store defines the models, interface and implementation of object as a repository
package store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Store collects all the CRUD of components
type Store interface {
	JobStore
	JobAttemptStore
	WorkerStore
}

// JobStore defines all the CRUD for jobs
type JobStore interface {
	JobCreator
	JobReader
	JobUpdater
	JobDeleter
}

type JobFilter struct {
	Status        []JobStatus
	Priority      *int16
	CreatedFrom   *time.Time
	CreatedTo     *time.Time
	ScheduledFrom *time.Time
	ScheduledTo   *time.Time
	UpdatedFrom   *time.Time
	UpdatedTo     *time.Time
}

type JobUpdate struct {
	Status      *JobStatus
	Priority    *int16
	MaxAttempts *int16
	CompletedAt *time.Time
	DeadAt      *int16
}

type JobCreator interface {
	CreateJob(ctx context.Context, job *Job) (*Job, error)
}

type JobReader interface {
	GetJob(ctx context.Context, jobID uuid.UUID) (*Job, error)
	ListJobs(ctx context.Context, filter JobFilter, page Pagination) ([]Job, error)
}

type JobUpdater interface {
	UpdateJob(ctx context.Context, jobID uuid.UUID, update JobUpdate) (*Job, error)
}

type JobDeleter interface {
	DeleteJob(ctx context.Context, jobID uuid.UUID) error
}

// JobAttemptStore defines all the CRUD for job attempts
type JobAttemptStore interface {
	JobAttemptCreator
	JobAttemptReader
	JobAttemptUpdater
	JobAttemptDeleter
}

type JobAttemptFilter struct {
	JobID          *uuid.UUID
	Error          []string
	FinishedFrom   *time.Time
	FinishedTo     *time.Time
	Success        *bool
	DurationMSFrom *int
	DurationMSTo   *int
}

type JobAttemptUpdate struct {
	JobID      *uuid.UUID
	WorkerID   *string
	StartedAt  *time.Time
	FinishedAt *time.Time
	Success    *bool
	Error      *string
	DurationMS *int
}

type JobAttemptCreator interface {
	CreateJobAttempt(ctx context.Context, jobAttempt *JobAttempt) (*JobAttempt, error)
}

type JobAttemptReader interface {
	GetJobAttempt(ctx context.Context, attemptID uuid.UUID) (*JobAttempt, error)
	ListJobAttempts(ctx context.Context, filter JobAttemptFilter, page Pagination) ([]JobAttempt, error)
}

type JobAttemptUpdater interface {
	UpdateJobAttempt(ctx context.Context, jobAttemptID uuid.UUID, update JobAttemptUpdate) (*JobAttempt, error)
}

type JobAttemptDeleter interface {
	DeleteJobAttempt(ctx context.Context, jobAttemptID uuid.UUID) error
}

// WorkerStore defines all the CRUD for the workers
type WorkerStore interface {
	WorkerCreator
	WorkerReader
	WorkerUpdater
	WorkerDeleter
}

type WorkerFilter struct {
	ID                []string
	Hostname          []string
	Status            *WorkerStatus
	LastHeartbeatFrom *time.Time
	LastHeartbeatTo   *time.Time
	RegisteredFrom    *time.Time
	RegisteredTo      *time.Time
}

type WorkerUpdate struct {
	ID            *string
	Hostname      *string
	Status        *WorkerStatus
	LastHeartbeat *time.Time
	RegisterdAt   *time.Time
}

type WorkerCreator interface {
	CreateWorker(ctx context.Context, worker *Worker) (*Worker, error)
}

type WorkerReader interface {
	GetWorker(ctx context.Context, workerID string) (*Worker, error)
	ListWorkers(ctx context.Context, filter WorkerFilter, page Pagination) ([]Worker, error)
}

type WorkerUpdater interface {
	UpdateWorker(ctx context.Context, workerID string, update WorkerUpdate) (*Worker, error)
}

type WorkerDeleter interface {
	DeleteWorker(ctx context.Context, workerID string) error
}

type Pagination struct {
	OffSet *int
	Limit  *int
}
