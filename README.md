# service-oriented-architectures

Гимадиев Рустем Аделевич, группа 216

Система - Трекер задач

# Клиентский сервис

## Подготовка к сборке сервиса

Установка переменных окружения для БД:
```
touch .env
echo "CLIENT_DB_NAME=<your-db-name>" >>.env
echo "CLIENT_DB_USER=<your-db-user>" >>.env
echo "CLIENT_DB_PASSWORD=<your-db-password>" >>.env
```

Опционально, для повышения безопасности нужно измененить ключи шифрования.

Для этого в директории client_service/crypto:
```
rm signature.pem signature.pub

openssl genrsa -des3 -out signature.pem 2048
openssl genrsa -out signature.pem 2048
openssl rsa -in signature.pem -outform PEM -pubout -out signature.pub
```

## Сборка клиентского сервиса

Из корневой директории
```
docker compose build
docker compose up --force-recreate client_service client_db
```

## Обращение к клиентскому сервису

```
# Регистрация
curl -X POST "localhost:8080/register?login=<your-login>&password=<your-password>"
# Авторизация
curl -X POST "localhost:8080/auth?login=<your-login>&password=<your-password>"
# Обновление данных
curl -X PUT "localhost:8080/update?token=<your-token-after-logging-in>" -H "Content-Type: application/json" -d '{"<user-field>": "<value>"}'
```

# Сервис задач

## Подготовка к сборке сервиса

Установка переменных окружения для БД:
```
touch .env
echo "TASK_DB_NAME=<your-db-name>" >>.env
echo "TASK_DB_USER=<your-db-user>" >>.env
echo "TASK_DB_PASSWORD=<your-db-password>" >>.env
```

Опционально, можно самостоятельно сгенерировать код по proto-файлам:
```
rm -rf client_service/proto
rm -rf task_service/proto

protoc -I=proto --go_out=client_service/ --go-grpc_out=client_service/ proto/task_service.proto
protoc -I=proto --go_out=task_service/ --go-grpc_out=task_service/ proto/task_service.proto
```

## Сборка сервиса задач

Из корневой директории
```
docker compose build
docker compose up --force-recreate task_service task_db
```

## Обращение к сервису задач

Во всех запросах используется токен, выданный при регистрации/авторизации.
```
# Создание задачи
curl -X POST "localhost:8080/task?token=<your-token>" -H "Content-Type: application/json" -d '{"<task-field>": "<value>"}'
# Обновление задачи
curl -X PUT "localhost:8080/task?token=<your-token>&task_id=<your-task-id>" -H "Content-Type: application/json" -d '{"<task-field>": "<value>"}'
# Удаление задачи
curl -X DELETE "localhost:8080/task?token=<your-token>&task_id=<your-task-id>"
# Получение задачи
curl -X GET "localhost:8080/task?token=<your-token>&task_id=<required-task-id>"
# Получение списка задач
curl -X GET "localhost:8080/tasks?token=<your-token>&author_login=<required-author-login>&page_index=<required-page-index>&tasks_per_page=<size-of-tasks-page-you-want>"
```
