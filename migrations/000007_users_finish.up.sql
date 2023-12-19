ALTER TABLE public.users
    DROP CONSTRAINT users_pkey,
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id),
    ADD CONSTRAINT wallet_address_unique UNIQUE (wallet_address),
    ALTER COLUMN wallet_address DROP NOT NULL;
