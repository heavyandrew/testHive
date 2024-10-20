# testHive

testHive — это API для управления пользователями и ассетами. С помощью API можно регистрировать пользователей, аутентифицировать их, добавлять, искать, удалять и покупать ассеты. Проект использует Go для разработки сервера и PostgreSQL для хранения данных.

## Сборка проекта

### Конфигурация .env

Создайте файл `.env` в корневом каталоге проекта и добавьте следующую конфигурацию:

```ini
# Конфигурация для базы данных
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=testHive

# Секрет для JWT
# что-то сгенеренное на https://jwtsecret.com/generate
JWT_SECRET=
```
### Запуск
```ini 
make build
```

## Другие команды Makefile
| Команда    |                                              	Описание                                              |
|------------|:---------------------------------------------------------------------------------------------------:|
| up         |                         Запускает Docker-контейнеры и пересобирает проект.                          |
| dev-up     |                               Запускает только контейнер базы данных.                               |
| restart    |                                    Перезапускает все контейнеры.                                    |
| build      |      Запускает контейнер базы данных, выполняет миграции и затем запускает остальные сервисы.       |
| rebuild    | Полностью пересобирает проект: останавливает контейнеры, удаляет данные и образы, запускает заново. |
| migrate    |                                   Выполняет миграции базы данных.                                   |
| down       |                               Останавливает и удаляет все контейнеры.                               |
| test-reg   |                           Запускает функциональные тесты для регистрации.                           |
| test-login |                         Запускает функциональные тесты для входа в систему.                         |
| test-asset |             Запускает функциональные тесты для создания, получения и удаления активов.              |

## Примеры запросов curl
| Описание                        |                                                                                               CURL команда                                                                                                |
|---------------------------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
| Регистрация нового пользователя |                              curl --location 'localhost:8080/api/register' \ --header 'Content-Type: application/json' \ --data '{ "username": "user", "password": "pass" }'                              |
| Вход в систему                  |                               curl --location 'localhost:8080/api/login' \ --header 'Content-Type: application/json' \ --data '{ "username": "user", "password": "pass" }'                                |
| Добавление нового актива        | curl --location 'localhost:8080/api/assets' \ --header 'Content-Type: application/json' \ --header 'Authorization: Bearer ****' \ --data '{ "name": "some_name4", "description": "desc", "price": 1500 }' |
| Получение списка активов        |                                                            curl --location 'localhost:8080/api/assets' \ --header 'Authorization: Bearer ****'                                                            |
| Удаление актива                 |                                                  curl --location --request DELETE 'localhost:8080/api/assets/1' \ --header 'Authorization: Bearer ****'                                                   |
| Покупка актива                  |              curl --location 'localhost:8080/api/assets/buy' \ --header 'Content-Type: application/json' \ --header 'Authorization: Bearer ****' \ --data '{ "asset_id": 4, "price": 1500 }'              |
| Получение активов пользователя  |                                                          curl --location 'localhost:8080/api/assets/my' \ --header 'Authorization: Bearer ****'                                                           |

Также можно воспользоваться коллекцией postman (testHive.postman_collection.json), которая лежит в корне репозитория. Там все те же запросы.