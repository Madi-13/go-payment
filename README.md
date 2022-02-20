![GitHub top language](https://img.shields.io/github/languages/top/Madi-13/go-payment?style=for-the-badge)
![Lines of code](https://img.shields.io/tokei/lines/github/Madi-13/go-payment?color=yellow&style=for-the-badge)
![GitHub repo size](https://img.shields.io/github/repo-size/Madi-13/go-payment?color=yellow&style=for-the-badge)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Madi-13/go-payment?style=for-the-badge)
![GitHub last commit](https://img.shields.io/github/last-commit/Madi-13/go-payment?style=for-the-badge)
![GitHub contributors](https://img.shields.io/github/contributors/Madi-13/go-payment?color=gree&style=for-the-badge)

![GitHub watchers](https://img.shields.io/github/watchers/Madi-13/go-payment?style=social)

# Финансовая грамотность

## Описание задания:



Разработать go API "Финансовая грамотностсь" 

по движению денежных средств и текущему финансовосу состоянию.


![logo](png/go.jpeg)
                                    
## Реализовано:

+ Стуктура Базы Данных (PostgreSQL):
  + Категории (справочник)
  + Платежи (Доход или расход)

+ Функции
  + Управление платежами (CRUD)
  + Управление категориями (CRUD)

## Пример работы:

Запускаем приложение командой:

```
go run main.go
```

### Протестируем функцию getPaymentsBYID:

Отправляем GET запрос на локальный хост (я это делаю через POSTMAN) localhost:8080/payments/3

Уже созданной заранее нами базе по id = 3 лежит:

```
{
    "id": 3,
    "name": "мгу",
    "price": 65.5,
    "date": "2021-12-12",
    "type": "payment",
    "comment": "",
    "CategoryID": 2,
    "Category": {
        "ID": 2,
        "Name": "travel"
    }
}
```

Соответственно, после нашего запроса по id = 3 должен вернуться результат с телом вышеуказонного сегмента.

### Получаемый результат:

![logo](png/1.png)

Аналогично с остальными запросами.
