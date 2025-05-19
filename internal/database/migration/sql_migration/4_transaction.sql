-- +migrate Up
CREATE TYPE transaction_type_enum AS ENUM ('income', 'expense');
-- +migrate StatementBegin

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    ammount BIGINT NOT NULL,
    user_id uuid NOT NULL,
    transaction_category_id int NOT NULL,
    transaction_sub_category_id int NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    notes TEXT,
    transaction_type transaction_type_enum NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP,
    updated_by VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (transaction_category_id) REFERENCES transaction_categories(id) ON DELETE CASCADE,
    FOREIGN KEY (transaction_sub_category_id) REFERENCES transaction_sub_categories(id) ON DELETE CASCADE
);

-- +migrate StatementEnd