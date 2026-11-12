CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    full_name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS user_email_lower_idx ON users (LOWER(email));

CREATE TABLE IF NOT EXISTS study_groups (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    year INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (name, year)
);

CREATE TABLE IF NOT EXISTS users_to_study_groups (
    user_id BIGINT NOT NULL REFERENCES users (id),
    study_group_id BIGINT NOT NULL REFERENCES study_groups (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, study_group_id),
    UNIQUE (user_id, study_group_id)
);

CREATE INDEX IF NOT EXISTS idx_users_to_study_groups_group_id ON users_to_study_groups (study_group_id);