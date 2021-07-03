FROM golang:latest

COPY src/api /go/src/api/

WORKDIR /go/src/api/

RUN go get -u github.com/cosmtrek/air \
  && go build -o /go/bin/air github.com/cosmtrek/air \
  && go get github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql

EXPOSE 8080

# CMD ["go", "run", "main.go"]

CMD ["air", "-c", ".air.toml"]


