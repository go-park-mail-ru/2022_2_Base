-- CREATE TYPE categoryOfProduct AS ENUM ('phones', 'tablets');
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    category VARCHAR(50) DEFAULT 'undefined',
    price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    discountPrice DECIMAL(10, 2) NULL,
    rating DECIMAL(2, 1) NOT NULL DEFAULT 0.0,
    imgSrc VARCHAR(50) NOT NULL,
    CHECK (price > 0),
    CHECK (discountPrice > 0),
    CHECK (rating >= 0)
);

-- Table products:
-- {id} -> name, category, price, discountPrice, rating, imgSrc
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(30) NOT NULL UNIQUE,
    name VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL,
    phone VARCHAR(15) NULL UNIQUE,
    avatar VARCHAR(30) NULL UNIQUE
);

-- Table users:
-- {id} -> email, name, password, phone, avatar
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
-- CREATE TYPE "PaymentStatus" AS ENUM ('paid', 'onReceive');
-- CREATE TYPE "OrderStatus" AS ENUM (
--     'cart',
--     'created',
--     'processed',
--     'delivery',
--     'delivered',
--     'reviewed',
--     'returned'
-- );
CREATE TABLE orderItems (
    id SERIAL PRIMARY KEY,
    productID INT PRIMARY KEY REFERENCES products (id) ON DELETE CASCADE,
    orderID INT PRIMARY KEY REFERENCES orders (id) ON DELETE CASCADE,
    count INT NOT NULL,
    pricePerUnit DECIMAL(10, 2) NOT NULL,
    CHECK (count > 0)
);

-- Table orderItems:
-- {id, productID, orderID} -> count, pricePerUnit
-- Цена товара за еденицу (pricePerUnit) есть в таблице,
-- потому что цена товара может изменится после создания заказа
CREATE TABLE address (
    id SERIAL PRIMARY KEY,
    userID INT REFERENCES users (id) ON DELETE CASCADE,
    city VARCHAR(50) NOT NULL,
    street VARCHAR(50) NOT NULL,
    house VARCHAR(50) NOT NULL,
    flat VARCHAR(50) NULL,
    isPrimary BOOLEAN NOT NULL DEFAULT FALSE,
    deleted BOOLEAN NULL
);

-- Table address:
-- {id} -> userID, city, street, house, flat, isPrimary, deleted
-- CREATE TYPE paymentType AS ENUM ('card');
CREATE TABLE payment (
    id SERIAL PRIMARY KEY,
    userID INT REFERENCES users (id) ON DELETE CASCADE,
    paymentType VARCHAR(50) NOT NULL,
    number VARCHAR(16) NOT NULL,
    expiryDate DATE NOT NULL,
    isPrimary BOOLEAN NOT NULL DEFAULT FALSE,
    deleted BOOLEAN NULL
);

-- Table payment:
-- {id} -> userID, paymentType, number, expiryDate, isPrimary, deleted