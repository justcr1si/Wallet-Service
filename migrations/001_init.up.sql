CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE
    users (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        email text UNIQUE NOT NULL,
        created_at timestamptz NOT NULL DEFAULT now ()
    );

CREATE TABLE
    wallets (
        user_id uuid PRIMARY KEY REFERENCES users (id) ON DELETE CASCADE,
        balance bigint NOT NULL DEFAULT 0,
        updated_at timestamptz NOT NULL DEFAULT now ()
    );

CREATE TABLE
    transactions (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        type text NOT NULL CHECK (type IN ('deposit', 'withdraw', 'transfer')),
        from_user_id uuid NULL,
        to_user_id uuid NULL,
        amount bigint NOT NULL CHECK (amount > 0),
        status text NOT NULL CHECK (status IN ('success', 'failed')),
        reason text NULL,
        created_at timestamptz NOT NULL DEFAULT now ()
    );

CREATE INDEX idx_tx_from_created ON transactions (from_user_id, created_at DESC);

CREATE INDEX idx_tx_to_created ON transactions (to_user_id, created_at DESC);