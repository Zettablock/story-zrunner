CREATE TABLE transaction (
    block_number bigint NOT NULL,
    block_time bigint NOT NULL,
    id text,
    tx_hash text,
    transaction_index int NOT NULL,
    log_index int NOT NULL,
    initiator text,
    ip_id text,
    resource_id text,
    resource_type text,
    action_type text,
    created_at bigint,
    primary key (id)
);
