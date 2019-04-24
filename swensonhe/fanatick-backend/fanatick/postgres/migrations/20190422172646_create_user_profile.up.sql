CREATE TABLE user_profile
(
    id              TEXT PRIMARY KEY,
    user_id         TEXT        NOT NULL UNIQUE,
    first_name      VARCHAR(45) NOT NULL,
    last_name       VARCHAR(45) NOT NULL,
    profile_pic_url TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);