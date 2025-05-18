-- +migrate Up
ALTER TABLE users ADD COLUMN email VARCHAR(255) NOT NULL UNIQUE;

-- +migrate Down
ALTER TABLE users DROP COLUMN email;