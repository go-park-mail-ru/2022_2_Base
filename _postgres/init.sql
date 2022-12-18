
CREATE TABLE products (
    id  SERIAL PRIMARY KEY,
	name VARCHAR (110) NOT NULL,
    category VARCHAR (50),
    price FLOAT NOT NULL,
    nominalPrice FLOAT NOT NULL,
    rating Float DEFAULT 0,
	imgsrc VARCHAR (50)
);

CREATE TABLE users (
    id  SERIAL PRIMARY KEY,
	email VARCHAR (30) NOT NULL,
    username VARCHAR (30) NOT NULL,
	password VARCHAR (80) NOT NULL, 
    phone VARCHAR (15), 
    avatar VARCHAR (30)
);

CREATE TABLE address
(
    id       SERIAL PRIMARY KEY,
    userID   INT REFERENCES users (id) ON DELETE CASCADE,
    city     VARCHAR(50) NOT NULL,
    street   VARCHAR(50) NOT NULL,
    house    VARCHAR(50) NOT NULL,
    flat     VARCHAR(50) NULL,
    priority BOOLEAN     NOT NULL DEFAULT FALSE,
    deleted  BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE TABLE payment
(
    id          SERIAL PRIMARY KEY,
    userID      INT REFERENCES users (id) ON DELETE CASCADE,
    paymentType VARCHAR(50) NOT NULL,
    number      VARCHAR(16) NOT NULL,
    expiryDate  DATE        NOT NULL,
    priority    BOOLEAN     NOT NULL DEFAULT FALSE,
    deleted     BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    userID INT REFERENCES users (id) ON DELETE CASCADE,
    orderStatus VARCHAR(20) NOT NULL,
    paymentStatus VARCHAR(30) NOT NULL,
    addressID INT REFERENCES address (id) ON DELETE RESTRICT,
    paymentCardID INT REFERENCES payment (id) ON DELETE RESTRICT,
    creationDate TIMESTAMP,
    deliveryDate TIMESTAMP,
    promocode VARCHAR(15)
);

CREATE TABLE orderItems (
    id  SERIAL PRIMARY KEY,
    itemID INT NOT NULL,
    orderID INT NOT NULL,
    price FLOAT NOT NULL,
    count INT NOT NULL
);

CREATE TABLE comments (
    id  SERIAL PRIMARY KEY,
    itemID INT REFERENCES products (id) ON DELETE CASCADE,
    userID INT REFERENCES users (id) ON DELETE CASCADE,
    pros VARCHAR(300) NOT NULL, 
    cons VARCHAR(300) NOT NULL,
    comment VARCHAR(300) NOT NULL,
    rating Float
);

CREATE TABLE usedpromocodes (
    id  SERIAL PRIMARY KEY,
    userID INT REFERENCES users (id) ON DELETE CASCADE,
    promocode VARCHAR(15) NOT NULL
);


CREATE TABLE properties (
    id  SERIAL PRIMARY KEY,
    category VARCHAR (80),
    propname1 VARCHAR(80), 
    propname2 VARCHAR(80),
    propname3 VARCHAR(80),
    propname4 VARCHAR(80),
    propname5 VARCHAR(80),
    propname6 VARCHAR(80)
);

CREATE TABLE monitors (
    id  SERIAL PRIMARY KEY,
    itemID INT REFERENCES products (id) ON DELETE CASCADE,
    propdesc1 VARCHAR(150), 
    propdesc2 VARCHAR(150),
    propdesc3 VARCHAR(150),
    propdesc4 VARCHAR(150),
    propdesc5 VARCHAR(150),
    propdesc6 VARCHAR(150)
);

CREATE TABLE phones (
    id  SERIAL PRIMARY KEY,
    itemID INT REFERENCES products (id) ON DELETE CASCADE,
    propdesc1 VARCHAR(150), 
    propdesc2 VARCHAR(150),
    propdesc3 VARCHAR(150),
    propdesc4 VARCHAR(150),
    propdesc5 VARCHAR(150),
    propdesc6 VARCHAR(150)
);

CREATE TABLE computers (
    id  SERIAL PRIMARY KEY,
    itemID INT REFERENCES products (id) ON DELETE CASCADE,
    propdesc1 VARCHAR(150), 
    propdesc2 VARCHAR(150),
    propdesc3 VARCHAR(150),
    propdesc4 VARCHAR(150),
    propdesc5 VARCHAR(150),
    propdesc6 VARCHAR(150)
);

CREATE TABLE tvs (
    id  SERIAL PRIMARY KEY,
    itemID INT REFERENCES products (id) ON DELETE CASCADE,
    propdesc1 VARCHAR(150), 
    propdesc2 VARCHAR(150),
    propdesc3 VARCHAR(150),
    propdesc4 VARCHAR(150),
    propdesc5 VARCHAR(150),
    propdesc6 VARCHAR(150)
);

CREATE TABLE watches (
    id  SERIAL PRIMARY KEY,
    itemID INT REFERENCES products (id) ON DELETE CASCADE,
    propdesc1 VARCHAR(150), 
    propdesc2 VARCHAR(150),
    propdesc3 VARCHAR(150),
    propdesc4 VARCHAR(150),
    propdesc5 VARCHAR(150),
    propdesc6 VARCHAR(150)
);

CREATE TABLE tablets (
    id  SERIAL PRIMARY KEY,
    itemID INT REFERENCES products (id) ON DELETE CASCADE,
    propdesc1 VARCHAR(150), 
    propdesc2 VARCHAR(150),
    propdesc3 VARCHAR(150),
    propdesc4 VARCHAR(150),
    propdesc5 VARCHAR(150),
    propdesc6 VARCHAR(150)
);

CREATE TABLE accessories (
    id  SERIAL PRIMARY KEY,
    itemID INT REFERENCES products (id) ON DELETE CASCADE,
    category VARCHAR (50),
    propdesc1 VARCHAR(150), 
    propdesc2 VARCHAR(150),
    propdesc3 VARCHAR(150),
    propdesc4 VARCHAR(150),
    propdesc5 VARCHAR(150),
    propdesc6 VARCHAR(150)
);

