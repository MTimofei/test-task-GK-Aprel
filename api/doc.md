# REST API

## EndPoints 
- Для авторизации пользователя, используем Get запрос с Query содержащим login и password по пути /auth. В ответе Header содержит StatusCod 200 и X-Token token 

```
Request
GET http://domain/auth?login=username&password=password
Response
    Header
        StatusCod   200 
        X-Token     token
```
Ответ с StatusCod 403 означает блокировку аккаунта.<br>
Ответ с StatusCod 400 означает что login или password неверный.<br>
Ответ с StatusCod 405 означает что используется не верный метод запроса.<br>
Ответ с StatusCod 500 означает что произошла ошибка на сервере.<br>

- Для получения данных аудита аккаунта, используется Get запрос по пути /audit и Header X-Token token.
В ответе Header содержит StatusCod 200, а в Body лежит масив содержащий время и действия с аккаунтом в формате json
```
Request
Get http://domain/audit
    Header
        X-Token     token
Response
    Header
       StatusCod    200 
    Body
        {
           [
                { 
                "event_time"="timestamp",
                "event_type"="text"
                }...
           ]
        }
```
Ответ с StatusCod 403 означает не валидный token.<br>
Ответ с StatusCod 404 означает что token не найден.<br>
Ответ с StatusCod 405 означает что используется не верный метод запроса.<br>
Ответ с StatusCod 500 означает что произошла ошибка на сервере.<br>
- Для удаления аудита аккаунта, используется Delete запрос по пути /audit/clear и Header X-Token token.
В ответе Header содержит StatusCod 200.
```
Request
Get http://domain/audit
    Header
        X-Token     token
Response
    Header
       StatusCod    200 
```
Ответ с StatusCod 403 означает не валидный token.<br>
Ответ с StatusCod 404 означает что token не найден.<br>
Ответ с StatusCod 405 означает что используется не верный метод запроса.<br>
Ответ с StatusCod 500 означает что произошла ошибка на сервере.<br>