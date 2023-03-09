package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/solenovex/web-tuborial/common"
	"github.com/solenovex/web-tuborial/controller"
)

const (
	server   = "localhost"
	port     = 1433
	user     = "sa"
	password = "123456"
	database = "user"
)

func init() {
	//连接sqlserver数据库
	connstr := fmt.Sprintf("server = %s;user id = %s;password = %s;port=%d; database=%s;",
		server, user, password, port, database)
	var err error
	common.Db, err = sql.Open("sqlserver", connstr)
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	err = common.Db.PingContext(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected!")
}

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: nil,
	}
	controller.RegisterRoutes()
	server.ListenAndServe()

}
