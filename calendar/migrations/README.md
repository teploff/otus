Использование goose-migration

1.Компилируем мигратор:
```shell script
go build -o goose *.go
```
2.Проверить статус миграций:
```shell script
./goose --host=127.0.0.1 --port=5438 --user=postgres --password=password --dbname=otus --sslmode=disable status
```
3.Накатить миграцию:
```shell script
./goose --host=127.0.0.1 --port=5438 --user=postgres --password=password --dbname=otus --sslmode=disable up
```
4.Откатить миграцию:
```shell script
./goose --host=127.0.0.1 --port=5438 --user=postgres --password=password --dbname=otus --sslmode=disable down
```