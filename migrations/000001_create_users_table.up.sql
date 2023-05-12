CREATE TABLE IF NOT EXISTS "users" (
  "id" serial NOT NULL,
  "unique_id" uuid NOT NULL,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
    PRIMARY KEY("id")
);
