-- Database access
GRANT CONNECT ON DATABASE chat TO chat_admin;
GRANT CONNECT ON DATABASE chat TO keycloak;

-- Assign ownership of public schema to chat_admin
ALTER SCHEMA public OWNER TO chat_admin;

-- Grant chat_admin full privileges on public
GRANT ALL PRIVILEGES ON SCHEMA public TO chat_admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO chat_admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO chat_admin;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT ALL ON TABLES TO chat_admin;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT ALL ON SEQUENCES TO chat_admin;

-- Create Keycloak schema and assign ownership
CREATE SCHEMA IF NOT EXISTS keycloak AUTHORIZATION keycloak;

-- Ensure the owner is correct
ALTER SCHEMA keycloak OWNER TO keycloak;

GRANT ALL PRIVILEGES ON SCHEMA keycloak TO keycloak;
GRANT CREATE ON SCHEMA keycloak TO keycloak;