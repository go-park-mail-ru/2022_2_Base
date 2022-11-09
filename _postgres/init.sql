
CREATE TABLE products (
    id  SERIAL PRIMARY KEY,
	name VARCHAR (30),
    description VARCHAR (50),
    price FLOAT,
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


CREATE TABLE orderTable (
    id  SERIAL PRIMARY KEY,
	userID INT NOT NULL,
    orderStatus VARCHAR (20) NOT NULL,
	paymentStatus VARCHAR (30) NOT NULL, 
    adress VARCHAR (50) NOT NULL
);

CREATE TABLE orderItems (
    id  SERIAL PRIMARY KEY,
    itemID INT NOT NULL,
    orderID INT NOT NULL,
    count INT NOT NULL
);

CREATE TABLE adress (
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
