CREATE TABLE "nfts"
(
    "token"        CHAR(42)       NOT NULL,
    "identifier"   NUMERIC(78, 0) NOT NULL,

    "owner"        CHAR(42)       NOT NULL,
    "is_burned"    BOOLEAN        NOT NULL,
    "token_uri"    VARCHAR(255),
    "metadata"     jsonb,

    "block_number" numeric(19, 0) NOT NULL,
    "tx_index"     bigint         NOT NULL,

    PRIMARY KEY ("token", "identifier")
);