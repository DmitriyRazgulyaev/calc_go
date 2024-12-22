# calc_go

## Описание

`calc_go` - это простой веб-сервис для вычисления математических выражений. Он принимает выражения в виде строки, обрабатывает их и возвращает результат.

## Установка

1. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/DmitriyRazgulyaev/calc_go.git
    ```
2. Перейдите в директорию проекта:
    ```sh
    cd calc_go
    ```
3. Установите зависимости:
    ```sh
    go mod tidy
    ```

## Использование

1. Запустите сервер:
    ```sh
    go run cmd/main.go
    ```
2. Отправьте POST-запрос на `http://localhost:8080/api/v1/calculate/` с JSON-телом, содержащим математическое выражение. Пример запроса:
    ```sh
    curl -X POST http://localhost:8080/api/v1/calculate/ -d '{"expression": "2+2*2"}' -H "Content-Type: application/json"
    ```
3. Сервер вернет результат вычисления.

## Пример ответа

```json
{
    "result": 6
}