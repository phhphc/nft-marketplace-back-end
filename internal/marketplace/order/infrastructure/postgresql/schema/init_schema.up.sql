CREATE TABLE IF NOT EXISTS marketplace_order (
    order_hash varchar NOT NULL,
    offerer char(42) NOT NULL,
    is_cancelled boolean NOT NULL,
    is_validated boolean NOT NULL,
    signature varchar,
    order_type numeric(78,0) NOT NULL,
    start_time numeric(78,0) NOT NULL,
    end_time numeric(78,0) NOT NULL,
    counter numeric(78,0) NOT NULL,
    salt varchar NOT NULL,
    zone char(42),
    zone_hash varchar,

    created_at timestamp,
    modified_at timestamp,

    PRIMARY KEY (order_hash)
);

CREATE TABLE IF NOT EXISTS marketplace_order_offer (
    id SERIAL,
    order_hash varchar NOT NULL,
    type_number numeric(78,0) NOT NULL,
    token_id numeric(78,0) NOT NULL,
    token_address char(42) NOT NULL,
    start_amount numeric(78,0) NOT NULL,
    end_amount numeric(78,0) NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (order_hash) REFERENCES marketplace_order(order_hash)
);

CREATE TABLE IF NOT EXISTS marketplace_order_consideration (
    id SERIAL,
    order_hash varchar NOT NULL,
    type_number numeric(78,0) NOT NULL,
    token_id numeric(78,0) NOT NULL,
    token_address char(42) NOT NULL,
    start_amount numeric(78,0) NOT NULL,
    end_amount numeric(78,0) NOT NULL,
    recipient char(42) NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (order_hash) REFERENCES marketplace_order(order_hash)
);
