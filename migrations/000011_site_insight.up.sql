CREATE TABLE site_insights (
    site_url TEXT PRIMARY KEY,
    insight TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);