
CREATE TABLE IF NOT EXISTS products (
    id  SERIAL PRIMARY KEY,
	name VARCHAR (30),
    description VARCHAR (50),
    price FLOAT,
    discountPrice FLOAT,
    rating Float,
	imgsrc VARCHAR (50)
);

CREATE TABLE IF NOT EXISTS users (
    id  SERIAL PRIMARY KEY,
	email VARCHAR (30) NOT NULL,
    username VARCHAR (30) NOT NULL,
	password VARCHAR (30) NOT NULL, 
    phone VARCHAR (15), 
    avatar VARCHAR (30)
);


CREATE TABLE IF NOT EXISTS orderTable (
    id  SERIAL PRIMARY KEY,
	userID INT NOT NULL,
    items INTEGER[] NOT NULL,
    orderStatus VARCHAR (20) NOT NULL,
	paymentStatus VARCHAR (30) NOT NULL, 
    adress VARCHAR (50) NOT NULL
);
