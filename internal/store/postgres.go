package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgStore struct {
	pool *pgxpool.Pool
}

func NewPGStore(p *pgxpool.Pool) Store {
	return &pgStore{pool: p}
}

// This might not be idiom go?
// func NewPGCreator(p *pgxpool.Pool) Creator {
// 	return &pgStore{pool: p}
// }
//
// func NewPGReader(p *pgxpool.Pool) Reader {
// 	return &pgStore{pool: p}
// }
//
// func NewPGUpdater(p *pgxpool.Pool) Updater {
// 	return &pgStore{pool: p}
// }
//
// func NewPGDeleter(p *pgxpool.Pool) Deleter {
// 	return &pgStore{pool: p}
// }
//
// func NewPGJobReader(p *pgxpool.Pool) JobReader {
// 	return &pgStore{pool: p}
// }
//
// func NewPGJobAttemptReader(p *pgxpool.Pool) JobAttemptReader {
// 	return &pgStore{pool: p}
// }
//
// func NewPGWorkerReader(p *pgxpool.Pool) WorkerReader {
// 	return &pgStore{pool: p}
// }
//
// func NewPGJobUpdater(p *pgxpool.Pool) JobUpdater {
// 	return &pgStore{pool: p}
// }
//
// func NewPGJobAttemptUpdater(p *pgxpool.Pool) JobAttemptUpdater {
// 	return &pgStore{pool: p}
// }
//
// func NewPGWorkerUpdater(p *pgxpool.Pool) WorkerUpdater {
// 	return &pgStore{pool: p}
// }

func (*pgStore) CreateJob(ctx context.Context, job *Job) (*Job, error) {
	return nil, nil
}

func (*pgStore) CreateJobAttempt(ctx context.Context, jobAttempt *JobAttempt) (*JobAttempt, error) {
	return nil, nil
}

func (*pgStore) CreateWorker(ctx context.Context, worker *Worker) (*Worker, error) {
	return nil, nil
}

func (*pgStore) GetJob(ctx context.Context, jobID uuid.UUID) (*Job, error) {
	return nil, nil
}

func (*pgStore) ListJobs(ctx context.Context, filter JobFilter, page Pagination) ([]Job, error) {
	return nil, nil
}

func (*pgStore) GetJobAttempt(ctx context.Context, attemptID uuid.UUID) (*JobAttempt, error) {
	return nil, nil
}

func (*pgStore) ListJobAttempts(ctx context.Context, filter JobAttemptFilter, page Pagination) ([]JobAttempt, error) {
	return nil, nil
}

func (*pgStore) GetWorker(ctx context.Context, workerID string) (*Worker, error) {
	return nil, nil
}

func (*pgStore) ListWorkers(ctx context.Context, filter WorkerFilter, page Pagination) ([]Worker, error) {
	return nil, nil
}

func (*pgStore) UpdateJob(ctx context.Context, jobID uuid.UUID, update JobUpdate) (*Job, error) {
	return nil, nil
}

func (*pgStore) UpdateJobAttempt(ctx context.Context, jobAttemptID uuid.UUID, update JobAttemptUpdate) (*JobAttempt, error) {
	return nil, nil
}

func (*pgStore) UpdateWorker(ctx context.Context, workerID string, update WorkerUpdate) (*Worker, error) {
	return nil, nil
}

func (*pgStore) DeleteJob(ctx context.Context, jobID uuid.UUID) error {
	return nil
}

func (*pgStore) DeleteJobAttempt(ctx context.Context, jobAttemptID uuid.UUID) error {
	return nil
}

func (*pgStore) DeleteWorker(ctx context.Context, workerID string) error {
	return nil
}
