CREATE TABLE transactions (
    id VARCHAR(255) PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    value DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP 
);
