CREATE TABLE public.erebrus (
    UUID text NOT NULL,
    name text NOT NULL,
    wallet_address text,
    region text
);

ALTER TABLE ONLY public.erebrus
    ADD CONSTRAINT erebrus_pkey PRIMARY KEY (UUID);
