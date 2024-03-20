package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{} //类型别名，为任意类型的映射

type Context struct {
	//origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	//request info
	Path   string
	Method string
	Params map[string]string
	//response info
	StatusCode int
	//middleware
	handlers []HandlerFunc
	index    int //记录执行到第几个中间件

	engine *Engine
}

// 创建实例
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 从请求中获取表单参数和URL查询参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 设置http响应状态码和响应头
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 返回字符串
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Context-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 返回json
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 返回字节数组
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// 返回html格式的http响应
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}

// 当在中间件在next方法时，控制权交给了下一个中间件
// 直到调用最后一个中间件，然后再从后往前调用每个中间件在next方法之后定义的部分
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}

}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})

}
