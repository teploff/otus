
```
git clone https://github.com/teploff/otus.git
cd otus/hw_2
go vet $(go list ./... | grep -v /vendor/)
$(go list -f {{.Target}} golang.org/x/lint/golint) $(go list ./... | grep -v /vendor/)
go test -v $(go list ./... | grep -v /vendor/)
```