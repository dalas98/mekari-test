-- +goose Up
-- +goose StatementBegin
CREATE TABLE employees (
  "id" serial PRIMARY KEY,
  "first_name" TEXT,
  "last_name" TEXT,
  "email" TEXT,
  "hire_date" DATE,
  "created_at" TIMESTAMPTZ NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE employees;
-- +goose StatementEnd
