-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS booking_quote (
    id INT PRIMARY KEY REFERENCES enquiry,
    is_quote BOOLEAN NOT NULL DEFAULT TRUE,
    valid_till VARCHAR(64) NOT NULL,
    vessel_liner VARCHAR(128) NOT NULL,
    vessel_total_transit VARCHAR(64) NOT NULL,
    vessel_free_days VARCHAR(64) NOT NULL,
    freight_charges_vendor VARCHAR(64) NOT NULL,
    freight_other_charges_vendor VARCHAR(64) NOT NULL,
    total_buy_amount INT NOT NULL,
    total_sell_amount INT NOT NULL,
    profit_percentage VARCHAR(64) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON booking_quote;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON booking_quote
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON booking_quote;

DROP TABLE IF EXISTS booking_quote;
-- +goose StatementEnd

