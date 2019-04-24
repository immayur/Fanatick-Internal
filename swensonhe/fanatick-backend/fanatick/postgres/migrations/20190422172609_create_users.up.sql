CREATE TABLE users
(
    id              TEXT                      NOT NULL primary key,
    firebase_uuid   TEXT                      NOT NULL unique,
    phone_number    VARCHAR(20)               NOT NULL,
    is_active       BOOLEAN     DEFAULT true,
    is_seller       boolean                   NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at      TIMESTAMPTZ DEFAULT now() NOT NULL,
    last_login_time TIMESTAMPTZ               NOT NULL
);