-- create table email_auths which will have id, email
CREATE TABLE email_auths (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE,
    created_at timestamp with time zone DEFAULT current_timestamp
);