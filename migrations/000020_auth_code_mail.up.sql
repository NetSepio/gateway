-- add auth_code column to email_auths table
ALTER TABLE email_auths
ADD COLUMN auth_code TEXT UNIQUE;
