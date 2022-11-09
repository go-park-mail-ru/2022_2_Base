INSERT INTO products (name, description, price, discountPrice, rating, imgsrc) VALUES
('Монитор Xiaomi Mi 27', 'good', 14999, 13999, 4, 'https://img.mvideo.ru/Big/30058309bb.jpg'),
('Монитор Xiaomi Mi 27', 'good', 14999, 13999, 4, 'https://img.mvideo.ru/Big/30058309bb.jpg'),
('A', 22);

INSERT INTO users (email, username, password) VALUES
('art@art',	'aaa', '12345678'),
('art2@art',	'bbb', '12345678');

INSERT INTO adress (userID, city, street, house, priority) VALUES
(1, 'Moscow', 'IZM', '123', false);

INSERT INTO payment (userID, type, number, expiryDate, priority) VALUES
(1, 'Card', '22222', '3333', false);