package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/solenovex/web-tuborial/common"
	"github.com/solenovex/web-tuborial/funcs"
	"github.com/solenovex/web-tuborial/model"
)

func registerIndexRoutes() {
	http.HandleFunc("/index", Lists)
	http.HandleFunc("/add", add_list)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/delete/", delete)
	http.HandleFunc("/test", test)
}

//查询所有记录
func Lists(w http.ResponseWriter, r *http.Request) {
	//读取cookie
	//cookie, _ := r.Cookie("uid")
	//fmt.Println(cookie.Value)
	//cha quan xian
	//des, err := model.Get_purview(cookie.Value)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for _, j := range des {
	//	fmt.Println(j)
	//}
	//查询数据
	apps, err := model.Get_all()
	if err != nil {
		log.Fatalln(err)
	}
	funcMap := template.FuncMap{"add": funcs.Add}
	t := template.New("companies").Funcs(funcMap)
	t, err = t.ParseFiles("templates/lyout.html", "templates/list.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "lyout", apps)
}

//添加一条记录
func add_list(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("uid")
	var ok bool
	//cha quan xian
	des, err := model.Get_purview(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}

	switch r.Method {
	case http.MethodGet:
		for _, j := range des {
			if j == "增加" {
				ok = true
				break
			}
		}
		if ok {
			t, err := template.ParseFiles("templates/lyout.html", "templates/add.html")
			if err != nil {
				fmt.Println(err)
			}
			t.ExecuteTemplate(w, "lyout", nil)
		} else {
			t, err := template.ParseFiles("templates/lyout.html", "templates/without.html")
			if err != nil {
				log.Fatalln(err)
			}
			t.ExecuteTemplate(w, "lyout", nil)
		}
	case http.MethodPost:
		record := common.App{}
		record.Id = r.FormValue("number")
		record.Key = r.FormValue("name")
		record.Role = r.FormValue("password")
		err := model.Add_one(record.Id, record.Key, record.Role)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalln(err)
		} else {
			http.Redirect(w, r, "/index", http.StatusFound)
		}
	}
}

//编辑
func edit(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("uid")
	var ok bool
	//cha quan xian
	des, err := model.Get_purview(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	switch r.Method {

	case http.MethodGet:
		//pan duan quanxian
		for _, j := range des {
			if j == "编辑" {
				ok = true
				break
			}
		}
		if ok {
			query := r.URL.Query()
			id := query.Get("id")
			list := model.Get_one(id)
			t, err := template.ParseFiles("templates/lyout.html", "templates/edit.html")
			if err != nil {
				log.Fatalln(err)
			}
			t.ExecuteTemplate(w, "lyout", list)
		} else {
			t, err := template.ParseFiles("templates/lyout.html", "templates/without.html")
			if err != nil {
				log.Fatalln(err)
			}
			t.ExecuteTemplate(w, "lyout", nil)
		}

	case http.MethodPost:
		record := common.App{}
		//query := r.URL.Query()
		//record.Id = query.Get("id")
		record.Id = r.FormValue("number")
		record.Key = r.FormValue("name")
		err := model.Edit(record)
		if err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("uid")
	var ok bool
	//cha quan xian
	des, err := model.Get_purview(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	//pan duan quanxian
	for _, j := range des {
		if j == "删除" {
			ok = true
			break
		}
	}
	if ok {
		query := r.URL.Query()
		Id := query.Get("id")
		err := model.Delete(Id)
		if err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, "/index", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		//t, err := template.ParseFiles("templates/lyout.html", "templates/without.html")
		//if err != nil {
		//log.Fatalln(err)
		//}
		//t.ExecuteTemplate(w, "lyout", nil)
	}

}

func test(w http.ResponseWriter, r *http.Request) {
	uid := "001"
	many, err := model.Test(uid)
	if err != nil {
		log.Fatalln(err)
	}
	for _, j := range many {
		fmt.Println(j.Id)
		fmt.Println(j.Name)
		for _, n := range j.Purview {
			fmt.Print(n)
		}

	}
}
