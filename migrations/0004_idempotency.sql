CREATE TABLE IF NOT EXISTS idempotency_keys (
    key       TEXT PRIMARY KEY,
    scope     TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);