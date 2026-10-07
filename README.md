# Task Management API

REST API на Go для управления пользовательскими проектами и задачами. Приложение поддерживает регистрацию и авторизацию, CRUD-операции и проверяет, что пользователь работает только со своими ресурсами.

## Стек

- Go и стандартный пакет `net/http`
- PostgreSQL
- `pgx/v5`
- JWT-аутентификация
- `bcrypt` для хеширования паролей

## Возможности

- регистрация и авторизация пользователей;
- создание, просмотр, обновление и удаление проектов;
- создание, просмотр, обновление и удаление задач;
- привязка проектов к пользователям, а задач — к проектам;
- проверка владельца проекта или задачи перед выполнением операций;
- каскадное удаление связанных данных на уровне PostgreSQL;
- логирование HTTP-запросов через middleware.

## Архитектура

Код разделён на слои:

```text
cmd/api          запуск HTTP-сервера и регистрация маршрутов
internal/handler обработка HTTP-запросов и ответов
internal/service бизнес-логика
internal/repository запросы к PostgreSQL
internal/midleware JWT-аутентификация, проверка доступа и логирование
internal/model   модели данных
internal/auth    создание и проверка JWT
internal/migrations SQL-схема базы данных
```

Поток обработки запроса:

```text
HTTP request -> middleware -> handler -> service -> repository -> PostgreSQL
```

## API

Публичные маршруты:

| Метод | Маршрут | Назначение |
| --- | --- | --- |
| `POST` | `/register` | Регистрация пользователя |
| `POST` | `/login` | Авторизация и получение JWT |

Маршруты ниже требуют JWT:

| Метод | Маршрут | Назначение |
| --- | --- | --- |
| `POST` | `/projects` | Создать проект |
| `GET` | `/projects` | Получить проекты пользователя |
| `GET` | `/projects/{id}` | Получить проект |
| `PUT` | `/projects/{id}` | Обновить проект |
| `DELETE` | `/projects/{id}` | Удалить проект |
| `POST` | `/projects/{projectID}/tasks` | Создать задачу в проекте |
| `GET` | `/projects/{projectID}/tasks` | Получить задачи проекта |
| `GET` | `/tasks/{id}` | Получить задачу |
| `PUT` | `/tasks/{id}` | Обновить задачу |
| `DELETE` | `/tasks/{id}` | Удалить задачу |

JWT передаётся в заголовке авторизации:

```http
Authorization: Bearer <token>
```

## Локальный запуск

### 1. Создайте базу данных

Создайте в PostgreSQL базу `taskmanager`, затем выполните миграцию:

```bash
psql -U postgres -d taskmanager -f internal/migrations/001_init.sql
```

### 2. Проверьте строку подключения

Сейчас параметры подключения указаны в `cmd/api/main.go`:

```text
postgres://postgres:1@localhost:5432/taskmanager
```

Перед запуском замените пользователя и пароль на значения своей локальной установки PostgreSQL.

### 3. Запустите сервер

```bash
go run ./cmd/api
```

API будет доступен по адресу `http://localhost:8080`.

## Примеры запросов

Регистрация:

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"strong-password"}'
```

Авторизация:

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"strong-password"}'
```

Создание проекта:

```bash
curl -X POST http://localhost:8080/projects \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Internship preparation"}'
```

## Что можно улучшить

- вынести строку подключения и секрет JWT в переменные окружения;
- добавить graceful shutdown HTTP-сервера;
- добавить unit- и интеграционные тесты;
- описать API в OpenAPI/Swagger;
- настроить проверку сборки и тестов в GitHub Actions;
- добавить Docker Compose для локального запуска PostgreSQL.
