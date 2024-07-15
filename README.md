# ATM-API

## Описание

ATM-API - это REST API, имитирующий работу банкомата, написанный на Go. Приложение поддерживает операции по созданию аккаунтов, пополнению баланса, снятию средств и проверке баланса. Все операции логируются в консоль.

## Использование

### Эндпоинты API

- Создание нового аккаунта `POST /accounts`.
- Пополнение баланса `POST /accounts/{id}/deposit`.
- Снятие средств `POST /accounts/{id}/withdraw`.
- Проверка баланса `GET /accounts/{id}/balance `.
