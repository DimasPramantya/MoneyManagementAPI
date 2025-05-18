-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transactions (
    id int PRIMARY KEY,
    ammount DECIMAL(10, 2) NOT NULL,
    user_id uuid NOT NULL,
    transaction_category_id int NOT NULL,
    transaction_sub_category_id int NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP,
    modified_by VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (transaction_category_id) REFERENCES transaction_categories(id) ON DELETE CASCADE,
    FOREIGN KEY (transaction_sub_category_id) REFERENCES transaction_sub_categories(id) ON DELETE CASCADE
);

-- +migrate StatementEnd