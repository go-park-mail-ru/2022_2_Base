
INSERT INTO users (email, username, password) VALUES
('art@art',	'aaa', 'QmFzZTIwMjLCiyMWgeZmxrfD2Wq5LYvkasDYAL3GtKvkim7P1uORBg');

INSERT INTO address (userID, city, street, house, flat, priority) VALUES
(1, 'default', 'default', 'default', 'default', false);

INSERT INTO payment (userID, paymentType, number, expiryDate, priority) VALUES
(1, 'Card', 'default', '1975-08-19T23:15:30.000Z', false);

INSERT INTO products (name, category, nominalPrice, price, imgsrc) VALUES
('Монитор Xiaomi Mi 27', 'monitors', 14999, 13999,'https://img.mvideo.ru/Big/30058309bb.jpg'),
('Монитор Xiaomi Mi 23.8', 'monitors', 11999, 11999, 'https://img.mvideo.ru/Big/30054775bb.jpg'),
('Монитор HUAWEI GT 27', 'monitors', 25999, 24999, 'https://img.mvideo.ru/Pdb/30059512b.jpg'),
('Монитор HUAWEI SE SSN-24', 'monitors', 10499, 10499, 'https://img.mvideo.ru/Big/30063630bb.jpg'),

('Tecno Spark 8с', 'phones', 12999, 6999, 'https://img.mvideo.ru/Big/30062036bb.jpg'),
('realme GT Master', 'phones', 29999, 29999, 'https://img.mvideo.ru/Big/30058843bb.jpg'),
('Apple iPhone 11', 'phones', 62999, 48999, 'https://img.mvideo.ru/Big/30063237bb.jpg'),
('Apple iPhone 13', 'phones', 79999, 66199, 'https://img.mvideo.ru/Big/30063534bb.jpg'),
('Apple iPhone 13 Pro', 'phones', 122999, 87699, 'https://img.mvideo.ru/Pdb/30063566b.jpg'),
('Apple iPhone 13 Pro Max', 'phones', 120999, 104399, 'https://img.mvideo.ru/Big/30063191bb.jpg'),
('Apple iPhone 13 mini', 'phones', 72999, 59799, 'https://img.mvideo.ru/Big/30063553bb.jpg'),
('Apple iPhone 12', 'phones', 69999, 58499, 'https://img.mvideo.ru/Big/30063499bb.jpg'),
('Apple iPhone 12 Pro Max', 'phones', 109999, 88000, 'https://img.mvideo.ru/Big/30052915bb.jpg'),
('Apple iPhone 12 mini', 'phones', 69999, 60999, 'https://img.mvideo.ru/Big/30063517bb.jpg'),
('Samsung Galaxy S22 Ultra', 'phones', 84999, 84999, 'https://img.mvideo.ru/Pdb/30066408b.jpg'),
('Samsung Galaxy Z Fold4', 'phones', 113999, 112999, 'https://img.mvideo.ru/Big/30065099bb.jpg'),

('Системный блок игровой MUST MBM114 (I5-11400F/16GB/512Gb+1TB/RTX3060/NOOS)', 'computers', 96999, 96999, 'https://img.mvideo.ru/Big/30064259bb.jpg'),
('Системный блок игровой MSI MAG Codex 5 11SC-1037XRU', 'computers', 93899, 74999, 'https://img.mvideo.ru/Big/30063347bb.jpg'),
('Системный блок HP Slim Desktop S01-aF0009ur 24U68EA', 'computers', 12799, 12799, 'https://img.mvideo.ru/Big/30058598bb.jpg'),
('Системный блок игровой Acer NITRO 50 N50-640 (DG.E2VER.003)', 'computers', 136999, 136999, 'https://img.mvideo.ru/Big/30063373bb.jpg'),

('Haier 55 Smart TV MX', 'tvs', 41999, 39999, 'https://img.mvideo.ru/Big/10030234bb.jpg'),
('Яндекс 43', 'tvs', 29999, 26999, 'https://img.mvideo.ru/Big/10031656bb.jpg'),
('Toshiba 50C350KE', 'tvs', 44999, 29999, 'https://img.mvideo.ru/Big/10030415bb.jpg'),
('Haier 65 Smart TV AX Pro', 'tvs', 109999, 79999, 'https://img.mvideo.ru/Big/10030671bb.jpg'),

