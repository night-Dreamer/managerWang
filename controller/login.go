package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/solenovex/web-tuborial/common"
)

func registerLoginRoutes() {
	http.HandleFunc("/login", login)
}

//登录表单handle
func login(w http.ResponseWriter, r *http.Request) {
	//处理Form post请求
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, nil)
	} else {
		name := r.FormValue("Name")
		password := r.FormValue("password")
		//cha mi ma
		a := common.App{}
		err := common.Db.QueryRow("SELECT userName, password FROM dbo.User_1 WHERE userName=@Id",
			sql.Named("Id", name)).Scan(&a.Id, &a.Key)
		if err != nil {
			fmt.Println(err)
		}
		//cookie 设置结构体
		b := common.App{}
		err = common.Db.QueryRow("SELECT u_id FROM dbo.User_1 WHERE userName=@Id",
			sql.Named("Id", name)).Scan(&b.Id)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(b.Id)

		if name != "" {
			if password == a.Key {
				//设置cookie
				cookie := http.Cookie{Name: "uid", Value: b.Id}
				http.SetCookie(w, &cookie)
				//定向到Index页
				http.Redirect(w, r, "/index", http.StatusFound)
			} else {
				fmt.Fprintln(w, "用户名或密码错误")
			}
		}
	}
}
