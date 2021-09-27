package gin

import (
	"net/http"
)

type HandleFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// ServeHTTP conforms to the http.Handler interface.
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.router.handle(newContext(w, req))
}

func (e *Engine) addRoute(method, path string, fun HandleFunc) {
	e.router.addRoute(method, path, fun)
}

func (e *Engine) GET(path string, fun HandleFunc) {
	e.addRoute(http.MethodGet, path, fun)
}

func (e *Engine) POST(path string, fun HandleFunc) {
	e.addRoute(http.MethodPost, path, fun)
}

func (e *Engine) Run(addr ...string) error {

	var defaultAddr = ":9999"

	if len(addr) == 1 {
		return http.ListenAndServe(addr[0], e)
	}
	return http.ListenAndServe(defaultAddr, e)
}
