
CREATE TABLE products (
    id  SERIAL PRIMARY KEY,
	name VARCHAR (30) NOT NULL,
    category VARCHAR (50),
    price FLOAT NOT NULL,
    discountPrice FLOAT,
    rating Float,
	imgsrc VARCHAR (50)
);

CREATE TABLE users (
    id  SERIAL PRIMARY KEY,
	email VARCHAR (30) NOT NULL,
    username VARCHAR (30) NOT NULL,
	password VARCHAR (30) NOT NULL, 
    phone VARCHAR (15), 
    avatar VARCHAR (30)
);

CREATE TABLE address (
    id  SERIAL PRIMARY KEY,
    userID INT NOT NULL,
    city VARCHAR (50) NOT NULL,
    street VARCHAR (50) NOT NULL,
    house VARCHAR (50) NOT NULL,
    priority BOOL NOT NULL
);

CREATE TABLE payment (
    id  SERIAL PRIMARY KEY,
    userID INT NOT NULL,
    type VARCHAR (50) NOT NULL,
    number VARCHAR (50) NOT NULL,
    expiryDate VARCHAR (50) NOT NULL,
    priority BOOLEAN NOT NULL
);

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

CREATE TABLE orderItems (
    id  SERIAL PRIMARY KEY,
    itemID INT NOT NULL,
    orderID INT NOT NULL,
    count INT NOT NULL
);
