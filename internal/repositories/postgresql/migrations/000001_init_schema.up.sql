CREATE TABLE "nfts"
(
    "token"        CHAR(42)       NOT NULL,
    "identifier"   NUMERIC(78, 0) NOT NULL,

    "owner"        CHAR(42)       NOT NULL,
    
    "token_uri"    VARCHAR,
    "metadata"     jsonb,

    "is_burned"    BOOLEAN        NOT NULL,
    "is_hidden" BOOLEAN NOT NULL DEFAULT FALSE,

    "block_number" numeric(19, 0) NOT NULL,
    "tx_index"     bigint         NOT NULL,

    PRIMARY KEY ("token", "identifier")
);