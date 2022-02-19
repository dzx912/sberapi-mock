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
sberapi-mock start --ignore-validation
```

## Примеры запросов

Cписок доступных API:

```sh
curl http://localhost:8084

(POST) /creation
(POST) /status
```

Создание QR:

```sh
# Запускаем заглушку без проверки запросов 
sberapi-mock start --ignore-validation

# Запрос
curl -X POST http://localhost:8080/creation -d '{}' | jq 

# Ответ
{
  "status": {
    "error_code": "000000",
    "error_description": "Описание ошибки выполнения запроса",
    "order_form_url": "https://sberbank.ru/qr/?uuid=111111111111111111",
    "order_id": "10001000518956637",
    "order_number": "774635526637",
    "order_state": "CREATED",
    "rq_tm": "2005-08-15T15:52:01Z",
    "rq_uid": "ac11cA1CEae1D1111dABf1fD1Bb0acAd"
  }
}
```

## Сборка

```sh 
make build 
```

## Docker

### Запуск

```sh 
docker run --publish 8080:8080 ghcr.io/rige1/sberapi-mock:main
```

### Запуск с параметрами

```sh 
docker run --publish 8080:8080 ghcr.io/rige1/sberapi-mock:main --ignore-validation
```

```sh 
docker run --publish 8080:8080 \
    --volume <absolume_path_to_cert_dir>:/app/cert/ \
    ghcr.io/rige1/sberapi-mock:main \
    --cert /app/cert/server_cert.pem \
    --key /app/cert/server_key.pem \
    --client-cert /app/cert/client_cert.pem 
```