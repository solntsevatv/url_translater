# url translater
Test task for ozon

## API

| Путь      | Метод   | Body (json)          | Описание                                         |
| --------- | --------| -------------------- | ------------------------------------------------ |
| /short    | POST    | {"url": "long_url"}  | Сформировать короткий url и добавить в хранилище |
| /long     | POST    | {"url": "short_url"} | Вернуть длинный url по значению короткого        |

## Проверка работоспособности
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
