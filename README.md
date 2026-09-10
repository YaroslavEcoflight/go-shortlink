# go-shortlink

Микросервисный сервис для создания коротких ссылок с авторизацией и аналитикой действий пользователей.

## Архитектура

```
          ┌─────────────┐
          │   Client    │
          └──────┬──────┘
                 │
          ┌──────┴──────┐
          │             │
        REST           REST
          │             │
          ▼             ▼
┌──────────────────┐  ┌───────────────────┐
│   auth-service   │  │ shortlink-service │
│      :8080       │  │      :8081        │
└────────┬─────────┘  └────────┬──────────┘
         │                     │
         │  gRPC (fire&forget) │
         └──────────┬──────────┘
                    ▼
          ┌─────────────────────┐
          │  analytics-service  │
          │    gRPC :50051      │
          │    HTTP  :8082      │
          └──────────┬──────────┘
                     │
            ┌────────┴────────┐
            ▼                 ▼
     ┌────────────┐    ┌────────────┐
     │ PostgreSQL │    │   Redis    │
     └────────────┘    └────────────┘
```

### Сервисы

| Сервис | Порт | Описание |
|---|---|---|
| auth-service | 8080 | Регистрация, авторизация, JWT токены |
| shortlink-service | 8081 | Создание коротких ссылок, редиректы |
| analytics-service | 8082 (HTTP), 50051 (gRPC) | Сбор и хранение событий |

### Взаимодействие сервисов

- **auth-service** и **shortlink-service** отправляют события в **analytics-service** по gRPC в режиме fire-and-forget (не блокируют основной запрос)
- JWT валидируется локально в каждом сервисе по shared secret — межсервисных запросов для проверки токена нет
- Все сервисы используют общий PostgreSQL и Redis

---

## Запуск

### Требования

- Docker + Docker Compose

### 1. Настройка переменных окружения

```bash
cp .env.example .env
```

Отредактируй `.env` — как минимум измени `JWT_SECRET`:

```env
JWT_SECRET=your-secret-key
```

### 2. Запуск

```bash
docker compose up -d --build
```

### 3. Остановка

```bash
docker compose down
```

---

## Переменные окружения

| Переменная | Описание | Пример |
|---|---|---|
| `POSTGRES_USER` | Пользователь БД | `postgres` |
| `POSTGRES_PASSWORD` | Пароль БД | `postgres` |
| `POSTGRES_DB` | Имя базы данных | `auth` |
| `POSTGRES_PORT` | Порт PostgreSQL | `5432` |
| `POSTGRES_SSLMODE` | SSL режим | `disable` |
| `REDIS_PORT` | Порт Redis | `6379` |
| `JWT_SECRET` | Секрет для подписи JWT | `change-me` |
| `AUTH_HTTP_PORT` | Порт auth-service | `8080` |
| `AUTH_APP_VERSION` | Версия auth-service | `0.1.0` |
| `SHORTLINK_HTTP_PORT` | Порт shortlink-service | `8081` |
| `SHORTLINK_APP_VERSION` | Версия shortlink-service | `0.1.0` |
| `ANALYTICS_HTTP_PORT` | HTTP порт analytics-service | `8082` |
| `ANALYTICS_GRPC_PORT` | gRPC порт analytics-service | `50051` |
| `ANALYTICS_ADDR` | Адрес analytics-service для клиентов | `analytics-service:50051` |
| `ANALYTICS_APP_VERSION` | Версия analytics-service | `0.1.0` |
| `LOG_LEVEL` | Уровень логирования всех сервисов | `info` |

---

## API

### auth-service `http://localhost:8080`

#### Регистрация
```http
POST /api/v1/auth
Content-Type: application/json

{
  "username": "john",
  "email": "john@example.com",
  "password": "secret"
}
```

#### Вход
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "secret"
}
```
Ответ:
```json
{
  "access_token": "<jwt>",
  "refresh_token": "<token>",
  "expires_in": 1234567890
}
```

#### Выход
```http
POST /api/v1/auth/logout
Content-Type: application/json

{ "refresh_token": "<token>" }
```

#### Обновление токена
```http
POST /api/v1/auth/refresh
Content-Type: application/json

{ "refresh_token": "<token>" }
```

#### Валидация токена
```http
POST /api/v1/auth/validate
Content-Type: application/json

{ "access_token": "<jwt>" }
```

---

### shortlink-service `http://localhost:8081`

#### Создать короткую ссылку
```http
POST /api/v1/link/shorten
Authorization: Bearer <access_token>
Content-Type: application/json

{ "url": "https://example.com" }
```
Ответ:
```json
{ "Id": 1, "Code": "1", "Url": "https://example.com" }
```

#### Редирект
```http
GET /:code
```
Возвращает `302` на оригинальный URL.

#### Удалить ссылку
```http
DELETE /api/v1/link/:code
Authorization: Bearer <access_token>
```

---

### analytics-service `http://localhost:8082`

#### События пользователя за день
```http
GET /api/v1/events?user_id=<id>&date=2026-09-10
```
Ответ:
```json
[
  {
    "ID": 1,
    "EventType": "login",
    "Status": "success",
    "ErrorCode": "",
    "UserID": "uuid",
    "IP": "192.168.1.1",
    "OccurredAt": "2026-09-10T08:00:00Z"
  }
]
```

**Типы событий:** `register`, `login`, `logout`, `refresh`, `validate`, `shorten`, `delete`

**Статусы:** `success`, `failure`

---

## Структура проекта

```
go-shortlink/
├── auth-service/        # Авторизация и управление пользователями
├── shortlink-service/   # Короткие ссылки
├── analytics-service/   # Аналитика событий
├── proto/               # Общий .proto файл
├── docker-compose.yml
└── .env.example
```

Каждый сервис — независимый Go-модуль с собственным `go.mod` и `Dockerfile`, следует принципам чистой архитектуры (domain / usecase / infrastructure / transport).