('Смарт-часы HUAWEI Watch GT2 Matte Black', 'watches', 10999, 9999, 'https://img.mvideo.ru/Big/30045872bb.jpg'),
('Смарт-часы Samsung Galaxy Watch5 40mm Pink Gold', 'watches', 21999, 17999, 'https://img.mvideo.ru/Big/30065408bb.jpg'),
('Смарт-часы Xiaomi Redmi Watch 2 Lite Black (BHR5436GL)', 'watches', 4599, 3999, 'https://img.mvideo.ru/Big/30061134bb.jpg'),
('Смарт-часы Apple Watch Series 8 41mm Starlight Aluminium Sport S/M', 'watches', 39999, 38999, 'https://img.mvideo.ru/Big/30066237bb.jpg'),

('Apple iPad 10.2 Wi-Fi 64GB Silver (MK2L3)', 'tablets', 27999, 27999, 'https://img.mvideo.ru/Pdb/30064044b.jpg'),
('Samsung Galaxy Tab A8 10.5 LTE 64GB Gray (SM-X205)', 'tablets', 21999, 21999, 'https://img.mvideo.ru/Pdb/30064063b.jpg'),
('HUAWEI MatePad 10.4 4/128GB LTE Grey (BAH4-L09)', 'tablets', 19999, 19999, 'https://img.mvideo.ru/Pdb/30064395b.jpg'),
('Digma 10 A502 3G', 'tablets', 4999, 4999, 'https://img.mvideo.ru/Big/400003048bb.jpg'),

