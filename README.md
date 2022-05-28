# goweb
* Build / How to RUN:
```
1. go build

2. run exe
```

* Important packages:
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

* Gorm Docs:
```
https://gorm.io/docs/
```

* Mod commands:
```
1. go mod init learn
2. go mod tidy
```
