CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    source_account_id UUID,
    destination_account_id UUID,
    currency VARCHAR(10) NOT NULL,
    amount NUMERIC NOT NULL,
    status VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    failure_reason VARCHAR(255) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);