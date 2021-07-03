# Dockerfile
FROM golang:latest

# コンテナ内に作業ディレクトリを作成
 
ENV SRC_DIR=/go/src/work

RUN mkdir $SRC_DIR

ENV GOBIN=/go/bin
 
# WORKDIR $GOBIN
WORKDIR $SRC_DIR
 
ADD ./api $SRC_DIR
 
RUN cd /go/src/;
 
# Install dependency module
RUN go get github.com/go-sql-driver/mysql \
    && go get -u github.com/gin-gonic/gin \
    && go get github.com/gorilla/mux \
    && go get -u github.com/jinzhu/gorm \
    && go get github.com/gin-contrib/cors \
    && go get gopkg.in/ini.v1
 
ENTRYPOINT ["go", "run", "main.go"]