('Сетевой фильтр ЭРА USF-4es-1.5m-B', 'accessories', 599, 599, 'https://img.mvideo.ru/Big/50131693bb.jpg'),
('Сетевой фильтр ЭРА USF-4es-1.5m-W', 'accessories', 599, 599, 'https://img.mvideo.ru/Big/50131692bb.jpg'),
('Салфетки для комп. техники Code Влажные 20шт. (СС-120)', 'accessories', 99, 99, 'https://img.mvideo.ru/Big/50144659bb.jpg'),
('Чистящее средство для компьютерной техники Home Protect 250мл (HP800033)', 'accessories', 299, 299, 'https://img.mvideo.ru/Big/50143133bb.jpg'),
('Чехол TFN Apple iPhone 13 Fade MagSafe Black', 'accessories', 1499, 999, 'https://img.mvideo.ru/Big/50163252bb.jpg'),
('Чехол TFN Apple iPhone 13 Hard MagSafe Clear', 'accessories', 1499, 1499, 'https://img.mvideo.ru/Big/50163244bb.jpg'),
('Внешний аккумулятор TFN PowerAid 10000мАч Black (TFN-PB-278-BK)', 'accessories', 999, 999, 'https://img.mvideo.ru/Big/50171468bb.jpg'),
('Внешний аккумулятор Carmega 20000mAh Charge 20 black (CAR-PB-202-BK)', 'accessories', 1499, 1299, 'https://img.mvideo.ru/Big/50170521bb.jpg'),
('Флеш-диск Hikvision 64GB USB 2.0 (HS-USB-M200 64G)', 'accessories', 799, 359, 'https://img.mvideo.ru/Pdb/50172829b.jpg'),
('Флеш-диск Netac 64GB U903 USB2.0 (NT03U903N-064G-20BK)', 'accessories', 799, 359, 'https://img.mvideo.ru/Big/50165477bb.jpg'),
('Wi-Fi роутер TP-Link TL-WR841N V14.0', 'accessories', 1349, 1349, 'https://img.mvideo.ru/Big/50123335bb.jpg'),
('Приемник Wi-Fi TP-Link Archer T2U Plus', 'accessories', 1299, 1299, 'https://img.mvideo.ru/Big/50130210bb.jpg'),
('Антенна телевизионная комнатная One For All Value Line SV9143', 'accessories', 2299, 2299, 'https://img.mvideo.ru/Pdb/50036636b.jpg'),
('Антенна телевизионная комнатная Рэмо BAS-5341-DX Мицар', 'accessories', 1899, 1899, 'https://img.mvideo.ru/Pdb/50050250b.jpg'),
('Антенна телевизионная комнатная Рэмо BAS-5341-DX Мицар', 'accessories', 1899, 1899, 'https://img.mvideo.ru/Pdb/50050250b.jpg'),
('Сетевой фильтр ЭРА USF-4es-1.5m-B', 'accessories', 599, 599, 'https://img.mvideo.ru/Big/50131693bb.jpg'),
('Сетевой фильтр ЭРА USF-4es-1.5m-W', 'accessories', 599, 599, 'https://img.mvideo.ru/Big/50131692bb.jpg'),
('Беспроводное зарядное устройство RIVACASE RivaPower VA4912 Fast Charger', 'accessories', 1299, 999, 'https://img.mvideo.ru/Pdb/50053449b.jpg'),
('Беспроводное зарядное устройство TFN Rapid 15W (TFN-QI03)', 'accessories', 2999, 2499, 'https://img.mvideo.ru/Big/50152389bb.jpg'),
('Беспроводное зарядное устройство Rombica Neo Energy Brown (NQ-00230)', 'accessories', 1999, 1599, 'https://img.mvideo.ru/Big/50132466bb.jpg'),
('Беспроводное зарядное устройство Deppa Qi 3 в 1,Galaxy Watch,Galaxy Buds, 17,5 Вт,Black', 'accessories', 2199, 1799, 'https://img.mvideo.ru/Big/50149730bb.jpg'),
('Стилус Apple Pencil (1-го поколения) (MK0C2ZM/A)', 'accessories', 9999, 9999, 'https://img.mvideo.ru/Pdb/50043834b.jpg'),
('Стилус SwitchEasy EasyPencil Pro для iPad 2018/2019/2020', 'accessories', 3999, 3999, 'https://img.mvideo.ru/Big/50131301bb.jpg'),
('Сетевое зарядное устройство TFN Ultra PD 20W White (TFN-WCRPD30W01)', 'accessories', 999, 999, 'https://img.mvideo.ru/Big/50145799bb.jpg'),
('Кабель Lightning TFN Type-C 1.0m TPE white (TFN-CLIGC1MTPEWH)', 'accessories', 499, 499, 'https://img.mvideo.ru/Pdb/50166520b.jpg');

INSERT INTO properties (category, propname1, propname2, propname3, propname4, propname5, propname6) VALUES
('monitors', 'Экран', 'Яркость', 'Частота обновления', 'Динамическая контрастность',' Интерфейс связи с ПК', 'Время отклика пикселя'),
('phones', 'Экран', 'Процессор', 'Оперативная память(RAM)', 'Встроенная память (ROM)','Основная камера МПикс','Фронтальная камера МПикс'),
('computers', 'Процессор', 'Операционная система','Оперативная память (RAM)','Графический контроллер','Диск','Блок питания'),
('tvs', 'Экран', 'Диагональ экрана', 'Звук','Габаритные размеры (В*Ш*Г)', 'Поддержка Smart TV','Операционная система'),
('watches', 'Диагональ/разрешение', 'Технология дисплея', 'Совместимость', 'Встроенная память (ROM)','Сенсорный экран','Размер'),
('tablets', 'Экран', 'Встроенная память (ROM)', 'Оперативная память (RAM)', 'Тип процессора','Количество ядер','Основная камера МПикс'),
('accessories', 'Гарантия', 'Страна', 'Серия', 'Модель', 'Цвет', 'Материал');

INSERT INTO monitors (itemID, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6) VALUES
(1, '27"/1920x1080 Пикс', '300 кд/кв.м', '75 Гц', '1 000 000:1', 'HDMI', '6 (GTG) мсек'),
(2, '23.8"/1920x1080 Пикс', '250 кд/кв.м', '60 Гц', '1000:1', 'D-Sub; HDMI', '6 (GTG) мсек'),
(3, '27"/2560x1440 Пикс', '350 кд/кв.м', '165 Гц', '4000:1', 'DisplayPort; HDMI', '4 (GTG) мсек'),
(4, '23.8"/1920x1080 Пикс', '250 кд/кв.м', '75 Гц', '1000:1', 'DisplayPort; HDMI', '4 (GTG) мсек');

