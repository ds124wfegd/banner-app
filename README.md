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

* Реализованы эндпоинты для регистрации и авторизации админа и пользователя. Для авторизации используется 2 вида JWT токенов: пользовательский и админский. Все действия кроме получения банеров доступны только для админов. Для пользователей используются id < 2000000000 назначаются автоматически. Для админов id начинаются с 2000000000, назначаются самим админом при регистрации.
* Для входа админу необходимо ввести пароль доступа к сервису "sekret"
* Для флага is_active принято следующие понятия: 1 - активен, 2 - неактивен
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

### регистрация пользователя
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

### авторизация пользователя
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
### регистрация админа
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
```

### авторизация админа
```
URL: localhost:8000/auth/sign-in-admin
```
![Снимок экрана (14681)](https://github.com/ds124wfegd/banner-app/assets/133537346/deb44a88-9139-4e0d-91fe-7aec0cf5ead6)

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
```
---
### Пример того, что пользователь не имеет доступа к другим эндпоинтам, кроме получения банеров
![Снимок экрана (14678)](https://github.com/ds124wfegd/banner-app/assets/133537346/4476bd55-daad-457f-b3b2-b4fded80e9be)
---
### создание банера админом
```
URL: localhost:8000/api_admin/bannerAdmin
```
![Снимок экрана (14682)](https://github.com/ds124wfegd/banner-app/assets/133537346/72fdf977-9a81-4258-853f-2d64a8ea2efe)

![Снимок экрана (14683)](https://github.com/ds124wfegd/banner-app/assets/133537346/a5d9234a-49c3-496b-a352-e6d701a9fc43)
входные данные:
```
{
    "isActive": 1,
    "featureId": 2,
    "tagId_1":1,
    "tagId_2":2,
    "tagId_3": 5
}
```
выходные:
```
{
    "id": 5
}
```
### Получение всех банеров админом или пользователем
```
URL: localhost:8000/api_admin/bannerAdmin/getAll/banner
```
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/a8ce86a9-c0fa-4b35-9a4d-fce97c5c1898)


входные данные:
```
{
}
```
выходные:
```
{
    "data": [
        {
            "id": 1,
            "isActive": 1,
            "featureId": 2,
            "tagId_1": 1,
            "tagId_2": 2,
            "tagId_3": 5
        },
        {
            "id": 2,
            "isActive": 2,
            "featureId": 4,
            "tagId_1": 5,
            "tagId_2": 1,
            "tagId_3": 5
        },
        {
            "id": 3,
            "isActive": 1,
            "featureId": 3,
            "tagId_1": 2,
            "tagId_2": 1,
            "tagId_3": 2
        },
        {
            "id": 4,
            "isActive": 2,
            "featureId": 3,
            "tagId_1": 2,
            "tagId_2": 4,
            "tagId_3": 2
        },
        {
            "id": 5,
            "isActive": 1,
            "featureId": 5,
            "tagId_1": 8,
            "tagId_2": 7,
            "tagId_3": 2
        }
    ]
}
```
### Получение банеров по id фичи админом или пользователем
```
URL: localhost:8000/api/banner/user_banner/feature=5
```
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/8db123ca-df10-4e35-a45f-a92a5c748a90)
входные данные:
```
{
}
```
выходные:
```
{
    "data": [
        {
            "id": 5,
            "isActive": 1,
            "featureId": 5,
            "tagId_1": 8,
            "tagId_2": 7,
            "tagId_3": 2
        }
    ]
}
```

### Получение банеров по id тега админом или пользователем
```
URL: localhost:8000/api/banner/user_banner/tag=1
```
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/ce2dbf18-cf0b-44e6-b0c0-6a55ff4a8328)

входные данные:
```
{
}
```
выходные:
```
{
    "data": [
        {
            "id": 1,
            "isActive": 1,
            "featureId": 2,
            "tagId_1": 1,
            "tagId_2": 2,
            "tagId_3": 5
        },
        {
            "id": 2,
            "isActive": 2,
            "featureId": 4,
            "tagId_1": 5,
            "tagId_2": 1,
            "tagId_3": 5
        },
        {
            "id": 3,
            "isActive": 1,
            "featureId": 3,
            "tagId_1": 2,
            "tagId_2": 1,
            "tagId_3": 2
        }
    ]
}
```
### Удаление банера по id фичи админом
```
URL: localhost:8000/api_admin/bannerAdmin/deleteByFeature/4
```

