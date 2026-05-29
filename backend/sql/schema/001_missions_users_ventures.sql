-- +goose Up

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE mission_member_role AS ENUM (
    'admin',
    'member'
);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    email TEXT UNIQUE NOT NULL,
    hashed_password TEXT NOT NULL DEFAULT 'unset'
);

CREATE TABLE ventures (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    created_by_user_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE venture_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role mission_member_role NOT NULL DEFAULT 'member',
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT venture_members_venture_user_unique UNIQUE (venture_id, user_id)
);

CREATE TABLE mission_organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE mission_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mission_id UUID NOT NULL REFERENCES mission_organizations(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role mission_member_role NOT NULL DEFAULT 'member',
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT mission_members_mission_user_unique UNIQUE (mission_id, user_id)
);

CREATE TABLE mission_ventures (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mission_id UUID NOT NULL REFERENCES mission_organizations(id) ON DELETE CASCADE,
    venture_id UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT mission_ventures_mission_venture_unique UNIQUE (mission_id, venture_id)
);

CREATE TABLE mission_venture_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mission_venture_id UUID NOT NULL REFERENCES mission_ventures(id) ON DELETE CASCADE,
    mission_member_id UUID NOT NULL REFERENCES mission_members(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role mission_member_role NOT NULL DEFAULT 'member',
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT mission_venture_members_unique UNIQUE (mission_venture_id, user_id)
);

CREATE INDEX users_email_idx
ON users(email);

CREATE INDEX ventures_created_by_user_id_idx
ON ventures(created_by_user_id);

CREATE INDEX venture_members_venture_id_idx
ON venture_members(venture_id);

CREATE INDEX venture_members_user_id_idx
ON venture_members(user_id);

CREATE INDEX mission_members_mission_id_idx
ON mission_members(mission_id);

CREATE INDEX mission_members_user_id_idx
ON mission_members(user_id);

CREATE INDEX mission_ventures_mission_id_idx
ON mission_ventures(mission_id);

CREATE INDEX mission_ventures_venture_id_idx
ON mission_ventures(venture_id);

CREATE INDEX mission_venture_members_mission_venture_id_idx
ON mission_venture_members(mission_venture_id);

CREATE INDEX mission_venture_members_mission_member_id_idx
ON mission_venture_members(mission_member_id);

CREATE INDEX mission_venture_members_user_id_idx
ON mission_venture_members(user_id);


-- +goose Down

DROP TABLE IF EXISTS mission_venture_members;
DROP TABLE IF EXISTS mission_ventures;
DROP TABLE IF EXISTS mission_members;
DROP TABLE IF EXISTS mission_organizations;
DROP TABLE IF EXISTS venture_members;
DROP TABLE IF EXISTS ventures;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS mission_member_role;