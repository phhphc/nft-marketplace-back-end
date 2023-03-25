INSERT INTO "nfts" (token, identifier, owner, is_burned, token_uri, metadata, block_number, tx_index)
VALUES
    ('0xB6D6A911A1995c58DAC0C986BA8F123ddFa58B52', 1, '0xa95c9bebac4d0d7d255a01a22fd15a61dca55a55', false, 'null', null, 1,1),
    ('0xB6D6A911A1995c58DAC0C986BA8F123ddFa58B52', 2, '0xa95c9bebac4d0d7d255a01a22fd15a61dca55a55', false, 'null', null, 1,1),
    ('0xB6D6A911A1995c58DAC0C986BA8F123ddFa58B52', 3, '0x29038b82c6d8c091233b9a9e07344444cca5df7e', false, 'null', null, 1,1),
    ('0xe802b9b736895cc11f5985394f0b73ba4d8aceca', 1, '0xa95c9bebac4d0d7d255a01a22fd15a61dca55a55', false, 'null', null, 1,1);

INSERT INTO "orders" (order_hash, offerer, zone, recipient, order_type, zone_hash, salt, start_time, end_time, signature, is_cancelled, is_validated, is_fulfilled)
VALUES
    ('6e4e17ffa4ebf5fff59d738cd9cba9fc90f67ec9d343419f27baa714b75c922665', '0xabc', '0xxyz', null, 1, '0x30a7abdfvsadfzvad31066sdfsdfdsfsdzd55c84365775dafcb9ed790956e208', '0x30a7abdfvsadfzvad31066sdfsdfdsfsdzd55c84365775dafcb9ed790956e208', 1, 123123132, null, false, false, false),
    ('cad9cdcd70272a7193c4a44d80dcbcd60824d2f20c90887c27103a08149034f293', '0xabc', '0xacyz', null, 1, '0x30a7ab3106655c84365argadadfadfsfgdsfgsdfgs775dafcb9ed790956e208', '0x30a7abdfvsadfzvad31066sdfsdfdsfsdzd55c84365775dafcb9ed790956e208', 1, 12356132, null, false, false, false);

INSERT INTO "offer_items" (id, order_hash, item_type, token, identifier, amount, start_amount, end_amount)
VALUES
    (1, '6e4e17ffa4ebf5fff59d738cd9cba9fc90f67ec9d343419f27baa714b75c922665', 1, '0xB6D6A911A1995c58DAC0C986BA8F123ddFa58B52', 1, 1, 1, 1),
    (2, 'cad9cdcd70272a7193c4a44d80dcbcd60824d2f20c90887c27103a08149034f293', 1, '0xB6D6A911A1995c58DAC0C986BA8F123ddFa58B52', 1, 1, 1, 1);

INSERT INTO "consideration_items" (id, order_hash, item_type, token, identifier, amount, start_amount, end_amount, recipient)
VALUES
    (1, '6e4e17ffa4ebf5fff59d738cd9cba9fc90f67ec9d343419f27baa714b75c922665', 0, '0x41599d5d643fbef1422ffa5ab9d651131e61e779', 1, 50, 1, 1, '0xb6d6a911a1995c58dac0c986ba8f123ddfa58b52'),
    (2, 'cad9cdcd70272a7193c4a44d80dcbcd60824d2f20c90887c27103a08149034f293', 0, '0x41599d5d643fbef1422ffa5ab9d651131e61e779', 1, 100, 1, 1, '0xb6d6a911a1995c58dac0c986ba8f123ddfa58b52');