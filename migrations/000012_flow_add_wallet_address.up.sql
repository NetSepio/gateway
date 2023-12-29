ALTER TABLE ONLY public.flow_ids
ADD COLUMN wallet_address text;

-- update the wallet_address in flow_ids table with the wallet_address from users table
UPDATE public.flow_ids
SET wallet_address = users.wallet_address
FROM users
WHERE flow_ids.user_id = users.user_id;

-- make the wallet_address column not null
ALTER TABLE ONLY public.flow_ids
ALTER COLUMN wallet_address
SET NOT NULL;