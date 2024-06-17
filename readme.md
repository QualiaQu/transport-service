# Transport Service

Transport Service - это веб-сервис на Go, предоставляющий API для управления транспортными маршрутами и бронирования.

## Установка и запуск

Для запуска проекта с использованием Docker и Docker Compose выполните следующие шаги:


### Шаги для запуска

1. Клонируйте репозиторий:

    ```sh
    git clone https://github.com/QualiaQu/transport-service.git
    cd transport-service
    ```

2. Постройте и запустите контейнеры для PostgreSQL:

    ```sh
    docker-compose up -d --build postgres
    ```

3. Постройте и запустите контейнеры для транспортного сервиса:

    ```sh
    docker-compose up -d --build transport-service
    ```

4. Проверьте состояние контейнеров:

    ```sh
    docker-compose ps
    ```

5. После запуска, сервис будет доступен по адресу `http://localhost:8099`.

## Документация API

Документация API, сгенерированная с использованием Swagger, доступна по следующему адресу:

[localhost:8099/swagger/index.html](http://localhost:8099/swagger/index.html)

