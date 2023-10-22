# Запуск
Для начала Алеша, тебе нужно:
1. скачать компилятор Golang
2. зайти в папку server
3. вести оттуда команду `go run cmd/main.go`
4. дальше покажется хост с портом в консоли
# Какие есть Роуты
- GET: /health -> проверям шо все работает, должен придти 200 код
- POST: /api/client/setDesc -> сохраняет входные данные на серв и отправляет эти же данные как ответ:
input data -> json { name string, last_name string, patronymic string, phone string}
- POST: /api/client/checkVIN -> обрабатывает и сохраняет фотку в папку `photos`
input data -> form-data { VIN photo }
# Траблы
При возникновении траблов пиши в тг
