# service-status
## Итоговая аттестация по Golang от SkillBox

Необходимо реализовать сервис для объединения данных нагрузки на сервис.

## Установка и запуск

### Установка
1. Для удобства работы рекомендую создать отдельную папку для проекта 
2. Необходимо установить [симулятор данных](https://github.com/antondzhukov/skillbox-diploma) в корневую папку  
   * Установка симулятора `git clone https://github.com/antondzhukov/skillbox-diploma.git`
   * Переход в папку симулятора `cd .\skillbox-diploma\`
   * Инициация модуля `go mod init github.com/antondzhukov/skillbox-diploma`
   * Установка зависимостей `go get`
3. Установка приложения в корневую папку  
   * Переход в корневую папку `cd ..`
   * Установка приложения `git clone https://github.com/CyclopsV/service-status-skillbox.git`
   В результате должен получиться следующий каталог:
```
\---root-folder
    +---service-status-skillbox
    \---skillbox-diploma
```
### Запуск
1. Запустить симулятор данных
   * Перейти в папку симулятора `cd .\skillbox-diploma\`
   * Запустить симулятор `go run .`
2. Запуск приложения
   * Открыть новый терминал
   * Прейти в папку приложения `cd .\service-status-skillbox\`
   * Запустить приложение `go run cmd/service-status.go`

## Использование
Для получения данных необходимо отправить `get`-запрос на адрес http://127.0.0.1:8080/. Либо можно перейти по этой ссылке в браузере.
