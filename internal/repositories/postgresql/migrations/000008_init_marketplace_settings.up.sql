CREATE TABLE "marketplace_settings"
(
    id SERIAL NOT NULL PRIMARY KEY,
    marketplace VARCHAR(42) NOT NULL,
    beneficiary VARCHAR(42) NOT NULL,
    royalty NUMERIC(12, 12) NOT NULL DEFAULT 0
);

ALTER TABLE "marketplace_settings"
    ADD CONSTRAINT posivite_transaction_fee CHECK (royalty >= 0);

INSERT INTO marketplace_settings (marketplace, beneficiary, royalty)
VALUES ('0x2FB1f7D206ECd5fD434ED623dA67BF269fD753Ef', '0x5cc163BCf461482813a0E2c26a00b9FAaBf612eF', 0.01);