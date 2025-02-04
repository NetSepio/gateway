ALTER TABLE
    ONLY public.users
ADD
    COLUMN email text UNIQUE;