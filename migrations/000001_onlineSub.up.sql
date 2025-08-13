CREATE TABLE IF NOT EXISTS Subscription (
    service_name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL DEFAULT 0,
    user_id VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    CONSTRAINT onlineSub UNIQUE (service_name, user_id, start_date)
);
