-- Create Keycloak role if it does not exist
-- done this because of liquibase shenanigans
DO
$do$
    BEGIN
        PERFORM 1
        FROM pg_catalog.pg_roles
        WHERE rolname = 'keycloak';

        IF NOT FOUND THEN
            CREATE ROLE keycloak LOGIN PASSWORD 'keycloak';
        END IF;
    END
$do$;

-- Grant database connection
GRANT CONNECT ON DATABASE chat TO keycloak;

-- Grant schema usage and creation rights
GRANT USAGE, CREATE ON SCHEMA public TO keycloak;

-- Grant privileges on all existing tables and sequences
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO keycloak;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO keycloak;

-- Set default privileges for future objects
ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT ALL ON TABLES TO keycloak;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT ALL ON SEQUENCES TO keycloak;

-- Create dedicated schema for Keycloak
CREATE SCHEMA IF NOT EXISTS keycloak AUTHORIZATION keycloak;
GRANT ALL PRIVILEGES ON SCHEMA keycloak TO keycloak;
