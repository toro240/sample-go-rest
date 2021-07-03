package mysql

import (
	"errors"
	"fmt"
	"log"
	"mj-app/exceptions"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CONSTANT VALUES
const (
	DBTYPE = "mysql"
	SCHEMA = "%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local"
)

var (
	Client         *gorm.DB
	username       = os.Getenv("MYSQL_USER")
	password       = os.Getenv("MYSQL_PASSWORD")
	dbName         = os.Getenv("MYSQL_DATABASE")
	datasourceName = fmt.Sprintf(SCHEMA, username, password, dbName)
)

func init() {
	var err error
	Client, err = gorm.Open(DBTYPE, datasourceName)
	count := 0
	if err != nil {
		for {
			if err == nil {
				break
			}
			time.Sleep(time.Second)
			count++
			if count > 180 {
				log.Fatal(err)
			}
			Client, err = gorm.Open(DBTYPE, datasourceName)
		}
	}

	log.Println("database successfully configure")
}

func ParseError(err *gorm.DB) *exceptions.ApiError {

	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return exceptions.NewNotFoundError("no record matching given id")
	}

	return exceptions.NewInternalServerError("error processing request")
}
