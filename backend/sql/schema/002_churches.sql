-- +goose Up
CREATE TYPE church_member_role AS ENUM (
    'admin',
    'member',
    'pastor'
);

CREATE TABLE churches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT church_name_unique UNIQUE (name)
);

CREATE TABLE church_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    church_id UUID NOT NULL REFERENCES churches(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role church_member_role NOT NULL DEFAULT 'member',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT church_members_church_user_unique UNIQUE (church_id, user_id),
    CONSTRAINT church_members_user_unique UNIQUE (user_id)
);

-- +goose Down
DROP TABLE IF EXISTS church_members;
DROP TABLE IF EXISTS churches;
DROP TYPE IF EXISTS church_member_role;