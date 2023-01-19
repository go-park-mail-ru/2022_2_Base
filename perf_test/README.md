# Выбор сущности
Для тестирования была выбрана сущность заказа, поскольку она является ключевой для маркетплейса
# Процесс тестирования
Тестирование происходит с помощью утилиты [vegeta](https://github.com/tsenart/vegeta). Данные для этой программы предоставляет [genTestCase.sh](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/genTestCase.sh), который для сушествующего пользователя в корзину добавляется от 1 до 10 случайных товаров и создается заказ.
Далее запрашиваем созданные заказы GET запросом с помощью той же утилиты и [genGetTest.sh](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/genGetTest.sh). Генерация пользователей происходит с помощью [genUser.sh](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/genUsers.sh), добавление адреса доставки с использванием [curlAddAddress.sh](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/curlAddAddress.sh), получение ID адреса с выполнением [curlGetAddressID.sh](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/curlGetAddressID.sh)
## Первая итерация
#### Результаты
![image](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/1/hist-1.jpg)
![image](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/1/plot-1.png)
RPS: 100, время выполнения: 28 минут
#### Анализ
#### Изменения
