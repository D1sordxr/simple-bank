ALTER TABLE accounts
ADD CONSTRAINT fk_accounts_clients
FOREIGN KEY (client_id)
REFERENCES clients (id);
