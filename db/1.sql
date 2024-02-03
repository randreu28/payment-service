CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    account_owner TEXT NOT NULL,
    balance MONEY NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    account_from INTEGER NOT NULL REFERENCES accounts(id),
    account_to INTEGER NOT NULL REFERENCES accounts(id),
    amount MONEY NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    description TEXT
);

