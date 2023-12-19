-- Drop existing constraint
ALTER TABLE public.flow_ids
DROP CONSTRAINT IF EXISTS fk_users_flow_ids;

-- Add new column
ALTER TABLE public.flow_ids
ADD COLUMN user_id UUID;

-- Update user_id
UPDATE public.flow_ids
SET user_id = u.user_id
FROM public.users u
WHERE u.wallet_address = flow_ids.wallet_address;

-- Add new foreign key constraint and drop old column
ALTER TABLE public.flow_ids
ADD CONSTRAINT fk_users_flow_ids FOREIGN KEY (user_id) REFERENCES public.users(user_id),
DROP COLUMN wallet_address;