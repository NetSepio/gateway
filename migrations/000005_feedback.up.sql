-- Drop existing constraint
ALTER TABLE public.user_feedbacks
DROP CONSTRAINT IF EXISTS fk_users_feedbacks;

-- Add new column
ALTER TABLE public.user_feedbacks
ADD COLUMN user_id UUID;

-- Update user_id
UPDATE public.user_feedbacks
SET user_id = u.user_id
FROM public.users u
WHERE u.wallet_address = user_feedbacks.wallet_address;

-- Add new foreign key constraint and drop old column
ALTER TABLE public.user_feedbacks
ADD CONSTRAINT fk_users_feedbacks FOREIGN KEY (user_id) REFERENCES public.users(user_id),
DROP COLUMN wallet_address;