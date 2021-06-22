# sberapi-mock [![build status](https://github.com/rige1/sberapi-mock/actions/workflows/workflow.yml/badge.svg?branch=main)](https://github.com/rige1/sberapi-mock/actions/workflows/workflow.yml)

Сервер заглушка Sber API. Используется для быстрой интеграции и тестирования 

## Что умеет 

* Генерирует и проверяет запросы используя OpenAPI 3
* Имеет поддержку TLS и mTLS

## Примеры использования 

Запуск без аргументов. Сервер будет слушать на порте 8080: 

```sh
sberapi-mock start
```

Запуск с указанием порта: 

```sh
sberapi-mock start --port 8084
```

Запуск с mTLS:

```sh 
sberapi-mock start --cert server_cert.pem --key server_key.pem --client-cert client_cert.pem 
```

Отключение валидации запроса: 
```sh
sberapi-mock start --validate false
```

## Сборка

```sh 
make build 
```