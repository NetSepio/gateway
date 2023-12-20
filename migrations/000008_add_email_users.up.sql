ALTER TABLE
    ONLY public.users
ADD
    COLUMN email_id text UNIQUE;