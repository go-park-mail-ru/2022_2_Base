INSERT INTO products (name, category, nominalPrice, price, imgsrc) VALUES
('Монитор Xiaomi Mi 27', 'monitors', 14999, 13999,'https://img.mvideo.ru/Big/30058309bb.jpg'),
('Tecno Spark 8с', 'phones', 12999, 8999, 'https://img.mvideo.ru/Big/30062036bb.jpg'),
('realme GT Master', 'phones', 29999, 21999, 'https://img.mvideo.ru/Big/30058843bb.jpg'),
('Apple iPhone 11', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 10', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 9', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 8', 'phones', 62999, 20000, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 7', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 6', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 5', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 4', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 3', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 2', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 21', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 0', 'phones', 62999, 54999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Системный блок игровой MUST MBM114 (I5-11400F/16GB/512Gb+1TB/RTX3060/NOOS)', 'computers', 109999, 109999, 'https://img.mvideo.ru/Big/30064259bb.jpg'),
('Haier 55 Smart TV MX', 'tvs', 49999, 49999, 'https://img.mvideo.ru/Big/10030234bb.jpg'),
('Смарт-часы HUAWEI Watch GT2 Matte Black', 'watches', 10999, 10999, 'https://img.mvideo.ru/Big/30045872bb.jpg'),
('Apple iPad 10.2 Wi-Fi 64GB Silver (MK2L3)', 'tablets', 27999, 27999, 'https://img.mvideo.ru/Pdb/30064044b.jpg'),
('Геймпад для консоли Xbox Hori Horipad Pro (AB01-001E) Hori', 'accessories', 4699, 4699, 'https://img.mvideo.ru/Big/40075059bb.jpg'),
('Монитор Xiaomi Mi 24', 'monitors', 14999, 14999, 'https://img.mvideo.ru/Big/30058309bb.jpg');

INSERT INTO users (email, username, password) VALUES
('art@art',	'aaa', '12345678');

INSERT INTO address (userID, city, street, house, flat, priority) VALUES
(1, 'default', 'default', 'default', 'default', false);

INSERT INTO payment (userID, paymentType, number, expiryDate, priority) VALUES
(1, 'Card', 'default', '1975-08-19T23:15:30.000Z', false);