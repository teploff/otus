Запуск линтеров
```
git clone https://github.com/teploff/otus.git
cd otus/calendar
go vet $(go list ./... | grep -v /vendor/)
$(go list -f {{.Target}} golang.org/x/lint/golint) $(go list ./... | grep -v /vendor/)
```
Запуск проекта
```shell script
make run
```
Остановка проекта
```shell script
make shutdown
```
Запуск интеграционных тестов
```shell script
make integration_test
```