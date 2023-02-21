package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	MysqlUsersUsername = "mysql_users_username"
	MysqlUsersPassword = "mysql_users_password"
	MysqlUsersHost     = "mysql_users_host"
	MysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(MysqlUsersUsername)
	password = os.Getenv(MysqlUsersPassword)
	host     = os.Getenv(MysqlUsersHost)
	schema   = os.Getenv(MysqlUsersSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("database successfully configured")
}
