INSERT INTO
    "typesOfItems" ("id", "title", "description")
VALUES
    (1, 'телефоны', 'Рассказать про телефоны'),
    (
        2,
        'компьютеры',
        'Рассказать про компьютеры'
    ),
    (
        3,
        'телевизоры',
        'Рассказать про телевизоры'
    ),
    (
        4,
        'аксессуары',
        'Рассказать про аксессуары'
    );

INSERT INTO
    "items" (
        "id",
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
        1,
        'iPhone 13',
        1,
        123,
        12,
        232,
        'photo1',
        'высота 13мм; ширина 13 мм; длинна 5мм',
        'Apple',
        0
    ),
    (
        2,
        'iPhone 12',
        1,
        1232,
        NULL,
        2312,
        'photo2',
        'высота 13мм; ширина 13 мм; длинна 5мм',
        'Apple',
        1
    ),
    (
        3,
        'iPhone 11',
        1,
        1236,
        NULL,
        2932,
        'photo3',
        'высота 13мм; ширина 13 мм; длинна 5мм',
        'Apple',
        2
    ),
    (
        4,
        'iPhone 10',
        1,
        1238,
        112,
        2320,
        'photo4',
        'высота 13мм; ширина 13 мм; длинна 5мм',
        'Apple',
        3
    );

INSERT INTO
    "users" (
        "id",
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
        NULL,
        NULL,
        'fido',
        'lokers',
        'fido@mail.ru',
        '123456',
        '28734872',
        'avatar.jpg'
    ),
    (
        2,
        NULL,
        NULL,
        NULL,
        NULL,
        'kenny@mail.ru',
        'qwerty123',
        NULL,
        NULL
    ),
    (
        3,
        NULL,
        NULL,
        'kendirck',
        'okatnuga',
        'kendirck@mail.ru',
        'lolkek123',
        '98456472',
        NULL
    ),
    (
        4,
        NULL,
        NULL,
        'james',
        'drisier',
        'james@mail.ru',
        'sheesh123',
        '5654872',
        NULL
    );

INSERT INTO
    "paymentsOfUsers" (
        "id",
        "userID",
        "type",
        "cardNumber",
        "exiryDate"
    )
VALUES
    (1, 1, 'card', '1234567890456689', '09/23'),
    (2, 2, 'card', '7234567890456189', '03/23'),
    (3, 3, 'card', '8934567890456489', '10/24'),
    (4, 4, 'card', '9234567890456709', '07/26');

INSERT INTO
    "addressesOfUsers" ("id", "userID", "city", "address")
VALUES
    (1, 1, 'nyc', 'times square'),
    (2, 2, 'la', 'broadway ave.'),
    (3, 3, 'london', 'fullham'),
    (4, 4, 'msk', 'arbat');

-- INSERT INTO
--     "ratingOfItems" (
--         "id",
--         "itemID",
--         "userID",
--         "pros",
--         "cons",
--         "comment",
--         "quilityRating",
--         "priceRating"
--     )
-- VALUES
--     (
--         1,
--         1,
--         1,
--         'good square',
--         'nothing',
--         'nice',
--         9,
--         3
--     ),
--     (
--         2,
--         2,
--         2,
--         'bad ave.',
--         NULL,
--         NULL,
--         9,
--         3
--     ),
--     (3, 3, 3, NULL, 'asd', '-', 9, 3),
--     (
--         4,
--         4,
--         4,
--         'kek',
--         'kind',
--         'cheap',
--         9,
--         3
--     );

-- INSERT INTO
--     "vendors" ("id", "name", "location")
-- VALUES
--     (1, 'seller', 'berlin'),
--     (2, 'reseller', 'minsk'),
--     (3, 'myseller', 'almata'),
--     (4, 'GGseller', 'msk');

-- INSERT INTO
--     "stockOfItems" ("id", "vendorID", "itemID", "amount")
-- VALUES
--     (1, 1, 1, 1),
--     (2, 2, 2, 2),
--     (3, 3, 3, 9),
--     (4, 4, 4, 0);

INSERT INTO
    "orders" (
        "id",
        "userID",
        "orderStatus",
        "paymentStatus",
        "address"
    )
VALUES
    (1, 1, 'cart', 'paid', 'msk'),
    (2, 2, 'created', 'paid', 'spb'),
    (3, 3, 'created', 'onRecive', 'kzn'),
    (4, 4, 'created', 'onRecive', 'ekat');

INSERT INTO
    "orderItem" (
        "id",
        "itemID",
        "orderID",
        "amount",
        "pricePerUnit"
    )
VALUES
    (1, 1, 1, 1, 123),
    (2, 2, 2, 3543, 234),
    (3, 3, 3, 123, 34534),
    (4, 4, 4, 0, 0);

-- INSERT INTO
--     "admins" ("id", "userID", "permissions")
-- VALUES
--     (1, 1, 'root'),
--     (2, 2, 'dev'),
--     (3, 3, 'admin'),
--     (4, 4, 'dev');