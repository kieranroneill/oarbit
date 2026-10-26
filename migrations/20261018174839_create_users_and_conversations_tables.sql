-- +goose Up

/* creates "users" table */
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    full_name TEXT NOT NULL
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    username TEXT NOT NULL
);

/* creates "conversations" table */
CREATE TABLE conversations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,

    active BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_conversations_user
        FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE RESTRICT
);
CREATE INDEX ix_conversations_user_id
    ON conversations (user_id);

/* creates "conversation_messages" table */
CREATE TABLE conversation_messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    conversation_id UUID NOT NULL,
    parent_message_id UUID NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    content TEXT NOT NULL,
    model TEXT NULL,
    role TEXT NOT NULL,
    sequence_number INTEGER NOT NULL,
    input_tokens INTEGER NULL,
    output_tokens INTEGER NULL,
    total_tokens INTEGER NULL,

    CONSTRAINT fk_conversation_messages_conversation
        FOREIGN KEY (conversation_id)
        REFERENCES conversations (id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_conversation_messages_parent_message
        FOREIGN KEY (parent_message_id)
        REFERENCES conversation_messages (id)
        ON DELETE RESTRICT
);
CREATE INDEX ix_conversation_messages_conversation_id
    ON conversation_messages (conversation_id);
CREATE INDEX ix_conversation_messages_parent_message_id
    ON conversation_messages (parent_message_id);

/* creates "tool_calls" table */
CREATE TABLE tool_calls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    message_id UUID NOT NULL,

    arguments JSONB NULL,
    call_id TEXT NOT NULL,
    name TEXT NOT NULL,

    CONSTRAINT fk_tool_calls_message
        FOREIGN KEY (message_id)
        REFERENCES conversation_messages (id)
        ON DELETE RESTRICT
);
CREATE INDEX ix_tool_calls_message_id
    ON tool_calls (message_id);
CREATE UNIQUE INDEX uq_tool_calls_call_id
    ON tool_calls (call_id);

-- +goose Down

DROP TABLE IF EXISTS tool_calls;
DROP TABLE IF EXISTS conversation_messages;
DROP TABLE IF EXISTS conversations;
DROP TABLE IF EXISTS users;
