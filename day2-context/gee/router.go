package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

// 创建映射实例
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 向路由表中添加新的路由信息
func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router %4s-%s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 处理请求
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND:%s\n", c.Path)
	}
}
