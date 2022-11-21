INSERT INTO products (name, category, price, discountPrice, rating, imgsrc) VALUES
('Монитор Xiaomi Mi 27', 'monitors', 14999, 13999, 4, 'https://img.mvideo.ru/Big/30058309bb.jpg'),
('Tecno Spark 8с', 'phones', 12999, 8999, 4.5, 'https://img.mvideo.ru/Big/30062036bb.jpg'),
('realme GT Master', 'phones', 29999, 21999, 4.3, 'https://img.mvideo.ru/Big/30058843bb.jpg'),
('Apple iPhone 11', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 10', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 9', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 8', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 7', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 6', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 5', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 4', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 3', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 2', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 21', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 0', 'phones', 62999, 54999, 5, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Монитор Xiaomi Mi 24', 'monitors', 14999, 13999, 4, 'https://img.mvideo.ru/Big/30058309bb.jpg');

INSERT INTO users (email, username, password) VALUES
('art@art',	'aaa', '12345678');

INSERT INTO address (userID, city, street, house, priority) VALUES
(1, 'default', 'default', 'default', false);

INSERT INTO payment (userID, paymentType, number, expiryDate, priority) VALUES
(1, 'Card', 'default', '1975-08-19T23:15:30.000Z', false);