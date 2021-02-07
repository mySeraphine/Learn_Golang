package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

　　 //数据库连接
    db,_:=sql.Open("mysql","root:root@(127.0.0.1:3306)/golang")
    
    err :=db.Ping()
    if err != nil{
        fmt.Println("数据库链接失败")
    }
　　　

　　defer db.Close()
  }
