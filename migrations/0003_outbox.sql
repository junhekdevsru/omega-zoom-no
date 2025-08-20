CREATE TABLE IF NOT EXISTS outbox (
    id            BIGSERIAL PRIMARY KEY,
    aggregate     TEXT NOT NULL,
    event_type    TEXT NOT NULL,
    payload       JSONB NOT NULL,
    headers       JSONB NOT NULL DEFAULT '{}',
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    state         TEXT NOT NULL DEFAULT 'NEW' -- NEW|SENT|ERROR
);
CREATE INDEX IF NOT EXISTS idx_outbox_state ON outbox(state, created_at);