CREATE TABLE events
(
    id         TEXT PRIMARY KEY,
    name       TEXT        NOT NULL,
    start_at   TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);