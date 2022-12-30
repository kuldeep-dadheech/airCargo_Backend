-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS enquiry (
    id SERIAL PRIMARY KEY NOT NULL,
    customer VARCHAR(128) NOT NULL,
    terms_of_shipment VARCHAR(128) NOT NULL,
    origin_customs BOOLEAN NOT NULL DEFAULT FALSE,
    door_pickup BOOLEAN NOT NULL DEFAULT FALSE,
    origin_port VARCHAR(64) NOT NULL,
    pickup_address VARCHAR(64),
    cargo_ready_date  VARCHAR(64),
    destination_customs  BOOLEAN  NOT NULL DEFAULT FALSE,
    door_delivery  BOOLEAN  NOT NULL DEFAULT FALSE,
    destination_port  VARCHAR(64) NOT NULL,
    delivery_address  VARCHAR(64),
    hs_code VARCHAR(64) NOT NULL,
    remarks TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON enquiry;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON enquiry
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON enquiry;

DROP TABLE IF EXISTS enquiry;
-- +goose StatementEnd

