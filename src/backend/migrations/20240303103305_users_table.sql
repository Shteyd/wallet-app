-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY,
    "username" VARCHAR(255),
    "email" VARCHAR(360) NOT NULL,
    "secret" VARCHAR(512) NOT NULL,
    "is_blocked" BOOLEAN NOT NULL DEFAULT FALSE,
    "is_confirmed" BOOLEAN NOT NULL DEFAULT FALSE,
    "is_deleted" BOOLEAN NOT NULL DEFAULT FALSE,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY ("id"),
    CONSTRAINT "users_email_unique" UNIQUE ("email")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
-- +goose StatementEnd
