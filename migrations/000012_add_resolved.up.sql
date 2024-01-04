ALTER TABLE
    public.reports
ADD
    COLUMN up_votes INT DEFAULT 0,
ADD
    COLUMN down_votes INT DEFAULT 0,
ADD
    COLUMN not_sure INT DEFAULT 0,
ADD
    COLUMN total_votes INT DEFAULT 0;