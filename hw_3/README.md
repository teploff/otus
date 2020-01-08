
```
git clone https://github.com/teploff/otus.git
cd otus/hw_3
go vet $(go list ./... | grep -v /vendor/)
golint $(go list ./... | grep -v /vendor/)
go test $(go list ./... | grep -v /vendor/)
```