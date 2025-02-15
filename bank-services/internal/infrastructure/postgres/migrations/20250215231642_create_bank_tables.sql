-- +goose Up
-- +goose StatementBegin

CREATE TABLE clients (
                         id UUID PRIMARY KEY,
                         first_name VARCHAR(255) NOT NULL,
                         last_name VARCHAR(255) NOT NULL,
                         middle_name VARCHAR(255) NOT NULL,
                         email VARCHAR(255) NOT NULL UNIQUE,
                         status VARCHAR(255) NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE phones (
                        id UUID PRIMARY KEY,
                        client_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
                        phone_number VARCHAR(25) NOT NULL UNIQUE,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
                          id UUID PRIMARY KEY,
                          client_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
                          available_money NUMERIC NOT NULL,
                          frozen_money NUMERIC NOT NULL,
                          currency VARCHAR(10) NOT NULL,
                          status VARCHAR(25) NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

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
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE events (
                        id UUID PRIMARY KEY,
                        aggregate_id UUID NOT NULL,
                        aggregate_type VARCHAR(50) NOT NULL,
                        event_type VARCHAR(50) NOT NULL,
                        payload TEXT NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE outbox (
                        id UUID PRIMARY KEY,
                        aggregate_id UUID NOT NULL,
                        aggregate_type VARCHAR(50) NOT NULL,
                        message_type VARCHAR(50) NOT NULL,
                        message_payload TEXT NOT NULL,
                        status VARCHAR(50) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS outbox;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS phones;
DROP TABLE IF EXISTS clients;

-- +goose StatementEnd
