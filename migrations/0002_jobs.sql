CREATE TYPE job_status AS ENUM ('QUEUED','RUNNING','SUCCEEDED','FAILED','CANCELED');
CREATE TYPE run_status AS ENUM ('RUNNING','SUCCEEDED','FAILED');
CREATE TYPE step_status AS ENUM ('RUNNING','SUCCEEDED','FAILED');

CREATE TABLE IF NOT EXISTS jobs (
    id           BIGSERIAL PRIMARY KEY,
    project_id   BIGINT NOT NULL REFERENCES projects(id),
    agent_id     BIGINT NOT NULL REFERENCES agents(id),
    prompt       TEXT   NOT NULL,
    status       job_status NOT NULL DEFAULT 'QUEUED',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    started_at   TIMESTAMPTZ,
    finished_at  TIMESTAMPTZ,
    cancel_reason TEXT
    );

CREATE TABLE IF NOT EXISTS runs (
    id           BIGSERIAL PRIMARY KEY,
    ]job_id       BIGINT NOT NULL REFERENCES jobs(id),
    status       run_status NOT NULL DEFAULT 'RUNNING',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    started_at   TIMESTAMPTZ,
    finished_at  TIMESTAMPTZ
    );

CREATE TABLE IF NOT EXISTS steps (
    id           BIGSERIAL PRIMARY KEY,
    run_id       BIGINT NOT NULL REFERENCES runs(id),
    idx          INT   NOT NULL,
    type         TEXT  NOT NULL,
    input        JSONB NOT NULL DEFAULT '{}',
    output       JSONB NOT NULL DEFAULT '{}',
    tool_name    TEXT,
    status       step_status NOT NULL DEFAULT 'RUNNING',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    finished_at  TIMESTAMPTZ
    );

CREATE UNIQUE INDEX IF NOT EXISTS uq_steps_run_idx ON steps(run_id, idx);
CREATE INDEX IF NOT EXISTS idx_jobs_project_created ON jobs(project_id, created_at DESC);