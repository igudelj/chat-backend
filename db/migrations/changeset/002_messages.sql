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