CREATE TABLE IF NOT EXISTS urls (
id UUID PRIMARY KEY,
user_id UUID NOT NULL,
short_url text NOT NULL,
orignal_url text NOT NULL,
created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);