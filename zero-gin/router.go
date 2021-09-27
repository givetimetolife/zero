package gin

import "fmt"

type IRoutes struct {
}

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		handlers: map[string]HandleFunc{},
	}
}

func (r router) addRoute(method, path string, fun HandleFunc) {
	r.handlers[fmt.Sprintf("%s-%s", method, path)] = fun
}

func (r *router) handle(c *Context) {
	key := fmt.Sprintf("%s-%s", c.Request.Method, c.Request.URL.Path)

	fmt.Printf("%+v\n", r.handlers)
	if handler, ok := r.handlers[key]; ok {
		handler(c)
		return
	}
	fmt.Fprintf(c.Writer, "404 not found !!")
}
