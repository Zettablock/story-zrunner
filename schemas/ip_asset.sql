CREATE TABLE ip_asset (
    block_number bigint NOT NULL,
    block_time bigint NOT NULL,
    id text,
    ip_id text,
    chain_id bigint,
    token_contract text,
    token_id bigint,
    metadata JSONB,
    child_ip_ids text[],
    parent_ip_ids text[],
    root_ip_ids text[],
    nft_name text,
    nft_token_uri text,
    nft_image_url text,
    primary key (id)
);