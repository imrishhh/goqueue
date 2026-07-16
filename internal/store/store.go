package store

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusRunning   JobStatus = "running"
	JobStatusSucceeded JobStatus = "succeeded"
	JobStatusFailed    JobStatus = "failed"
	JobStatusDead      JobStatus = "dead"
)

type Job struct {
	ID           uuid.UUID       `json:"id" db:"id"`
	Type         string          `json:"type" db:"type"`
	Status       JobStatus       `json:"status" db:"status"`
	Payload      json.RawMessage `json:"payload" db:"payload"`
	Priority     int16           `json:"priority" db:"priority"`
	MaxAttempts  int16           `json:"max_attempts" db:"max_attempts"`
	AttemptCount int16           `json:"attempt_count" db:"attempt_count"`
	CreatedAt    time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at" db:"updated_at"`
	ScheduledAt  time.Time       `json:"scheduled_at" db:"scheduled_at"`
	CompletedAt  time.Time       `json:"completed_at" db:"completed_at"`
	DeadAt       time.Time       `json:"dead_at" db:"dead_at"`
}

type JobAttempt struct {
	ID         uuid.UUID `json:"id" db:"id"`
	JobID      uuid.UUID `json:"job_id" db:"job_id"`
	WorkerID   string    `json:"worker_id" db:"worker_id"`
	StartedAt  time.Time `json:"started_at" db:"started_at"`
	FinishedAt time.Time `json:"finished_at" db:"finished_at"`
	Success    bool      `json:"success" db:"success"`
	Error      string    `json:"error" db:"error"`
	DurationMS int       `json:"duration_ms" db:"duration_ms"`
}

type WorkerStatus string

const (
	WorkerStatusAlive WorkerStatus = "alive"
	WorkerStatusDead  WorkerStatus = "dead"
)

type Worker struct {
	ID            string       `json:"id" db:"id"`
	Hostname      string       `json:"hostname" db:"hostname"`
	Status        WorkerStatus `json:"status" db:"status"`
	Capabilities  []string     `json:"capabilities" db:"capabilities"`
	LastHeartbeat time.Time    `json:"last_heartbeat" db:"last_heartbeat"`
	RegisterdAt   time.Time    `json:"registered_at" db:"registered_at"`
}

type Pagination struct {
	OffSet *int
	Limit  *int
}

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
