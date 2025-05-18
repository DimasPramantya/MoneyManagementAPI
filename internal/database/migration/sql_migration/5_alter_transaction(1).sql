-- +migrate Up
ALTER TABLE transactions
ALTER COLUMN transaction_sub_category_id DROP NOT NULL;

-- +migrate Down
ALTER TABLE transactions
ALTER COLUMN transaction_sub_category_id SET NOT NULL;