CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Add user_id and set it for existing users
ALTER TABLE public.users ADD COLUMN user_id UUID DEFAULT uuid_generate_v4() UNIQUE;
DO $$
DECLARE 
    user_wallet_address TEXT;
    user_user_id UUID;
BEGIN
    FOR user_wallet_address, user_user_id IN SELECT wallet_address, user_id FROM public.users LOOP
        IF user_user_id IS NULL THEN
            UPDATE public.users SET user_id = uuid_generate_v4() WHERE wallet_address = user_wallet_address;
        END IF;
    END LOOP;
END $$;