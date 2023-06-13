CREATE TABLE "marketplace_settings"
(
    id SERIAL NOT NULL PRIMARY KEY,
    marketplace VARCHAR(42) NOT NULL,
    admin VARCHAR(42) NOT NULL,
    signer VARCHAR(42) NOT NULL,
    royalty NUMERIC(12, 12) NOT NULL DEFAULT 0,
    typed_data jsonb,
    sighash VARCHAR,
    signature VARCHAR,
    created_at NUMERIC(78,0)
);

ALTER TABLE "marketplace_settings"
    ADD CONSTRAINT posivite_transaction_fee CHECK (royalty >= 0);


