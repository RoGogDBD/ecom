# TODO API Server

HTTP-сервер на Go для управления задачами (TODO list) с поддержкой CRUD операций.

## Описание

Сервер реализует REST API для управления задачами с использованием только стандартной библиотеки Go. Данные хранятся в памяти приложения.

## Функциональность

- Создание новой задачи
- Получение списка всех задач
- Получение задачи по ID
- Обновление задачи
- Удаление задачи
- Валидация входных данных
- Логирование всех запросов
- Graceful shutdown

## API Endpoints

| Метод  | Путь          | Описание                    |
|--------|---------------|-----------------------------|
| POST   | /todos        | Создать новую задачу        |
| GET    | /todos        | Получить список всех задач  |
| GET    | /todos/{id}   | Получить задачу по ID       |
| PUT    | /todos/{id}   | Обновить задачу             |
| DELETE | /todos/{id}   | Удалить задачу              |

### Структура задачи

```json
{
  "id": 1,
  "title": "Название задачи",
  "description": "Описание задачи",
  "completed": false
}
```

### Примеры запросов

**Создание задачи:**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Купить молоко","description":"В магазине у дома","completed":false}'
```

**Получение всех задач:**
```bash
curl http://localhost:8080/todos
```

**Получение задачи по ID:**
```bash
curl http://localhost:8080/todos/1
```

**Обновление задачи:**
```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Купить молоко и хлеб","description":"В магазине у дома","completed":true}'
```

**Удаление задачи:**
```bash
curl -X DELETE http://localhost:8080/todos/1
```

## Требования

- Go 1.25 или выше
- Docker и Docker Compose (для запуска в контейнере)
- Make (опционально, для удобства)

## Установка и запуск

### Локальный запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/RoGogDBD/ecom.git
cd ecom
```

2. Установите зависимости:
```bash
go mod download
```

3. Запустите приложение:
```bash
go run ./cmd/server
```

Или с использованием Makefile:
```bash
make run
```

Сервер запустится на `http://localhost:8080`

### Запуск с настройкой конфигурации

Создайте или отредактируйте файл `config.json`:
```json
{
  "server": {
    "host": "0.0.0.0",
    "port": 8080
  }
}
```

Запустите с указанием файла конфигурации:
```bash
go run ./cmd/server -config config.json
```

Или используйте переменные окружения:
```bash
SERVER_HOST=0.0.0.0 SERVER_PORT=8080 go run ./cmd/server
```

### Запуск через Docker

1. Соберите и запустите контейнер:
```bash
docker compose up -d --build
```

Или используйте Makefile:
```bash
make docker-up
```

2. Просмотр логов:
```bash
docker compose logs -f
```

Или:
```bash
make docker-logs
```

3. Остановка контейнера:
```bash
docker compose down
```

Или:
```bash
make docker-down
```

### Сборка бинарного файла

```bash
go build -o bin/ecom ./cmd/server
./bin/ecom
```

Или с помощью Makefile:
```bash
make build
./bin/ecom
```

## Тестирование

Запуск всех unit-тестов:
```bash
go test ./...
```

Или с помощью Makefile:
```bash
make test
```

Запуск тестов с подробным выводом:
```bash
go test -v ./...
```

Запуск тестов с покрытием кода:
```bash
go test -cover ./...
```

## Структура проекта

```
.
├── api/                    # Swagger документация
├── cmd/
│   └── server/
│       └── main.go        # Точка входа приложения
├── internal/
│   ├── config/            # Конфигурация приложения
│   ├── handler/           # HTTP обработчики и роутинг
│   ├── models/            # Модели данных
│   ├── repository/        # Слой работы с хранилищем
│   └── service/           # Бизнес-логика
├── logs/                  # Директория для логов
├── config.json            # Файл конфигурации
├── docker-compose.yml     # Docker Compose конфигурация
├── Dockerfile            # Dockerfile для сборки образа
├── go.mod                # Go модули
├── go.sum                # Зависимости модулей
├── Makefile              # Команды для сборки и запуска
└── README.md             # Документация
```

## Конфигурация

### Файл конфигурации

Приложение поддерживает загрузку конфигурации из JSON файла. По умолчанию используется `config.json` в корне проекта.

```json
{
  "server": {
    "host": "localhost",
    "port": 8080
  }
}
```

### Переменные окружения

Переменные окружения имеют приоритет над файлом конфигурации:

- `CONFIG` - путь к файлу конфигурации
- `SERVER_HOST` - хост сервера (по умолчанию: localhost)
- `SERVER_PORT` - порт сервера (по умолчанию: 8080)

### Флаги командной строки

- `-config` или `-c` - путь к файлу конфигурации

Пример:
```bash
./bin/ecom -config /path/to/config.json
```

## Обработка ошибок

Сервер возвращает соответствующие HTTP статус-коды:

- `200 OK` - успешное выполнение
- `201 Created` - задача успешно создана
- `400 Bad Request` - ошибка валидации (пустой заголовок, некорректные данные)
- `404 Not Found` - задача не найдена
- `405 Method Not Allowed` - метод не поддерживается
- `409 Conflict` - задача с таким ID уже существует
- `500 Internal Server Error` - внутренняя ошибка сервера

## Особенности реализации

- Использование только стандартной библиотеки Go (для runtime)
- Хранение данных в памяти с использованием sync.RWMutex для безопасности
- Middleware для логирования всех HTTP запросов
- Graceful shutdown с таймаутом 10 секунд
- Валидация входных данных
- Unit-тесты для всех слоев приложения
- Многоступенчатая сборка Docker образа для минимизации размера

## Makefile команды

```bash
make build          # Сборка бинарного файла
make run            # Запуск приложения
make test           # Запуск тестов
make clean          # Очистка собранных файлов
make docker-build   # Сборка Docker образа
make docker-up      # Запуск через Docker Compose
make docker-down    # Остановка контейнеров
make docker-logs    # Просмотр логов
make docker-clean   # Полная очистка Docker ресурсов
```