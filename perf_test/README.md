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
RPS: 450, запросов выполнено 1000351, успешных запросов 24%
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
#### Изменения
Много 500 кодов, попробуем добавить индексы orderItems (itemID), orderItems (orderID) и снизить RPS. Так же сменим db drive с pgx на pgx pool.
## Вторая итерация
#### Результаты
![image](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/3/hist3.png)
![image](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/3/plot3.png)
RPS: 300, запросов выполнено 666901, успешных запросов 32%
```json
"status_codes":
{
    "0":7240,
    "200":219351,
    "400":87,
    "401":23100,
    "500":417123
}
```
Количество успешных запросов значительно улучшилось, изменения приведи к положительному результату
#### Изменения
Попробуем [изменить](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/_postgres/postgres.conf) конфгурацию базы данных со стандартной. Так же оптимизируем микросервис заказов
## Третья итерация
#### Результаты
POST:
![image](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/5/hist5.png)
![image](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/5/plot5.png)
RPS: 250, запросов выполнено 1000000, успешных запросов 34%
```json 
"status_codes":
{
    "0":9343,
    "200":343906,
    "400":120,
    "401":35959,
    "500":610672
}
```
POST:
![image](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/5get/hist5.png)
![image](https://github.com/go-park-mail-ru/2022_2_Base/blob/Base-5.2_srs_3/perf_test/5get/plot5.png)
RPS: 450, запросов выполнено 1000351, успешных запросов 27%
GET:
```json 
"status_codes":
{
    "0":15517,
    "200":278271,
    "500":706563
}
```
#### Результаты
Получили небольшое улучшение, процент успешных запросов вырос на 2%
## Итог
Адекватных результатов по RPS и проценту успешных запросов добиться не удалось,
причинами считаем огранниченость технических характеристик серера (single core CPU, 4gb RAM). Так же проблемы создают микросервисы, они не выдерживают такой нагрузки.