CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       username VARCHAR(50) NOT NULL UNIQUE,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Messages table
CREATE TABLE messages (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

                          sender_id UUID NOT NULL,
                          receiver_id UUID NOT NULL,

                          content TEXT NOT NULL,

                          created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

                          CONSTRAINT fk_sender
                              FOREIGN KEY (sender_id)
                                  REFERENCES users(id)
                                  ON DELETE CASCADE,

                          CONSTRAINT fk_receiver
                              FOREIGN KEY (receiver_id)
                                  REFERENCES users(id)
                                  ON DELETE CASCADE,

                          CONSTRAINT chk_not_self_message
                              CHECK (sender_id <> receiver_id)
);

-- Useful indexes
CREATE INDEX idx_messages_sender ON messages(sender_id);
CREATE INDEX idx_messages_receiver ON messages(receiver_id);
CREATE INDEX idx_messages_created_at ON messages(created_at);