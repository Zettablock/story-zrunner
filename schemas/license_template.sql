CREATE TABLE license_template (
    block_number bigint NOT NULL,
    block_time bigint NOT NULL,
    id text,
    name text,
    metadata_uri text,
    primary key (id)
);