INSERT INTO phones (itemID, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6) VALUES
(5, '6.6"/720x1612 Пикс', 'UNISOC Tiger T606 2 x 1.6ГГц + 6 x 1.4ГГц', '4 ГБ', '64 ГБ', '13/TOF/AF', '8/TOF'),
(6, '6.43"/2400x1080 Пикс', 'Qualcomm Snapdragon 778G 1 х 2.4ГГц + 3 х 2.2ГГц + 4 х 1.9ГГц', '6 ГБ', '128 ГБ', '64/8/2', '32'),
(7, '6.1"/1792x828 Пикс', 'A13 Bionic', '4 ГБ', '128 ГБ', '12/12', '12'),
(8, '6.1"/2532x1170 Пикс', 'A15 Bionic', '4 ГБ', '128 ГБ', '12/12', '12'),
(9, '6.1"/2532x1170 Пикс', 'A15 Bionic', '4 ГБ', '128 ГБ', '12/12/12', '12'),
(10, '6.7"/2778x1284 Пикс', 'A15 Bionic', '4 ГБ', '256 ГБ', '12/12/12', '12'),
(11, '5.4"/2340x1080 Пикс', 'A15 Bionic', '4 ГБ', '128 ГБ', '12/12', '12'),
(12, '6.1"/2532x1170 Пикс', 'A14 Bionic', '4 ГБ', '128 ГБ', '12/12', '12'),
(13, '6.7"/2778x1284 Пикс', 'A14 Bionic', '4 ГБ', '128 ГБ', '12/12/12/LiDAR', '12'),
(14, '5.4"/2340x1080 Пикс', 'A14 Bionic', '4 ГБ', '128 ГБ', '12/12', '12'),
(15, '6.8"/3088x1440 Пикс', 'Qualcomm Snapdragon 8 Gen 1 1x3.0 ГГц + 3x2.5 ГГц + 4x1.8 ГГц', '12 ГБ', '256 ГБ', '12/108/10/10/TOF', '40'),
(16, '7.6"/1812x2176 Пикс', 'Qualcomm 1x3.18 + 3x2.7 + 4x2 ГГц', '12 ГБ', '256 ГБ', '50/12/10', '10');

INSERT INTO computers (itemID, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6) VALUES
(17, 'Intel Core i5 11400F 2.6 ГГц; 6 ядер; максимальная тактовая частота 4.4 ГГц', 'не установлена', '16 ГБ', 'GeForce RTX 3060', 'HDD 1 ТБ; SSD 500 ГБ', '700 Вт'),
(18, 'Intel Core i5 11400F 2.6 ГГц; 6 ядер; максимальная тактовая частота 4.4 ГГц', 'не установлена', '16 ГБ', 'GeForce RTX 2060', 'SSD 512 ГБ', '500 Вт'),
(19, 'AMD Athlon Silver 3050U 2.3 ГГц; 2 ядра; максимальная тактовая частота 3.2 ГГц', 'DOS', '4 ГБ', 'Radeon Graphics', 'SSD 128 ГБ', null),
(20, 'Intel Core i5 12400F 2.5 ГГц; 6 ядер; максимальная тактовая частота 4.4 ГГц', 'Windows 11 Домашняя', '16 ГБ', 'GeForce RTX 3060 Ti', 'SSD 512 ГБ', '500 Вт');

INSERT INTO tvs (itemID, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6) VALUES
(21, '55"/3840x2160 Пикс', '139.6 см', 'NICAM стерео', '70.9*122.9*6.1 см', 'Да', 'Android 9.0'),
(22, '43"/3840x2160 Пикс', '109 см', 'Dolby Audio', '624*955*78 см', 'Да', 'Яндекс.ТВ+Android 9.0'),
(23, '50"/3840x2160 Пикс', '126 см', 'DTS Studio Sound', '69.2*111.7*22.6 см', 'Да', 'Vidaa'),
(24, '65"/3840x2160 Пикс', '139.6 см', 'NICAM стерео', '85.2*144.8*6.7 см', 'Да', 'Android 11');

