CREATE TABLE "notifications"
(
    "id"            SERIAL PRIMARY KEY,
    "info"          VARCHAR NOT NULL,
    "event_name"    VARCHAR NOT NULL,
    "order_hash"    CHAR(66) NOT NULL,
    "address"       CHAR(42) NOT NULL,
    "is_viewed"     BOOLEAN DEFAULT FALSE
);