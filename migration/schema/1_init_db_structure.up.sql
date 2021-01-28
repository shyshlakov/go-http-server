CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table tags (
	id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name text NOT NULL
);

create table authors (
	id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name text NOT null,
    register_on date,
    image text
);

create table articles (
	id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    slug text,
    title text,
    description text,
    body text,
    favorites_count int,
    comment jsonb,
    score float,
    liked_users text array,
    author_id UUID REFERENCES authors(id) ON DELETE CASCADE
);