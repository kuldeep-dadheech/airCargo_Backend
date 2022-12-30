-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cargo (
    id SERIAL PRIMARY KEY NOT NULL,
    enquiry_id INT NOT NULL ,
    count INT NOT NULL,
    length INT NOT NULL,
    width INT NOT NULL,
    height INT NOT NULL,
    volume VARCHAR(64) NOT NULL,
    unit VARCHAR(64) NOT NULL,
    weight INT NOT NULL,
    total_volume VARCHAR(64) NOT NULL,
    total_weight VARCHAR(64) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    CONSTRAINT fk_enquiry
     FOREIGN KEY(enquiry_id) 
        REFERENCES enquiry(id)

);

DROP TRIGGER IF EXISTS set_timestamp ON cargo;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON cargo
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON cargo;

DROP TABLE IF EXISTS cargo;
-- +goose StatementEnd

