## Web app in golang

### Requirements

- go compiler
- mysql server with user letme -p hello
- use sql code in ./db

### For building and running the server

```console
$ go get golang.org/x/oauth2
$ go get golang.org/x/oauth2/google
$ go get cloud.google.com/go
$ go get github.com/go-sql-driver/mysql
$ git clone https://github.com/augusto-mantilla/letgoweb.git
$ cd letgoweb
letgoweb $ go build src/main/main.go
letgoweb $ ./main
```