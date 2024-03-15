# service-oriented-architectures

Гимадиев Рустем Аделевич, группа 216

Система - Трекер задач

# Подготовка к сборке

Установка переменных окружения для БД:
```
touch .env
echo "DB_NAME=<your-db-name>" >>.env
echo "DB_USER=<your-db-user>" >>.env
echo "DB_PASSWORD=<your-db-password>" >>.env
```

Опционально, для повышения безопасности нужно измененить ключи шифрования.

Для этого в директории client_service/crypto:
```
rm signature.pem signature.pub

openssl genrsa -des3 -out signature.pem 2048
openssl genrsa -out signature.pem 2048
openssl rsa -in signature.pem -outform PEM -pubout -out signature.pub
```

# Сборка клиентского сервиса

Из корневой директории
```
docker compose build
docker compose up --force-recreate
```

# Обращение к клиентскому сервису

```
# Регистрация
curl -X POST "localhost:8080/register?login=<your-login>&password=<your-password>"
# Авторизация
curl -X POST "localhost:8080/auth?login=<your-login>&password=<your-password>"
# Обновление данных
curl -X PUT "localhost:8080/update?token=<your-token-after-logging-in>" -H "Content-Type: application/json" -d '{"<user-field>": "<value>"}'
```