![image](https://github.com/ds124wfegd/banner-app/assets/133537346/48f693b1-418e-4eb4-9125-e98d566b28f4)

![image](https://github.com/ds124wfegd/banner-app/assets/133537346/9801cb37-08fe-4b00-83bf-f19fb40b826b)
входные данные:
```
{
}
```
выходные:
```
{
    "status": "ok"
}
```

### обновление тегов, фич, статуса работы банера админом
```
URL: localhost:8000/api_admin/bannerAdmin/updateBanner/5
```

![image](https://github.com/ds124wfegd/banner-app/assets/133537346/9842315b-d294-4b98-ba3e-6891d100ddfd)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/d64b34c1-5dd9-48c6-927c-11fad2fb8335)
входные данные:
```
{
}
```
выходные:
```
{
}
```
### Создание контента админом
```
URL: localhost:8000/api_admin/bannerAdmin/content/createContent=1
```
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/beb26414-379a-40fd-bae2-d876a07e1564)
входные данные:
```
{
    "bannerId":22,
	"title":"fdsf",
	"someTitle":"fdgdvdgfvsdf",
	"text":"fdfgds",
	"someText":"fdddthts",
	"someUrl":"fdf"
}
```
выходные:
```
{
    "id": 1
}
```
### Получение контента по Id банера
```
URL: localhost:8000/api_admin/bannerAdmin/content/createContent=1
```
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/6d4526e2-ce32-4a4c-9dfe-443ea236c563)
входные данные:
```
{
}
```
выходные:
```
[
    {
        "id": 2,
        "bannerId": 1,
        "title": "fddsfdsfsdfsdfsdfsf",
        "someTitle": "fdgdfdsfdsvdgfvsdf",
        "text": "fdfgds",
        "someText": "sdfsdfdsfsd",
        "someUrl": "fddfsfdsff"
    }
]
```
### удаление содержимого контентов банера по id контента
```
URL: localhost:8000/api_admin/bannerAdmin/content/deleteContnetById=2
```
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/610a067d-ae0c-48ec-a71e-0a782d290f2f)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/d25c61eb-c9fc-46f4-a006-afa9e5708999)
входные данные:
```
{
}
```
выходные:
```
{
    "status": "ok"
}
```

#Нагрузочное тестирование
авторизация для админа
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/635aeb95-fe0a-46d4-a33e-d6f0ddebc057)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/d4a0371b-ffeb-4a3c-bb26-36ee6b929789)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/78bd9659-2ef9-495a-a20a-db4891840e75)
![Снимок экрана (14714)](https://github.com/ds124wfegd/banner-app/assets/133537346/c46ee927-8b35-4794-9df5-04b679ee8c4e)


Тестирование: /api/banner/user_banner/feature=3
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/b6866659-f963-44db-bd41-536028d44bf2)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/0267baa4-b6bf-49a1-b4bf-8bcab6fb0f0e)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/cc1c6a8f-a38b-4538-bd71-e7e2b31ac5f2)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/6a819ed5-2487-4a9e-880b-96da0201de91)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/a26a7b57-5176-4bed-8eb1-f3631a6560fc)


Тестирование localhost:8000/api_admin/bannerAdmin
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/7641a53a-cca7-49f4-ac64-17f05788150c)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/91c19f59-7eed-49d5-abbe-4b68c5c0182c)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/9bf68a45-0230-432f-8c62-6e58132c9c22)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/b353b735-77fa-4c2a-93e5-f5492d9b752a)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/96e350cf-efcf-465e-bc07-3d4218acdbf3)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/051eb817-9bff-45c6-ac3f-b4079424b049)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/587a4040-e055-4739-a415-e2379bb06311)

Тестирование localhost:8000/api_admin/bannerAdmin/content/createContent=1

![image](https://github.com/ds124wfegd/banner-app/assets/133537346/9fe67873-7cf6-43f5-8b3f-54675b52361f)


![image](https://github.com/ds124wfegd/banner-app/assets/133537346/3b488b5a-7976-4bee-bb0a-623218e54f95)

![image](https://github.com/ds124wfegd/banner-app/assets/133537346/c7e69718-d3a6-4208-9f48-2931fe45aa1f)

![image](https://github.com/ds124wfegd/banner-app/assets/133537346/f5998160-edd0-4451-b98c-c98f15ce0d15)

![image](https://github.com/ds124wfegd/banner-app/assets/133537346/5c108c84-ee57-4166-ba3e-8aa16c0add59)

Тестирование: localhost:8000/api_admin/bannerAdmin/content/getAllContnetById=3
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/2b6b78d5-6a70-49ba-965f-a5039001dea4)

![image](https://github.com/ds124wfegd/banner-app/assets/133537346/1db6a2b7-860f-40cb-ac5d-9e95e6f08b72)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/2af02bbb-c632-4076-b28b-2ce46dfb581d)
![image](https://github.com/ds124wfegd/banner-app/assets/133537346/2023c944-be3e-46ba-8ceb-81c1399c6977)



























## Вопросы:
1. Есть ли необходимость реализовывать эндпоитны для регистрации и авторизации?
- Если рассматривать то, что разрабатывается микросервис, в этом необходимости нет, необходимо предоставить интерфейсы для сервисов регистрации и авторизации. Однако для данного задания принято решение реализовать данные методы.
2. 
