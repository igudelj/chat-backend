-- Create application role
DO $$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'chat_admin') THEN
            CREATE ROLE chat_admin LOGIN PASSWORD 'chat_admin';
        END IF;
    END
$$;

