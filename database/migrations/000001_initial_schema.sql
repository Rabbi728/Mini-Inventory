CREATE TABLE IF NOT EXISTS users (
    "id" SERIAL PRIMARY KEY,
    "name" character varying(64) NOT NULL,
    "email" character varying(128) NOT NULL,
    "password" character varying(191) NOT NULL,
    "created_at" timestamp(3) NOT NULL,
    "updated_at" timestamp(3) NOT NULL
);

CREATE TABLE IF NOT EXISTS personal_access_tokens (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    "token" character varying(191) NOT NULL,
    "expires_at" timestamp(3) NOT NULL,
    "created_at" timestamp(3) NOT NULL,
    "updated_at" timestamp(3) NOT NULL
);