CREATE TABLE user_stripe_pis (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(user_id),
    stripe_pi_id TEXT UNIQUE,
    stripe_pi_type TEXT,
    created_at timestamp with time zone DEFAULT current_timestamp
);