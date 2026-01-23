-- Database access
GRANT CONNECT ON DATABASE chat TO chat_admin;

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
