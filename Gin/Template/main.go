package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*type User struct {
	Name   string
	Gander string
	Age    int
}*/

func f(w http.ResponseWriter, r *http.Request) {
	//定义模板
	k := func(name string) (string, error) {
		return name + "不要失去发芽的心情！", nil
	}
	t := template.New("f.tmpl")
	t.Funcs(template.FuncMap{
		"kua": k,
	})

	_, err := t.ParseFiles("/home/zhuheqin/clone/zhuWeb/Gin/Template/f.tmpl")
	if err != nil {
		//fmt.Println("http server failed:%V", err)
		return
	}

	//渲染模板
	err = t.Execute(w, "zhuheqin!")
	if err != nil {
		//fmt.Println("http server failed:%V", err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", f)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http server failed:%V", err)
		return
	}
}
