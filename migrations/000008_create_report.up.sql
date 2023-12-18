CREATE TABLE public.reports (
    id uuid PRIMARY KEY,
    title text NOT NULL,
    description text,
    document text,
    project_name text,
    project_domain text,
    status text CHECK (status IN ('accepted', 'rejected', 'running')),
    created_by uuid REFERENCES public.users(user_id),
    end_time timestamp with time zone,
    created_at timestamp with time zone DEFAULT current_timestamp
);

CREATE TABLE public.report_tags (
    report_id uuid REFERENCES public.reports(id),
    tag text,
    UNIQUE(report_id, tag)
);

CREATE TABLE public.report_images (
    report_id uuid REFERENCES public.reports(id),
    image_url text,
    UNIQUE(report_id, image_url)
);

CREATE TABLE public.report_votes (
    report_id uuid REFERENCES public.reports(id),
    voter_id uuid REFERENCES public.users(user_id),
    vote_type text CHECK (vote_type IN ('upvote', 'downvote', 'notsure')),
    created_at timestamp with time zone DEFAULT current_timestamp,
    PRIMARY KEY (report_id, voter_id)
);