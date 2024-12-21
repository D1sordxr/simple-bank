CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    source_account_id UUID,
    destination_account_id UUID,
    currency VARCHAR(10),
    amount NUMERIC,
    status VARCHAR(255),
    type VARCHAR(255),
    description VARCHAR(255),
    failure_reason VARCHAR(255) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
