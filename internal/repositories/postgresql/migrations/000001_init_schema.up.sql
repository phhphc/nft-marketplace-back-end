CREATE TABLE "nfts" (
    "token_id" numeric(78,0) NOT NULL,
    "contract_addr" char(42) NOT NULL,
    "owner" char(42) NOT NULL,
    "is_burned" boolean NOT NULL,

    "block_number" numeric(19,0) NOT NULL,
    "tx_index" bigint NOT NULL,
    
    PRIMARY KEY ("token_id", "contract_addr")
);

CREATE TABLE "listings" (
    "listing_id" numeric(78,0) NOT NULL,
    "collection" char(42) NOT NULL,
    "token_id" numeric(78,0) NOT NULL,
    "seller" char(42) NOT NULL,
    "price" numeric(78,0) NOT NULL,
    "status" char(42) NOT NULL,

    "block_number" numeric(19,0) NOT NULL,
    "tx_index" bigint NOT NULL,

    PRIMARY KEY ("listing_id"),
    UNIQUE ("collection", "token_id")
);