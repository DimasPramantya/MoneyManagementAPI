-- +migrate Up

-- +migrate StatementBegin
CREATE TABLE transaction_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id uuid NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP,
    updated_by VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +migrate StatementEnd