# url translater
Link shortener: test task for ozon <br/> 

## API

| Путь      | Метод   | Body (json)          | Описание                                         |
| --------- | --------| -------------------- | ------------------------------------------------ |
| /short    | POST    | {"url": "long_url"}  | Сформировать короткий url и добавить в хранилище |
| /long     | POST    | {"url": "short_url"} | Вернуть длинный url по значению короткого        |


## Возможные ошибки

| Ошибка    |  Описание                    |
| --------- | ---------------------------- |
| 400       |  Переданно некорректное боди |
| 500       |  internal error              |

## Запуск приложения
Eсли приложение запускается впервые, нужно сначала ввести следующую команду:
```
make build
```
Запуск:
```
make run
```
База данных уже запущена на heroku, веб-сервер доступен по адресу [http://localhost:8000](http://localhost:8000). 

## Проверка работоспособности
Запустите тесты, проверяющие правильность работы приложения, командой:
```
make test
```

Можно также проверить работу приложение через консоль:

```
curl -X POST -v -H "Connection: close" "Content-Type: application/json" -d '{
"url": "long_url"
}' http://localhost:8000/short
```

```
curl -X POST -v -H "Connection: close" "Content-Type: application/json" -d '{
"url": "short_url"
}' http://localhost:8000/long
```


