package gin

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request

	Path   string
	Method string

	StatusCode int

	mu sync.RWMutex

	Keys map[string]interface{}
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Path:    req.URL.String(),
		Method:  req.Method,
		// Keys:    map[string]interface{}{},
	}
}

func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	if c.Keys == nil {
		c.Keys = map[string]interface{}{}
	}
	c.Keys[key] = value

	c.mu.Unlock()
}

func (c *Context) Get(key string) (value interface{}, exists bool) {

	if c.Keys == nil {
		return
	}

	c.mu.RLock()

	value, exists = c.Keys[key]

	c.mu.RUnlock()
	return
}

// request
func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// response
func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, value string) {
	c.Status(code)
	c.Writer.Header().Set("Context-Type", "text/plain")

	c.Writer.Write([]byte(value))
}

func (c *Context) JSON(code int, value interface{}) {
	c.Status(code)
	c.Writer.Header().Set("Context-Type", "application/json")

	encoder := json.NewEncoder(c.Writer)

	if err := encoder.Encode(value); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.Writer.Header().Set("Context-Type", "text/html")

	c.Writer.Write([]byte(html))
}

func (c *Context) DATA(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