INSERT INTO watches (itemID, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6) VALUES
(25, '1.39"/454х454 Пикс', 'AMOLED', 'Android 4.4 и выше, iOS 9.0 и выше', '4 ГБ', 'Да', '45.9 мм'),
(26, null, null, 'Android', '16 ГБ', 'Да', '40 мм'),
(27, '1.55"/360x320 Пикс', 'TFT', 'Android/ iOS', null, 'Да', '41.2 мм'),
(28, null, null, null, '32 ГБ', 'Да', '41 мм');

INSERT INTO tablets (itemID, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6) VALUES
(29, '10.2"/1620x2160 Пикс', '64 ГБ', '4 ГБ', 'A13 Bionic', '6', '8'),
(30, '10.5"/1920x1200 Пикс', '64 ГБ', '4 ГБ', 'Tiger T618', '8', '8'),
(31, '10.4"/1200x2000 Пикс', '128 ГБ', '4 ГБ', 'Kirin 710', '8', '13'),
(32, '10.1"/1280x800 Пикс', '16 ГБ', '1 ГБ', 'SC7731E', '4', '2');

INSERT INTO accessories (itemID, category, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6) VALUES
(33, 'monitors', '1 год', 'Китай', 'USF-M', null, 'черный', 'полипропилен'),
(34, 'monitors', '1 год', 'Китай', 'USF-M', null, 'белый', 'полипропилен'),
(35, 'monitors', null, 'Россия', null, null, 'белый', 'крепированная бумага'),
(36, 'monitors', null, 'Россия', null, null, 'прозрачный', 'пластик'),
(37, 'phones', '6 месяцев', 'Китай', 'Fade MagSafe Black', 'TFN-SС-IP13FMSTBK', 'черный', 'силикон'),
(38, 'phones', '6 месяцев', 'Китай', 'Hard MagSafe Clear', 'TFN-SС-IP13HMSTR', 'прозрачный', 'силикон'),
(39, 'phones', '1 год', 'Китай', 'TFN-PB-278-BK', null, 'черный', 'пластик'),
(40, 'phones', '1 год', 'Китай', null, null, 'черный', 'пластик'),
(41, 'computers', '1 год', 'Китай', 'M200', null, 'серый', 'авиационный алюминий'),
(42, 'computers', '3 года', 'Китай', 'U903', null, 'черный', 'пластик'),
(43, 'computers', '3 года', 'Китай', null, null, 'белый', null),
(44, 'computers', '3 года', 'Китай', null, null, 'черный', null),
(45, 'tvs', '1 год', 'Китай', 'SV9143', null, 'черный', 'пластик'),
(46, 'tvs', '1 год', 'Россия', 'BAS-5341-DX', null, 'черный', 'пластик'),
(47, 'tvs', '1 год', 'Китай', 'USF-M', null, 'черный', 'полипропилен'),
(48, 'tvs', '1 год', 'Китай', 'USF-M', null, 'белый', 'полипропилен'),
(49, 'watches', '2 года', 'Китай', 'VA', 'VA4912', 'белый', 'пластик'),
(50, 'watches', '6 месяцев', 'Китай', 'TFQI03', 'TFN-QI03', 'черный', 'пластик'),
(51, 'watches', '1 год', 'Китай', 'NQ', 'NQ-00230', 'коричневый', 'металл/ пластик'),
(52, 'watches', '1 год', 'Китай', 'Qi', '24011', 'черный', 'пластик'),
(53, 'tablets', '1 год', 'Китай', 'Apple Pencil', 'MK0C2ZM/A', 'белый', 'пластик'),
(54, 'tablets', '1 год', 'Китай', null, null, 'белый', 'алюминий'),
(55, 'tablets', '1 год', 'Китай', 'Ultra', 'TFN-WCRPD30W01', 'белый', 'пластик'),
(56, 'tablets', '3 месяца', 'Китай', 'TPE', null, 'белый', 'TPE');