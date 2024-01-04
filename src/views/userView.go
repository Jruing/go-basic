package views

import (
	"basic/src/core"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/tidwall/gjson"
	"log"
	"net/http"
)

// 视图函数用户新增
func Useradd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("新增用户")
	// PostgreSQL 连接信息
	connStr := "user=myuser dbname=mydb sslmode=disable"
	// 打开数据库连接
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	gjson

	// 自定义json返回值
	//data := []string{"aaa", "bbb"}
	//returnVal := core.Response{
	//	Code: 0,
	//	Msg:  "this is response body",
	//	Data: data,
	//}

	core.Err.ToString()
	// 内置json返回值
	fmt.Fprintf(w, core.Err.ToString())
}
