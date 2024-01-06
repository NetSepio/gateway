CREATE TABLE domain_claims (
    id UUID PRIMARY KEY,
    domain_id text REFERENCES domains (id) NOT NULL,
    txt TEXT NOT NULL UNIQUE,
    admin_id uuid REFERENCES users (user_id) NOT NULL
);