package repository

import (
	"database/sql"
	"fmt"
	"goTodoProject/config"

	_ "github.com/go-sql-driver/mysql"
)

const (
	TableNameTodos = "todos"
)

var DbConnection *sql.DB

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.Config.SQLUsername, config.Config.SQLPassword, config.Config.SQLAddress, config.Config.DbName)
	DbConnection, err = sql.Open(config.Config.SQLDriver, dsn)
	if err != nil {
		fmt.Printf("Fail to open DB: %v", err)
		return
	}
	cmd := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id BINARY(16) PRIMARY KEY NOT NULL,
		title varchar(256) NOT NULL,
		is_complited tinyint(1) NOT NULL)`, TableNameTodos)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		fmt.Printf("Fail to open DB: %v", err)
		return
	}
}
