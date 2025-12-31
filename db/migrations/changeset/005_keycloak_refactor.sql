-- 1. Add keycloak_id column
ALTER TABLE users
    ADD COLUMN IF NOT EXISTS keycloak_id UUID;

-- 2. Backfill existing users with random UUIDs
-- (temporary until real Keycloak IDs are wired)
UPDATE users
SET keycloak_id = gen_random_uuid()
WHERE keycloak_id IS NULL;

-- 3. Enforce constraints
ALTER TABLE users
    ALTER COLUMN keycloak_id SET NOT NULL;

ALTER TABLE users
    ADD CONSTRAINT users_keycloak_id_unique UNIQUE (keycloak_id);

-- 4. Drop auth-related columns
ALTER TABLE users
    DROP COLUMN IF EXISTS password_hash,
    DROP COLUMN IF EXISTS username,
    DROP COLUMN IF EXISTS email;