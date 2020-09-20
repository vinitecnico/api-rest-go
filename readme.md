## Exemplo api restful in go

```
go run main.go
http://localhost:8000/contato -- GET
http://localhost:8000/contato/{id} -- GET
http://localhost:8000/contato/{id} -- POST
http://localhost:8000/contato/{id} -- DELETE
```

### run test
```
go test -v

go test -coverprofile cover.out
```