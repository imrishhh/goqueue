-- +goose Up

CREATE TYPE job_status AS ENUM (
    'pending',
    'running',
    'succeeded',
    'failed',
    'retrying',
    'dead'
);

CREATE TYPE worker_status AS ENUM (
    'alive',
    'dead'
);

CREATE TABLE jobs (
    id              UUID            PRIMARY KEY DEFAULT gen_random_uuid(),
    type            TEXT            NOT NULL,
    payload         JSONB           NOT NULL,
    status          job_status      NOT NULL DEFAULT 'pending',
    priority        SMALLINT        NOT NULL DEFAULT 0,
    max_attempts    SMALLINT        NOT NULL DEFAULT 3,
    attempt_count   SMALLINT        NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT now(),
    scheduled_at    TIMESTAMPTZ     NOT NULL DEFAULT now(),
    completed_at    TIMESTAMPTZ,
    dead_at         TIMESTAMPTZ
);

CREATE INDEX idx_jobs_status          ON jobs (status);
CREATE INDEX idx_jobs_type_status     ON jobs (type, status);
CREATE INDEX idx_jobs_scheduled_at    ON jobs (scheduled_at) WHERE status = 'pending';
CREATE INDEX idx_jobs_updated_at      ON jobs (updated_at);

CREATE TABLE job_attempts (
    id              UUID            PRIMARY KEY DEFAULT gen_random_uuid(),
    job_id          UUID            NOT NULL REFERENCES jobs (id) ON DELETE CASCADE,
    worker_id       TEXT            NOT NULL,
    started_at      TIMESTAMPTZ     NOT NULL DEFAULT now(),
    finished_at     TIMESTAMPTZ,
    success         BOOLEAN,
    error           TEXT,
    duration_ms     INT             GENERATED ALWAYS AS (
                        EXTRACT(EPOCH FROM (finished_at - started_at)) * 1000
                    ) STORED
);

CREATE INDEX idx_attempts_job_id ON job_attempts (job_id);

CREATE TABLE workers (
    id              TEXT            PRIMARY KEY,
    hostname        TEXT            NOT NULL,
    status          worker_status   NOT NULL DEFAULT 'alive',
    capabilities    TEXT[]          NOT NULL DEFAULT '{}',
    last_heartbeat  TIMESTAMPTZ     NOT NULL DEFAULT now(),
    registered_at   TIMESTAMPTZ     NOT NULL DEFAULT now()
);

CREATE INDEX idx_workers_status         ON workers (status);
CREATE INDEX idx_workers_last_heartbeat ON workers (last_heartbeat) WHERE status = 'alive';

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER trg_jobs_updated_at
BEFORE UPDATE ON jobs
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- +goose Down

DROP TRIGGER IF EXISTS trg_jobs_updated_at ON jobs;
DROP FUNCTION IF EXISTS set_updated_at;
DROP TABLE IF EXISTS job_attempts;
DROP TABLE IF EXISTS workers;
DROP TABLE IF EXISTS jobs;
DROP TYPE IF EXISTS worker_status;
DROP TYPE IF EXISTS job_status;
