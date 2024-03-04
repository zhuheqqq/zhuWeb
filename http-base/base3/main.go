package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) { //第二个参数是一个匿名函数
		fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path) //返回请求给客户端
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	})

	r.Run(":9999")
}

//没有启动服务器之前，前面的注册路由和处理函数会正常执行，但不会立即出发处理函数，处理函数会在客户端发起请求并且服务器正常运行时才会被调用
