package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	learnDb, err := sql.Open("mysql", "learn.go:123@1234@tcp(127.0.0.1:3389)/LearnGo")
	if err != nil {
		fmt.Println("连接数据库失败:", err)
	}
	defer learnDb.Close()
	if err = learnDb.Ping(); err != nil {
		fmt.Println("DB测试失败:", err)
	}
	fmt.Println("连接数据库成功")
	rows, err := learnDb.Query("select name from personal_information")
	if err != nil {
		fmt.Println("查询表数据失败")
	}
	fmt.Println(rows)
	// 解析
	if rows == nil {
		return
	}
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return
		}
		fmt.Println("name: ", name)
	}
	//fmt.Println(time.Now().String())
}
