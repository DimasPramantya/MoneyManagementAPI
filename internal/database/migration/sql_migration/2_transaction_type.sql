-- +migrate Up
CREATE TYPE transaction_type_enum AS ENUM ('income', 'expense');

-- +migrate StatementBegin
CREATE TABLE transaction_categories (
    id int PRIMARY KEY,
    type transaction_type_enum NOT NULL,
    name VARCHAR(255) NOT NULL,
    user_id uuid NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +migrate StatementEnd