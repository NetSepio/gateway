ALTER TABLE public.reports
ADD COLUMN transaction_hash text,
ADD COLUMN transaction_version bigint,
ADD COLUMN end_transaction_hash text,
ADD COLUMN end_transaction_version bigint,
ADD COLUMN meta_data_hash text,
ADD COLUMN end_meta_data_hash text;