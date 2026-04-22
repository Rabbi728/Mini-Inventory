DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'inventory_record_type') THEN
        CREATE TYPE inventory_record_type AS ENUM ('IN', 'OUT');
    END IF;
END
$$;

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

CREATE TABLE IF NOT EXISTS products (
    "id" SERIAL PRIMARY KEY,
    "title" character varying(191) NOT NULL,
    "description" text,
    "color" character varying(50),
    "size" character varying(50),
    "uom" character varying(50),
    "product_code" character varying(32) NOT NULL UNIQUE,
    "created_at" timestamp(3) NOT NULL,
    "updated_at" timestamp(3) NOT NULL
);

CREATE TABLE IF NOT EXISTS locations (
    "id" SERIAL PRIMARY KEY,
    "title" character varying(191) NOT NULL,
    "description" text,
    "created_by" INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    "created_at" timestamp(3) NOT NULL,
    "updated_at" timestamp(3) NOT NULL
);

CREATE TABLE IF NOT EXISTS inventories (
    "id" SERIAL PRIMARY KEY,
    "product_id" INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    "location_id" INTEGER NOT NULL REFERENCES locations(id) ON DELETE CASCADE,
    "record_type" inventory_record_type NOT NULL,
    "items" INTEGER NOT NULL DEFAULT 0,
    "note" text,
    "created_by" INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    "created_at" timestamp(3) NOT NULL,
    "updated_at" timestamp(3) NOT NULL
);