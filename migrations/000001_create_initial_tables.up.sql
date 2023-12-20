
CREATE TYPE public.flow_id_type AS ENUM (
    'AUTH',
    'ROLE'
);

CREATE TABLE public.domain_admins (
    domain_id text,
    admin_wallet_address text,
    updated_by_address text,
    name text,
    role text
);

CREATE TABLE public.domains (
    id text NOT NULL,
    domain_name text,
    txt_value text,
    verified boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone,
    title text,
    headline text,
    description text,
    cover_image_hash text,
    logo_hash text,
    category text,
    blockchain text,
    created_by_address text,
    updated_by_address text
);

CREATE TABLE public.flow_ids (
    flow_id_type text,
    wallet_address text,
    flow_id text NOT NULL,
    related_role_id text
);

CREATE TABLE public.reviews (
    voter text,
    meta_data_uri text NOT NULL,
    category text,
    domain_address text,
    site_url text,
    site_type text,
    site_tag text,
    site_safety text,
    site_ipfs_hash text,
    transaction_hash text,
    transaction_version bigint,
    deleted_at timestamp with time zone,
    created_at timestamp with time zone
);

CREATE TABLE public.roles (
    name text,
    role_id text NOT NULL,
    eula text
);

CREATE TABLE public.sotreus (
    name text NOT NULL,
    wallet_address text,
    region text
);

CREATE TABLE public.erebrus (
    UUID text NOT NULL,
    name text NOT NULL,
    wallet_address text,
    region text
);

CREATE TABLE public.user_feedbacks (
    wallet_address text NOT NULL,
    feedback text NOT NULL,
    rating bigint NOT NULL,
    created_at timestamp with time zone
);

CREATE TABLE public.user_roles (
    wallet_address text,
    role_id text
);

CREATE TABLE public.users (
    name text,
    wallet_address text NOT NULL,
    profile_picture_url text,
    country text,
    discord text,
    twitter text
);

CREATE TABLE public.wait_lists (
    email_id text NOT NULL,
    wallet_address text,
    twitter text,
    discord text
);


ALTER TABLE ONLY public.domains
    ADD CONSTRAINT domains_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.flow_ids
    ADD CONSTRAINT flow_ids_pkey PRIMARY KEY (flow_id);


ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (meta_data_uri);


ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_name_key UNIQUE (name);


ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (role_id);

ALTER TABLE ONLY public.sotreus
    ADD CONSTRAINT sotreus_pkey PRIMARY KEY (name);

ALTER TABLE ONLY public.erebrus
    ADD CONSTRAINT erebrus_pkey PRIMARY KEY (UUID);

ALTER TABLE ONLY public.user_feedbacks
    ADD CONSTRAINT user_feedbacks_pkey PRIMARY KEY (wallet_address, feedback, rating);


ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_wallet_address_role_id_key UNIQUE (wallet_address, role_id);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (wallet_address);

ALTER TABLE ONLY public.wait_lists
    ADD CONSTRAINT wait_lists_pkey PRIMARY KEY (email_id);


ALTER TABLE ONLY public.domain_admins
    ADD CONSTRAINT fk_domain_admins_admin FOREIGN KEY (admin_wallet_address) REFERENCES public.users(wallet_address);




ALTER TABLE ONLY public.domain_admins
    ADD CONSTRAINT fk_domain_admins_domain FOREIGN KEY (domain_id) REFERENCES public.domains(id);




ALTER TABLE ONLY public.domain_admins
    ADD CONSTRAINT fk_domain_admins_updated_by FOREIGN KEY (updated_by_address) REFERENCES public.users(wallet_address);




ALTER TABLE ONLY public.domains
    ADD CONSTRAINT fk_domains_created_by FOREIGN KEY (created_by_address) REFERENCES public.users(wallet_address);


ALTER TABLE ONLY public.domains
    ADD CONSTRAINT fk_domains_updated_by FOREIGN KEY (updated_by_address) REFERENCES public.users(wallet_address);



ALTER TABLE ONLY public.user_feedbacks
    ADD CONSTRAINT fk_users_feedbacks FOREIGN KEY (wallet_address) REFERENCES public.users(wallet_address);


ALTER TABLE ONLY public.flow_ids
    ADD CONSTRAINT fk_users_flow_ids FOREIGN KEY (wallet_address) REFERENCES public.users(wallet_address);