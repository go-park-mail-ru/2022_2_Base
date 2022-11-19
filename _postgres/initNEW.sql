CREATE TABLE products
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(30)    NOT NULL,
    category      VARCHAR(50)             default 'undefined',
    price         DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    discountPrice DECIMAL(10, 2) NULL,
    rating        DECIMAL(2, 1)  NOT NULL DEFAULT 0.0,
    imgSrc        VARCHAR(50)    NOT NULL,
    CHECK ( price > 0 ),
    CHECK ( discountPrice > 0 ),
    constraint validDiscount CHECK ( discountPrice < price ),
    CHECK ( rating >= 0)
);

-- Table products:
-- {id} -> name, category, price, discountPrice, rating, imgSrc

CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    email    VARCHAR(30) NOT NULL unique,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL,
    phone    VARCHAR(15) NULL unique,
    avatar   VARCHAR(30) NULL unique
);

-- Table users:
-- {id} -> email, username, password, phone, avatar

-- CREATE TYPE "OrderStatus" AS ENUM (
--     'cart',
--     'created',
--     'processed',
--     'delivery',
--     'delivered',
--     'received',
--     'returned'
-- );

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    userID INT REFERENCES users (id) ON DELETE CASCADE,
    orderStatus VARCHAR(20) NOT NULL,
    paymentStatus VARCHAR(30) NOT NULL,
    addressID INT REFERENCES address (id) ON DELETE RESTRICT,
    paymentCardID INT REFERENCES payment (id) ON DELETE RESTRICT,
    creationDate TIMESTAMP,
    deliveryDate TIMESTAMP
);

-- Table orders:
-- {id} -> userID, orderStatus, paymentStatus, address, creationDate, deliveryDate

-- CREATE TYPE "PaymentStatus" AS ENUM ('paid', 'onReceive', 'not started');

-- CREATE TYPE "OrderStatus" AS ENUM (
--     'cart',
--     'created',
--     'processed',
--     'delivery',
--     'delivered',
--     'reviewed',
--     'returned'
-- );

CREATE TABLE orderItems
(
    id           SERIAL PRIMARY KEY,
    productID    INT REFERENCES products (id) ON DELETE CASCADE,
    orderID      INT REFERENCES orders (id) ON DELETE CASCADE,
    count        INT            NOT NULL,
    pricePerUnit DECIMAL(10, 2) NOT NULL,
    CHECK ( count > 0 )
);

-- Table orderItems:
-- {id} -> productID, orderID, count
-- Цена товара за еденицу (pricePerUnit) есть в таблице,
-- потому что цена товара может изменится после создания заказа

CREATE TABLE address
(
    id       SERIAL PRIMARY KEY,
    userID   INT REFERENCES users (id) ON DELETE CASCADE,
    city     VARCHAR(50) NOT NULL,
    street   VARCHAR(50) NOT NULL,
    house    VARCHAR(50) NOT NULL,
    flat     VARCHAR(50) NULL,
    priority BOOLEAN     NOT NULL DEFAULT FALSE
);

-- Table address:
-- {id} -> userID, city, street, house, flat, priority

-- CREATE TYPE paymentType AS ENUM ('card');

CREATE TABLE payment
(
    id          SERIAL PRIMARY KEY,
    userID      INT REFERENCES users (id) ON DELETE CASCADE,
    paymentType VARCHAR(50) NOT NULL,
    number      VARCHAR(16) NOT NULL,
    expiryDate  DATE        NOT NULL,
    priority    BOOLEAN     NOT NULL DEFAULT FALSE
);

-- Table payment:
-- {id} -> userID, paymentType, number, expiryDate, priority