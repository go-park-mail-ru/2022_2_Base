# Выбор сущности
Для тестирования была выбрана сущность заказа, поскольку она является ключевой для маркетплейса
# Процесс тестирования
Тестирование происходит с помощью утилиты [vegeta](https://github.com/tsenart/vegeta). Данные для этой программы предоставляет [genTestCase.sh](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/genTestCase.sh), который для сушествующего пользователя в корзину добавляется от 1 до 10 случайных товаров и создается заказ.
Далее запрашиваем созданные заказы GET запросом с помощью той же утилиты и [genGetTest.sh](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/genGetTest.sh). Генерация пользователей происходит с помощью [genUser.sh](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/genUsers.sh), добавление адреса доставки с использванием [curlAddAddress.sh](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/curlAddAddress.sh), получение ID адреса с выполнением [curlGetAddressID.sh](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/curlGetAddressID.sh).

Для подготовки данных
```console
$ bash genItems.sh
$ bash genUsers.sh
```

Для запуска POST makeorder тестирования
```console
$ make run-post-test
```

Для запуска GET order тестирования
```console
$ make run-get-test
```
## Первая итерация
#### Результаты
![image](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/1/hist1.png)
![image](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/1/plot1.png)
RPS: 450, время выполнения: 2223 с, запросов выполнено 1000351
```json
"status_codes":
{
    "0":14279,
    "200":245660,
    "400":218,
    "401":37207,
    "500":702987
}
```
#### Анализ
#### Изменения
