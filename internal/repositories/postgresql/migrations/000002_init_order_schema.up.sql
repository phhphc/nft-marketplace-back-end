CREATE TABLE "orders"
(
    "order_hash"   CHAR(66) PRIMARY KEY,

    "offerer"      CHAR(42) NOT NULL,
    "zone"         CHAR(42) NOT NULL,
    "recipient"    CHAR(42),
    "order_type"   INT,
    "zone_hash"    CHAR(66) NOT NULL,
    "salt"         CHAR(66),
    "start_time"   NUMERIC(78, 0),
    "end_time"     NUMERIC(78, 0),

    "signature"    VARCHAR,

    "is_cancelled" BOOLEAN  NOT NULL,
    "is_validated" BOOLEAN  NOT NULL,
    "is_fulfilled" BOOLEAN  NOT NULL
);

CREATE TABLE "offer_items"
(
    "id"           BIGSERIAL PRIMARY KEY,
    "order_hash"   CHAR(66)       NOT NULL,

    "item_type"    INT            NOT NULL,
    "token"        CHAR(42)       NOT NULL,
    "identifier"   NUMERIC(78, 0) NOT NULL,

    "amount"       NUMERIC(78, 0),
    "start_amount" NUMERIC(78, 0),
    "end_amount"   NUMERIC(78, 0),


    FOREIGN KEY ("order_hash") REFERENCES orders ("order_hash")
);

CREATE TABLE "consideration_items"
(
    "id"           BIGSERIAL PRIMARY KEY,
    "order_hash"   CHAR(66)       NOT NULL,

    "item_type"    INT            NOT NULL,
    "token"        CHAR(42)       NOT NULL,
    "identifier"   NUMERIC(78, 0) NOT NULL,

    "amount"       NUMERIC(78, 0),
    "start_amount" NUMERIC(78, 0),
    "end_amount"   NUMERIC(78, 0),
    "recipient"    CHAR(42)       NOT NULL,

    FOREIGN KEY ("order_hash") REFERENCES orders ("order_hash")
);