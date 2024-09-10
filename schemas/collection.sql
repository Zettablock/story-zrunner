CREATE TABLE collection (
    block_number bigint NOT NULL,
    block_time bigint NOT NULL,
    id text,
    asset_count bigint,
    raised_dispute_count bigint,
    cancelled_dispute_count bigint,
    resolved_dispute_count bigint,
    judged_dispute_count bigint,
    licenses_count bigint,
    primary key (id)
);