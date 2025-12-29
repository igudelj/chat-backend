-- Create Keycloak schema and assign ownership
CREATE SCHEMA IF NOT EXISTS keycloak AUTHORIZATION keycloak;

-- Ensure the owner is correct
ALTER SCHEMA keycloak OWNER TO keycloak;

GRANT ALL PRIVILEGES ON SCHEMA keycloak TO keycloak;
GRANT CREATE ON SCHEMA keycloak TO keycloak;
