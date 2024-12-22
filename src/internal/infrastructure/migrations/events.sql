CREATE TABLE events (
    id UUID PRIMARY KEY,                -- Уникальный идентификатор события
    aggregate_id UUID NOT NULL,         -- Идентификатор агрегата (клиента, счёта, транзакции)
    aggregate_type VARCHAR(50) NOT NULL, -- Тип агрегата (client, account, transaction)
    event_type VARCHAR(50) NOT NULL,    -- Тип события (created, updated, deleted и т.д.)
    payload JSONB NOT NULL,             -- Данные события в формате JSON
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Время создания события
    version INT NOT NULL                -- Версия события для обеспечения последовательности
);
