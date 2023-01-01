-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY NOT NULL,
    bq_id INT NOT NULL ,
    count INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    CONSTRAINT fk_booking_quote
     FOREIGN KEY(bq_id) 
        REFERENCES booking_quote(id)

);

DROP TRIGGER IF EXISTS set_timestamp ON tasks;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON tasks
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON tasks;

DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd

