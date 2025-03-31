# 📰 REST News API
Разработан в соответствии с [техническим заданием](https://gist.github.com/bethrezen/d6f17fbb039a4366fe6baafdf189ff9a).
## 📌 Функциональность

- Регистрация и авторизация пользователей (JWT)
- Изменение новостей
- Подключение к базе данных PostgreSQL
- Миграции через Goose
- Логирование через Zerolog

## 🚀 Запуск проекта

### 1. Клонирование репозитория

```sh
git clone https://github.com/Epicpt/rest-news.git
cd rest-news
```

### 2. Измените .env файл

```sh
# HTTP settings
HTTP_PORT=8080
# Logger
LOG_LEVEL=debug
# PG
PG_POOL_MAX=2
PG_URL=postgres://postgres:password@postgres:5432/news
POSTGRES_DB=news
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
# JWT
JWT_SECRET=my-32-character-ultra-secure-and-ultra-long-secret
```

### 3. Запуск через Docker Compose

```sh
docker-compose up --build
```
### 4. Применение миграций (если нужно вручную)

```sh
docker-compose run --rm goose sh -c "goose -dir /migrations postgres $PG_URL up"
```

### 📡 API Эндпоинты
#### 🔑 Аутентификация

- POST /register – регистрация пользователя
- POST /login – вход (возвращает JWT-токен)

#### 📰 Новости

- GET /news/list – список новостей (требуется аутентификация)
- POST /news/edit:Id – изменить новость (требуется аутентификация)

### 🛠 Используемые технологии

- Golang 1.24
- Fiber
- PostgreSQL
- Goose (миграции)
- Docker & Docker Compose
- JWT