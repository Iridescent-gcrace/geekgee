package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Egine struct {
	router *router
}

func New() *Egine {
	return &Egine{router: newRouter()}
}

func (e *Egine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.router.addRouter(method, pattern, handler)
}

func (e *Egine) GET(path string, handler HandlerFunc) {
	e.addRouter("GET", path, handler)
}

func (e *Egine) POST(path string, handler HandlerFunc) {
	e.addRouter("POST", path, handler)
}

func (e *Egine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Egine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)

}
