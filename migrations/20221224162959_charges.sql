-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS charges (
    id SERIAL PRIMARY KEY NOT NULL,
    bq_id INT NOT NULL,
    charges_name VARCHAR(128) NOT NULL,
    units VARCHAR(64) NOT NULL,
    tax VARCHAR(128) NOT NULL ,
    buy_rate VARCHAR(64) NOT NULL,
    sell_rate VARCHAR(64) NOT NULL,
    total_buy_amount VARCHAR(64) NOT NULL,
    total_sell_amount VARCHAR(64) NOT NULL,
    charges_type VARCHAR(64) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
