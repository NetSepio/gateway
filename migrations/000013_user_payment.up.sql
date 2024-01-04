ALTER TABLE
    users
ADD
    COLUMN subscription_status text CHECK (
        subscription_status IN (
            'basic',
            'pro monthly',
            'pro yearly'
        )
    ),
ADD
    COLUMN stripe_customer_id text UNIQUE,
ADD
    COLUMN stripe_subscription_status text CHECK (
        stripe_subscription_status IN (
            'incomplete',
            'incomplete_expired',
            'trialing',
            'active',
            'past_due',
            'canceled',
            'unpaid',
            'unset'
        )
    ) DEFAULT 'unset',
ADD
    COLUMN stripe_subscription_id text UNIQUE;

-- set basic subscription status and unset stripe_subscription_status for all users
UPDATE
    users
SET
    subscription_status = 'basic',
    stripe_subscription_status = 'unset';