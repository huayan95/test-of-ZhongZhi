package database

import (
 "database/sql"
 _ "github.com/go-sql-driver/mysql"
 "log"
 
)

var SqlDB *sql.DB

func init() {
 var err error
 SqlDB, err = sql.Open("mysql", "root:@/test?parseTime=true") //链接mysql中的数据库test
 if err != nil {
  log.Fatal(err.Error())
 }
 err = SqlDB.Ping()
 if err != nil {
  log.Fatal(err.Error())
 }
}