CREATE TABLE license_term (
    block_number bigint NOT NULL,
    block_time bigint NOT NULL,
    id text,
    json text,
    license_template text,
    primary key (id)
);