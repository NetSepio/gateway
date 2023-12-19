-- Add new columns
ALTER TABLE public.domains
ADD COLUMN created_by_id UUID,
ADD COLUMN updated_by_id UUID;

-- Update created_by_id and updated_by_id
UPDATE public.domains
SET created_by_id = u.user_id
FROM public.users u
WHERE u.wallet_address = domains.created_by_address;

UPDATE public.domains
SET updated_by_id = u.user_id
FROM public.users u
WHERE u.wallet_address = domains.updated_by_address;

-- Add foreign key constraints and drop old columns
ALTER TABLE public.domains
ADD CONSTRAINT fk_domains_created_by FOREIGN KEY (created_by_id) REFERENCES public.users(user_id),
ADD CONSTRAINT fk_domains_updated_by FOREIGN KEY (updated_by_id) REFERENCES public.users(user_id),
DROP COLUMN created_by_address,
DROP COLUMN updated_by_address;