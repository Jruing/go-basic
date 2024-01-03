package views

import (
	"fmt"
	"go-basic/src/core"
	"net/http"
)

// 视图函数用户新增
func Useradd(w http.ResponseWriter, r *http.Request) {

	fmt.Println("新增用户")

	data := []string{"aaa", "bbb"}
	// json返回值
	returnVal := core.Response{
		Code: 0,
		Msg:  "",
		Data: data,
	}
	fmt.Fprintf(w, returnVal.ToString())
}
