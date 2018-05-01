package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)

func main() {

	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go?charset=utf8");
	if err != nil {
		fmt.Println(err);
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用

	//defer db.Close();
	stmt, err := db.Prepare(`INSERT into userinfo(username,departname,created) values (?,?,?)`)
	result,_:=stmt.Exec("xiongtao","ui","2018-05-01")
	id,_:=result.LastInsertId();
	fmt.Println("id:",id,"has been inserted");
	rows, _ := db.Query(" SELECT uid,username,departname FROM userinfo")
	uid := 0
	username := ""
	departname := ""

	//查询所有
	for rows.Next() {
		rows.Scan(&uid, &username, &departname);
		fmt.Println("id:", uid, "departname", departname, "username", username)
	}
}
