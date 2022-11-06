CREATE TABLE IF NOT EXISTS "typesOfItems" (
    "id" uuid PRIMARY KEY,
    "title" varchar(100) UNIQUE NOT NULL,
    "description" varchar(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS "items" (
    "id" uuid PRIMARY KEY,
    "title" varchar(100) UNIQUE NOT NULL,
    "typeID" uuid REFERENCES "typesOfItems" ("id") ON DELETE CASCADE,
    "price" money NOT NULL,
    "salePrice" money NULL,
    "weight" int NOT NULL,
    "photo" varchar(100) NOT NULL,
    "size" varchar(50) NOT NULL,
    "manufacturer" varchar(50),
    "rating" numeric NULL,
    CHECK ("weight" > 0),
    CHECK (
        "rating" >= 0
        AND "rating" <= 10
    ),
    CHECK ("price" > "salePrice")
);

CREATE TYPE "TypesOfPayment" AS ENUM ('card');

CREATE TABLE IF NOT EXISTS "paymentsOfUsers" (
    "id" uuid PRIMARY KEY,
    -- "userID" uuid REFERENCES "users" ("id") ON DELETE CASCADE,
    "type" "TypesOfPayment" NOT NULL,
    "cardNumber" int2 NOT NULL,
    "exiryDate" varchar(5) NOT NULL
);

CREATE TABLE IF NOT EXISTS "addressesOfUsers" (
    "id" uuid PRIMARY KEY,
    -- "userID" uuid REFERENCES "users" ("id") ON DELETE CASCADE,
    "city" varchar(30) NOT NULL,
    "address" int NOT NULL
);

CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid PRIMARY KEY,
    "addressID" uuid REFERENCES "addressesOfUsers" ("id") NULL,
    "paymentID" uuid REFERENCES "paymentsOfUsers" ("id") NULL,
    "name" varchar(50) NULL,
    "lastname" varchar(50) NULL,
    "email" varchar(50) UNIQUE NOT NULL,
    "password" varchar(30) NOT NULL,
    "phone" varchar(30) UNIQUE NULL,
    "avatar" varchar(100) NULL
);

ALTER TABLE
    "paymentsOfUsers"
ADD
    IF NOT EXISTS "userID" uuid REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE
    "addressesOfUsers"
ADD
    IF NOT EXISTS "userID" uuid REFERENCES "users" ("id") ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS "ratingOfItems" (
    "id" uuid PRIMARY KEY,
    "itemID" uuid REFERENCES "items" ("id") ON DELETE CASCADE,
    "userID" uuid REFERENCES "users" ("id") ON DELETE CASCADE,
    "pros" text NULL,
    "cons" text NULL,
    "comment" text NULL,
    "quilityRating" int2 NOT NULL,
    "priceRating" int2 NOT NULL,
    CHECK (
        "quilityRating" > 0
        AND "quilityRating" <= 10
    ),
    CHECK (
        "priceRating" > 0
        AND "priceRating" <= 10
    )
);

CREATE TABLE IF NOT EXISTS "vendors" (
    "id" uuid PRIMARY KEY,
    "name" varchar(50) NOT NULL,
    "location" varchar(50) NOT NULL
);

CREATE TABLE "stockOfItems" (
    "id" uuid PRIMARY KEY,
    "vendorID" uuid REFERENCES "vendors" ("id") ON DELETE CASCADE,
    "itemID" uuid REFERENCES "items" ("id") ON DELETE CASCADE,
    "amount" int NOT NULL,
    CHECK ("amount" >= 0)
);

CREATE TYPE "PaymentStatus" AS ENUM ('paid', 'onRecive');

CREATE TYPE "OrderStatus" AS ENUM (
    'cart',
    'created',
    'processed',
    'delivery',
    'delivered',
    'reciewed',
    'returned'
);

CREATE TABLE IF NOT EXISTS "orders" (
    "id" uuid PRIMARY KEY,
    "userID" uuid REFERENCES "users" ("id") ON DELETE CASCADE,
    "orderStatus" "PaymentStatus" NOT NULL,
    "paymentStatus" "OrderStatus" NOT NULL,
    "address" varchar(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS "orderItem" (
    "id" uuid PRIMARY KEY,
    "itemID" uuid REFERENCES "items" ("id") ON DELETE CASCADE,
    "orderID" uuid REFERENCES "orders" ("id") ON DELETE CASCADE,
    "amount" int NOT NULL,
    "pricePerUnit" money NOT NULL,
    CHECK ("amount" >= 0)
);

CREATE TYPE "AdminPermissions" AS ENUM ('root', 'dev', 'admin');

CREATE TABLE IF NOT EXISTS "admins" (
    "id" uuid PRIMARY KEY,
    "userID" uuid REFERENCES "users" ("id") ON DELETE CASCADE,
    "permissions" "AdminPermissions" NOT NULL
);