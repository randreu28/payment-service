CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    account_owner TEXT NOT NULL,
    balance MONEY NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    account_from INTEGER REFERENCES accounts(id) ON DELETE SET NULL,
    account_to INTEGER REFERENCES accounts(id) ON DELETE SET NULL,
    amount MONEY NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    description TEXT
);

