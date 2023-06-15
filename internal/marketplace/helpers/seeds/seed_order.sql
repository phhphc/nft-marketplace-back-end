-- seed for me 10 order with these spec
-- INSERT INTO marketplace_order(order_hash, offerer, order_value, is_cancelled, is_validated, signature, order_type, start_time, end_time, counter, salt, zone)
-- VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);

INSERT INTO marketplace_order(order_hash, offerer, order_value, is_cancelled, is_validated, signature, order_type, start_time, end_time, counter, salt, zone)
VALUES('0x2', '0x3', 4, false, false, '0x5', 1, 7, 8, 9, 10, '0x11');

INSERT INTO marketplace_order(order_hash, offerer, order_value, is_cancelled, is_validated, signature, order_type, start_time, end_time, counter, salt, zone)
VALUES('0x3', '0x3', 4, false, false, '0x5', 1, 7, 8, 9, 10, '0x11');

INSERT INTO marketplace_order(order_hash, offerer, order_value, is_cancelled, is_validated, signature, order_type, start_time, end_time, counter, salt, zone)
VALUES('0x4', '0x3', 4, false, false, '0x5', 1, 7, 8, 9, 10, '0x11');

-- seed offer for order
INSERT INTO marketplace_order_offer(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount)
VALUES('0x2', 1, '0x3', 4, '0x5', 6, 7);

INSERT INTO marketplace_order_offer(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount)
VALUES('0x2', 2, '0x3', 5, '0x5', 6, 7);

INSERT INTO marketplace_order_offer(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount)
VALUES('0x2', 3, '0x3', 6, '0x5', 6, 7);

INSERT INTO marketplace_order_offer(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount)
VALUES('0x3', 4, '0x3', 1, '0x11', 1, 1);

INSERT INTO marketplace_order_offer(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount)
VALUES('0x3', 2, '0x3', 2, '0x11', 1, 1);

INSERT INTO marketplace_order_offer(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount)
VALUES('0x4', 3, '0x3', 3, '0x11', 1, 1);

-- seed order consideration
INSERT INTO marketplace_order_consideration(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount, recipient)
VALUES('0x2', 1, '0x3', 4, '0x5', 6, 7, '0x11');

INSERT INTO marketplace_order_consideration(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount, recipient)
VALUES('0x2', 2, '0x3', 5, '0x5', 6, 7, '0x11');

INSERT INTO marketplace_order_consideration(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount, recipient)
VALUES('0x3', 3, '0x3', 6, '0x5', 6, 7, '0x11');

INSERT INTO marketplace_order_consideration(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount, recipient)
VALUES('0x3', 4, '0x3', 1, '0x11', 1, 1, '0x11');

INSERT INTO marketplace_order_consideration(order_hash, type_number, item_type, token_id, token_address, start_amount, end_amount, recipient)
VALUES('0x4', 5, '0x3', 2, '0x11', 1, 1, '0x11');
