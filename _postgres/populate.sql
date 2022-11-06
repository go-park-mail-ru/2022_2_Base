INSERT INTO
    "typesOfItems" ("title", "description")
VALUES
    ('телефоны', 'Рассказать про телефоны'),
    ('компьютеры', 'Рассказать про компьютеры'),
    ('телевизоры', 'Рассказать про телевизоры'),
    ('аксессуары', 'Рассказать про аксессуары');

INSERT INTO
    "items" (
        "title",
        "typeID",
        "price",
        "salePrice",
        "weight",
        "photo",
        "size",
        "manufacturer",
        "rating"
    )
VALUES
    (
        'iPhone 13',
        1,
        123,
        12,
        232,
        'photo1',
        'высота: 13мм; ширина: 13 мм; длинна: 5мм',
        'Apple',
        0
    ),
    (
        'iPhone 12',
        1,
        1232,
        NULL,
        2312,
        'photo2',
        'высота: 13мм; ширина: 13 мм; длинна: 5мм',
        'Apple',
        1
    ),
    (
        'iPhone 11',
        1,
        1236,
        NULL,
        2932,
        'photo3',
        'высота: 13мм; ширина: 13 мм; длинна: 5мм',
        'Apple',
        2
    ),
    (
        'iPhone 10',
        1,
        1238,
        112,
        2320,
        'photo4',
        'высота: 13мм; ширина: 13 мм; длинна: 5мм',
        'Apple',
        3
    );

INSERT INTO
    "users" (
        "addressID",
        "paymentID",
        "name",
        "lastname",
        "email",
        "password",
        "phone",
        "avatar"
    )
VALUES
    (
        1,
        1,
        "fido",
        'lokers',
        'fido@mail.ru',
        '123456',
        '28734872',
        'avatar.jpg'
    ),
    (
        2,
        2,
        NULL,
        NULL,
        'kenny@mail.ru',
        'qwerty123',
        NULL,
        NULL
    ),
    (
        3,
        3,
        "kendirck",
        'okatnuga',
        'kendirck@mail.ru',
        'lolkek123',
        '98456472',
        NULL
    ),
    (
        4,
        4,
        "james",
        'drisier',
        'james@mail.ru',
        'sheesh123',
        '5654872',
        NULL
    );

INSERT INTO
    "paymentsOfUsers" ("userID", "type", "cardNumber", "exiryDate")
VALUES
    (1, 'card', 1234567890),
    (2, 'card', 7234567890),
    (3, 'card', 8934567890),
    (4, 'card', 9234567890);

INSERT INTO
    "addressesOfUsers" ("userID", "city", "address")
VALUES
    (1, 'nyc', "times square"),
    (2, 'la', "broadway ave."),
    (3, 'london', "fullham"),
    (4, 'msk', "arbat");

INSERT INTO
    "ratingOfItems" (
        "itemID",
        "userID",
        "pros",
        "cons",
        "COMMENT",
        "quilityRating",
        "priceRating"
    )
VALUES
    (1, 1, "good square", "nothing", "nice", 9, 3),
    (2, 2, "bad ave.", NULL, NULL, 9, 3),
    (3, 3, NULL, "asd", "-", 9, 3),
    (4, 4, "kek", "kind", "cheap", 9, 3);

INSERT INTO
    "vendors" ("name", "location")
VALUES
    ("seller", "berlin"),
    ("reseller", "minsk"),
    ("myseller", "almata"),
    ("GGseller", "msk");

INSERT INTO
    "stockOfItems" ("vendorID", "userID", "amount")
VALUES
    (1, 1, 9),
    (2, 2, 8),
    (3, 3, 7),
    (4, 4, 0);

INSERT INTO
    "orders" (
        "userID",
        "orderStatus",
        "paymentStatus",
        "address"
    )
VALUES
    (1, 'cart', "paid", "msk"),
    (2, 'created', "paid", "spb"),
    (3, 'created', "onRecive", "kzn"),
    (4, 'created', "onRecive", "ekat");

INSERT INTO
    "orderItem" ("itemID", "orderID", "amount", "pricePerUnit")
VALUES
    (1, 1, 1, 1),
    (2, 2, 2, 3543),
    (3, 3, 3, 123),
    (4, 4, 0, 0);

INSERT INTO
    "admins" ("userID", "permissions")
VALUES
    (1, "root"),
    (2, "dev"),
    (3, "admin"),
    (4, "dev");