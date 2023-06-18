CREATE TABLE "events"
(
    "id"            SERIAL PRIMARY KEY,
    "name"          VARCHAR  NOT NULL,
    "token"         CHAR(42) NOT NULL,
    "token_id"      VARCHAR NOT NULL,
    "quantity"      INT,
    "type"          VARCHAR,
    "price"         VARCHAR,
    "from"          CHAR(42) NOT NULL,
    "to"            CHAR(42),
    "date"          TIMESTAMP DEFAULT current_timestamp,
    "tx_hash"       VARCHAR,
    
    "order_hash"   CHAR(66)
);

