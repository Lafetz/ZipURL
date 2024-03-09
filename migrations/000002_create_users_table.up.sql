CREATE TABLE IF NOT EXISTS users (
id UUID PRIMARY KEY,
created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
username text UNIQUE NOT NULL,
email citext UNIQUE NOT NULL,
hashed_password bytea NOT NULL
);