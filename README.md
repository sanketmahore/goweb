# goweb
* Build / How to RUN:
```
1. go build

2. run exe
```

* Importent packages:
```
"github.com/gorilla/mux"
"net/http"
"github.com/jinzhu/gorm"
"github.com/go-sql-driver/mysql"
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/mock"
```

* Unit tests:
```
1. install
    go get github.com/vektra/mockery/v2/.../
2. Generate mock interface:
    mockery --case snake --name BookingService --dir domain
```

* GRPC Integration
```
Login api interacts with grpc server to fetch login is valid or not
i.e. POST http://localhost:10000/login - returns true or false

1. Installation of the following libraries for grpc:-
   1. go get -u google.golang.org/grpc
   2. go get -u github.com/golang/protobuf/protoc-gen-go

2. Download protoc-3.18.0-win64.zip from https://github.com/protocolbuffers/protobuf/releases

3. Unzip protoc-3.18.0-win64 and place the protoc.exe file in your GOPATH/bin folder

4. For creating a go equivalent file use the following command:-
   protoc authentication/login.proto --go_out=plugins=grpc:.

5. Use the following commands to run grpc server:-
   go run authServer.go
```

* Gorm Docs:
```
https://gorm.io/docs/
```

* Mod commands:
```
1. go mod init learn
2. go mod tidy
```