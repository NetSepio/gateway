ALTER TABLE ONLY public.reviews ADD COLUMN site_rating INT DEFAULT 0 NOT NULL;
-- Update 0 with original value for existing reviews