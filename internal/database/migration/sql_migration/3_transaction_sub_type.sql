-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transaction_sub_categories (
    id int PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    transaction_category_id int NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255),
    FOREIGN KEY (transaction_category_id) REFERENCES transaction_categories(id) ON DELETE CASCADE
);

-- +migrate StatementEnd