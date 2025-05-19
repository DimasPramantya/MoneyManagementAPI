-- +migrate Up
ALTER TABLE transactions
ALTER COLUMN transaction_date TYPE DATE;

-- +migrate Down
ALTER TABLE transactions
ALTER COLUMN transaction_date TYPE TIMESTAMP;
