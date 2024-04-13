# Тестовое задание для стажёра Backend в Авито (информация будет добавляться в течение дня...)
                                                        
В Авито есть большое количество неоднородного контента, для которого необходимо иметь единую систему управления.  В частности, необходимо показывать разный контент пользователям в зависимости от их принадлежности к какой-либо группе. Данный контент мы будем предоставлять с помощью баннеров.

# Описание задачи
Необходимо реализовать сервис, который позволяет показывать пользователям баннеры, в зависимости от требуемой фичи и тега пользователя, а также управлять баннерами и связанными с ними тегами и фичами.

# Используемые технологии:
Технология | Применение
:----------|:--------:|
Visual Studio Code | редактор кода
PostgreSQL | в качестве хранилища данных
Docker |для запуска образа PostgreSQL
Postman |тестирование сценариев API
Apache JMeter| инструмент для проведения нагрузочного тестирования
golang-migrate/migrate | для миграций БД
github.com/spf13/viper| чтение из конфига
github.com/jmoiron/sqlx|для работы с postgres
github.com/joho/godotenv|для получения паролей из переменных окружения
github.com/sirupsen/logrus|для логирования
github.com/dgrijalva/jwt-go|для JWT-токенов
github.com/gin-gonic/gin| высокопроизводительный веб-фреймворк HTTP (для маршрутизации)

# Решение
* Сервис REST API написан на Golang v 1.22 с использованием Clean Architecture, что позволяет легко расширять функционал сервиса и тестировать его. Также был реализован Graceful Shutdown для корректного завершения работы сервиса.

* Реализованы эндпоинты для регистрации и авторизации админа и пользователя. Для авторизации используется 2 вида JWT токенов: пользовательский и админский. Все действия кроме получения банеров доступны только для админов.
* Реализованы все интеграционные-тесты для различных сценариев.
* Считаю, что при небольшом количестве фич и тегов нет необходимости использовать флаг use_last_revision. Все операции в данном сервисе осуществляются напрямую с базой. При повышении нагрузки будет необходимо реализовать кеширование на стороне пользователя
* При использовании метода GetAll() пользователи не получают банеры, которые отключены. К ним доступ есть только для админов.
* Добавлен метод удаления банеров по id фичи



## Иструкция по запуску докер контейнера Postgres:

в CMD:

* docker pull postgres
* docker run --name=banner-db -e POSTGRES_PASSWORD='12345' -p 5438:5432 -d --rm postgres
* docker ps // для копирования номера контейнера
* migrate create -ext sql -dir ./schema -seq init (это строка не нужна, но она использовалась, при первоначальной инициализации)
* migrate -path ./schema -database postgres://postgres:'12345'@localhost:5438/postgres?sslmode=disable up
* docker exec -it  ___номер контейнера___  /bin/bash
* psql -U postgres
* \d (для просмотра таблиц)
* \q (для выхода)
* exit  (для выхода)

(если пароль для базы данных неверен)
* \password postgres (для сброса пароля)
* дважды ввести новый пароль 12345

## Для запуска самого приложения:
* go run cmd/main.go
Для остановки:
* Ctrl + C


## Описание эндпоинтов:
метод|путь| назначение
:----|:---|:----------
POST  | /auth/sign-up      | регистрация пользователя       
POST  | /auth/sign-in      | авторизация пользователя     
POST  |/auth/sign-up-admin| регистрация админа
POST  | /auth/sign-in-admin|авторизация админа   
GET   | /api/banner/user_banner/feature=:id |получение банера пользователем по id фичи
GET   | /api/banner/user_banner/tag=:id     |получение банера пользователем по id тега
GET   | /api/banner/user_banner/banner      |получение все банеров пользователем, которые включены
POST  | /api_admin/bannerAdmin/             | создание админом банера
GET   | /api_admin/bannerAdmin/admin_banner/feature=:id     | получение банера админом по id фичи
GET   | /api_admin/bannerAdmin/admin_banner/tag=:id     | получение банера админом по id тега
GET   | /api_admin/bannerAdmin/getAll/banner            |получение всех банеров админом
GET   | /api_admin/bannerAdmin/updateBanner/:id         | обновление тегов, фич, статуса работы банера админом
DELETE |/api_admin/bannerAdmin/deleteByFeature/:id      | удаление банера админом по id фичи
POST   |/api_admin/bannerAdmin/content/createContent=:id    | создание контента банера    
GET    |/api_admin/bannerAdmin/content/getAllContnetById=:id  | получение содержимое контентов всех банеров 
DELETE |/api_admin/bannerAdmin/content/deleteContnetById=:id  | удаление содержимого контентов банера по id контента


## Usage

###регистрация пользователя
```
URL: localhost:8000/auth/sign-up
```

![image](https://github.com/ds124wfegd/banner-app/assets/133537346/111e6f54-cd50-4422-b1b9-c24fed2cc1e6)

входные данные:
```
{
    "name": "Nicolay_2",
    "username": "Nicolas_2",
    "password": "12345"
}
```
выходные:
```
"id": 6
```

###авторизация пользователя
```
URL: localhost:8000/auth/sign-in
```
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/e0e4367a-6d07-46b1-92cc-ac24a570fefd)

входные данные:
```
{
    "username": "Nicolas_2",
    "password": "12345"
}
```
выходные:
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMwNzM4OTUsImlhdCI6MTcxMzAzMDY5NSwidXNlcl9pZCI6Nn0.cIhKXB6nTFLlfZGt5z3cR6yQPu1aKbQmW0DoEcaT5zw"
}
```
###регистрация админа
```
URL: localhost:8000/auth/sign-up-admin
```
![Снимок экрана (14680)](https://github.com/ds124wfegd/banner-app/assets/133537346/4fc6eb76-b968-4b74-80c7-0531ac58da33)
входные данные:
```
{
    "id":2000000002,
    "adminUsername": "Vasiliy12345",
    "adminPassword": "12345",
    "adminStatus":true,
    "systemPasword":"sekret"
}
```
выходные:
```
{
    "id": 2000000002
}

### авторизация админа
```
URL: localhost:8000/auth/sign-in-admin
```
![Снимок экрана (14680)](https://github.com/ds124wfegd/banner-app/assets/133537346/72d94ca4-c099-41fc-ae60-dfaf2556689d)
входные данные:
```
{
    "adminUsername": "Vasiliy12345",
    "adminPassword": "12345"
}
```
выходные:
```
{
    "tokenAdmin": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMwNzQ1NTEsImlhdCI6MTcxMzAzMTM1MSwidXNlcl9pZCI6MjAwMDAwMDAwMn0.vc9vbxp2oPEzaihV-J1Y1R177_2us2lm7DAstW-0_6g"
}

---
### Пример того, что пользователь не имеет доступа к другим эндпоинтам, кроме получения банеров
![Снимок экрана (14678)](https://github.com/ds124wfegd/banner-app/assets/133537346/4476bd55-daad-457f-b3b2-b4fded80e9be)

---
## Вопросы:
1. Есть ли необходимость реализовывать эндпоитны для регистрации и авторизации?
- Если рассматривать то, что разрабатывается микросервис, в этом необходимости нет, необходимо предоставить интерфейсы для сервисов регистрации и авторизации. Однако для данного задания принято решение реализовать данные методы.
2. 
