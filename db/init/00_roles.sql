-- Create application role
DO $$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'chat_admin') THEN
            CREATE ROLE chat_admin LOGIN PASSWORD 'chat_admin';
        END IF;
    END
$$;

-- Create Keycloak role if it does not exist
DO $$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'keycloak') THEN
            CREATE ROLE keycloak LOGIN PASSWORD 'keycloak';
        END IF;
    END
$$;
