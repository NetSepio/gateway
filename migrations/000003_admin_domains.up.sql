ALTER TABLE public.domain_admins
ADD COLUMN admin_id UUID,
ADD COLUMN updated_by_id UUID,
DROP CONSTRAINT fk_domain_admins_admin,
DROP CONSTRAINT fk_domain_admins_updated_by;

-- Update admin_id and updated_by_id
UPDATE public.domain_admins
SET admin_id = u.user_id
FROM public.users u
WHERE u.wallet_address = domain_admins.admin_wallet_address;

UPDATE public.domain_admins
SET updated_by_id = u.user_id
FROM public.users u
WHERE u.wallet_address = domain_admins.updated_by_address;

-- Add foreign key constraints and drop old columns
ALTER TABLE public.domain_admins
ADD CONSTRAINT fk_domain_admins_admin FOREIGN KEY (admin_id) REFERENCES public.users(user_id),
DROP COLUMN admin_wallet_address,
DROP COLUMN updated_by_address;